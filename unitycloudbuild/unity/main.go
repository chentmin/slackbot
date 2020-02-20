package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/antihax/optional"
	swagger "github.com/chentmin/slackbot/unitycloudbuild/unity/api"
	"github.com/pkg/errors"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

var help = `未知的命令
build tag [clean]
`

var buildCommand = regexp.MustCompile(`<@.+> build (\S+)( clean)?`)
var pingCommand = regexp.MustCompile(`<@.+> ping(.*)`)

var TOKEN = os.Getenv("SLACK_TOKEN")
var api = slack.New(TOKEN)

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, event *events.APIGatewayProxyRequest) (Response, error) {
	var VERIFICATION_TOKEN = os.Getenv("SLACK_VERIFICATION_TOKEN")

	body := event.Body

	eventsAPIEvent, e := slackevents.ParseEvent(json.RawMessage(body), slackevents.OptionVerifyToken(slackevents.TokenComparator{VerificationToken: VERIFICATION_TOKEN}))
	if e != nil {
		return newResponse(http.StatusInternalServerError, e.Error()), e
	}

	if eventsAPIEvent.Type == slackevents.URLVerification {
		var r *slackevents.ChallengeResponse
		err := json.Unmarshal([]byte(body), &r)
		if err != nil {
			return newResponse(http.StatusInternalServerError, err.Error()), err
		}

		return newResponse(http.StatusOK, r.Challenge), nil
	}

	if eventsAPIEvent.Type == slackevents.CallbackEvent {
		innerEvent := eventsAPIEvent.InnerEvent
		switch ev := innerEvent.Data.(type) {
		case *slackevents.AppMentionEvent:
			fmt.Printf("收到mention事件: %s: %s\n", ev.User, ev.Text)

			text := strings.TrimSpace(ev.Text)

			switch{
			// 启动构建
			case buildCommand.MatchString(text):
				match := buildCommand.FindStringSubmatch(text)
				tag := match[1]
				clean := len(match) > 2 && match[2] == " clean"

				if err := triggerUnityBuild(ctx, tag, clean, ev.Channel); err != nil {
					api.PostMessage(ev.Channel, slack.MsgOptionText(err.Error(), false))
				}

			// 返回ping的内容, 测试
			case pingCommand.MatchString(text):
				match := pingCommand.FindStringSubmatch(text)
				msg := "pong"
				if len(match) > 1 && match[1] != ""{
					msg = match[1]
				}
				api.PostMessage(ev.Channel, slack.MsgOptionText(msg, false))

			default:
				// 没有match
				api.PostMessage(ev.Channel, slack.MsgOptionText(help, false))
			}

			return newResponse(http.StatusOK, ""), nil
		}
	}

	fmt.Printf("收到未知的事件: %s\n", body)
	return newResponse(http.StatusBadRequest, "unknown event type: "+body), errors.New("unknown request")
}

func unityClient() *swagger.APIClient {
	cfg := swagger.NewConfiguration()
	cfg.BasePath = "/api/v1"
	cfg.Host = "build-api.cloud.unity3d.com"
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", "Basic "+os.Getenv("UNITY_TOKEN"))

	client := swagger.NewAPIClient(cfg)

	return client
}

func triggerUnityBuild(ctx context.Context, tag string, clean bool, slackChannel string) error {

	client := unityClient()

	option := &swagger.StartBuildsOpts{
		Options: optional.NewInterface(swagger.InlineObject9{
			Clean: clean,
			Delay: 5,
		}),
	}

	builds, _, err := client.BuildsApi.StartBuilds(ctx, os.Getenv("UNITY_ORG"), os.Getenv("UNITY_PROJECT"), tag, option)

	if err != nil {
		return errors.Wrap(err, "调用unity接口出错")
	}

	if payload := builds; payload == nil || len(payload) == 0 {
		return errors.New("unity没有返回错误, 也没有返回payload...")
	} else {
		fmt.Printf("收到%d个payload\n", len(payload))

		for i, p := range payload {
			fmt.Printf("payload[%d] build: %v, error: %s\n",i, p.Build, p.Error)
			if i > 0{
				continue
			}
			if err := p.Error; err != "" {
				return errors.New(fmt.Sprintf("unity返回错误: build: %v: %s", p.Build, err))
			}

			api.PostMessage(slackChannel, slack.MsgOptionText(fmt.Sprintf("启动成功: %s: %v", tag, p.Build), false))
		}

		return nil
	}

	return errors.New("unity返回没处理?")
}

func triggerUnityCancel(tag string, buildNumber int) (err error) {
	return
}

func newResponse(statusCode int, content string) Response {
	resp := Response{
		StatusCode:      statusCode,
		IsBase64Encoded: false,
		Body:            content,
		Headers: map[string]string{
			"Content-Type": "text",
		},
	}
	return resp
}

func main() {
	lambda.Start(Handler)
}
