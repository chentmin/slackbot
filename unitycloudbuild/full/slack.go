package main

import (
	"bytes"
	"context"
	"fmt"

	"github.com/antihax/optional"
	swagger "github.com/chentmin/slackbot/unitycloudbuild/api"
	"github.com/pkg/errors"
	qrcode "github.com/skip2/go-qrcode"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"os"

	"strings"
)

var (
	buildCommand   = `<@.+> build (\S+)( clean)?`
	pingCommand    = `<@.+> ping(.*)`
	installCommand = `<@.+> install (\S+) (\S+)`

	TOKEN         = os.Getenv("SLACK_TOKEN")
	api           = slack.New(TOKEN)
	UNITY_ORG     = os.Getenv("UNITY_ORG")
	UNITY_PROJECT = os.Getenv("UNITY_PROJECT")

	help = `未知的命令
build tag [clean] 新建构建
install tag number 获得安装二维码
`
)

func processCancelBuild(ctx context.Context, action slack.InteractionCallback) {
	if action.ActionCallback.AttachmentActions == nil {
		fmt.Printf("没有actions字段\n")
		return
	}

	if l := len(action.ActionCallback.AttachmentActions); l == 0 {
		fmt.Printf("没有actions字段\n")
		return
	} else if l > 1 {
		fmt.Printf("竟然超过了1个actions, 有%d个, 取第1个\n", l)
	}

	click := action.ActionCallback.AttachmentActions[0]

	value := click.Value
	fmt.Printf("取消构建: %s\n", value)
	v := strings.SplitN(value, "_", 2)
	if len(v) != 2 {
		fmt.Printf("value malform: %s\n", value)
		return
	}

	buildNum := v[0]
	tag := v[1]
	if err := triggerUnityCancel(ctx, tag, buildNum, action); err != nil {
		fmt.Printf("取消失败: %s\n", err)
		return
	}
}

func processInstallCommand(ctx context.Context, ev *slackevents.AppMentionEvent, cmd []string) {
	tag := cmd[1]
	buildNumber := cmd[2]

	url := fmt.Sprintf("%s/install?tag=%s&build=%s", os.Getenv("SELF_URL"), tag, buildNumber)

	fmt.Printf("image url: %s\n", url)

	png, err := qrcode.Encode(url, qrcode.Medium, 256)
	if err != nil {
		fmt.Printf("生成二维码失败: %s\n", err)
		api.PostMessage(ev.Channel, slack.MsgOptionPostEphemeral(ev.User), slack.MsgOptionText(fmt.Sprintf("生成二维码失败: %s\n", err), false))
		return
	}

	if _, err := api.UploadFile(slack.FileUploadParameters{
		Reader:          bytes.NewBuffer(png),
		Filetype:        "png",
		Filename:        fmt.Sprintf("%s_%s.png", buildNumber, tag),
		Channels:        []string{ev.Channel},
		ThreadTimestamp: ev.ThreadTimeStamp,
	}); err != nil {
		fmt.Printf("上传图片失败: %s\n", err)
		api.PostMessage(ev.Channel, slack.MsgOptionPostEphemeral(ev.User), slack.MsgOptionText(fmt.Sprintf("上传图片失败: %s\n", err), false))
		return
	}
}

func processBuildCommand(ctx context.Context, ev *slackevents.AppMentionEvent, cmd []string) {
	tag := cmd[1]
	clean := len(cmd) > 2 && cmd[2] == " clean"

	if err := triggerUnityBuild(ctx, tag, clean, ev.Channel); err != nil {
		api.PostMessage(ev.Channel, slack.MsgOptionText(err.Error(), false))
	}
}

func processPingCommand(ctx context.Context, ev *slackevents.AppMentionEvent, cmd []string) {
	msg := "pong"
	if len(cmd) > 1 && cmd[1] != "" {
		msg = cmd[1]
	}
	api.PostMessage(ev.Channel, slack.MsgOptionText(msg, false))

	attachment := slack.Attachment{
		Pretext:    "",
		Fallback:   "",
		CallbackID: "accept_or_reject",
		Color:      "#3AA3E3",
		Actions: []slack.AttachmentAction{
			slack.AttachmentAction{
				Name:  "accept",
				Text:  "Accept",
				Type:  "button",
				Value: "accept",
			},
			slack.AttachmentAction{
				Name:  "reject",
				Text:  "Reject",
				Type:  "button",
				Value: "reject",
				Style: "danger",
			},
		},
	}

	message := slack.MsgOptionAttachments(attachment)
	api.PostMessage(ev.Channel, slack.MsgOptionText(msg, false), message)
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

func triggerUnityCancel(ctx context.Context, tag string, buildNumber string, action slack.InteractionCallback) error {
	client := unityClient()

	result, _, err := client.BuildsApi.CancelBuild(ctx, UNITY_ORG, UNITY_PROJECT, tag, buildNumber)

	if err != nil || strings.TrimSpace(result) != "" {
		fmt.Printf("取消result: %s error: %s\n", result, err)
		api.PostMessage(action.Channel.ID, slack.MsgOptionPostEphemeral(action.User.ID), slack.MsgOptionText(fmt.Sprintf("调用unity接口出错: result: %s error: %s", result, err.Error()), false))

		return errors.Wrap(err, "调用unity接口出错")
	}

	emptySlice := make([]slack.Attachment, 0)
	api.UpdateMessageContext(ctx, action.Channel.ID, action.OriginalMessage.Timestamp, slack.MsgOptionText(fmt.Sprintf("%s %s 已取消 by @%s", tag, buildNumber, action.User.Name), false), slack.MsgOptionAttachments(emptySlice...))
	return nil
}

func triggerUnityBuild(ctx context.Context, tag string, clean bool, slackChannel string) error {
	client := unityClient()

	option := &swagger.StartBuildsOpts{
		Options: optional.NewInterface(swagger.InlineObject9{
			Clean: clean,
			Delay: 5,
		}),
	}

	builds, _, err := client.BuildsApi.StartBuilds(ctx, UNITY_ORG, UNITY_PROJECT, tag, option)

	if err != nil {
		return errors.Wrap(err, "调用unity接口出错")
	}

	if payload := builds; payload == nil || len(payload) == 0 {
		return errors.New("unity没有返回错误, 也没有返回payload...")
	} else {
		fmt.Printf("收到%d个payload\n", len(payload))

		for i, p := range payload {
			fmt.Printf("payload[%d] build: %v, error: %s\n", i, p.Build, p.Error)
			if i > 0 {
				continue
			}
			if err := p.Error; err != "" {
				if strings.Contains(err, "already a build pending") {
					return nil
				}

				return errors.New(fmt.Sprintf("unity返回错误: build: %v: %s", p.Build, err))
			}

			attachment := slack.Attachment{
				Pretext:    fmt.Sprintf("启动成功: %s %v", tag, p.Build),
				Fallback:   fmt.Sprintf("启动成功: %s %v", tag, p.Build),
				CallbackID: "cancel_build",
				Color:      "#3AA3E3",
				Actions: []slack.AttachmentAction{
					slack.AttachmentAction{
						Name:  "cancel",
						Text:  "取消",
						Type:  "button",
						Value: fmt.Sprintf("%v_%s", p.Build, tag),
						Confirm: &slack.ConfirmationField{
							Title:       "确认",
							Text:        "确定要取消构建吗?",
							OkText:      "Yes",
							DismissText: "No",
						},
					},
				},
			}

			message := slack.MsgOptionAttachments(attachment)
			api.PostMessage(slackChannel, message)
		}

		return nil
	}

	return errors.New("unity返回没处理?")
}
