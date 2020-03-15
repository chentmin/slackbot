package slackbot

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/chentmin/once"
	"github.com/gin-gonic/gin"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

var (
	commandMap map[string]Command
)

type Command func(ctx context.Context, ev *slackevents.AppMentionEvent, cmd []string)

type Callback func(ctx context.Context, action slack.InteractionCallback)

type Manager struct {
	slackToken             string
	slackVerificationToken string

	commandMap  map[string]Command
	callbackMap map[string]Callback

	onceDynamoTable string
}

func New(token, verificationToken string, options ...option) *Manager {
	result := &Manager{
		slackToken:             token,
		slackVerificationToken: verificationToken,
		commandMap:             make(map[string]Command),
		callbackMap:            make(map[string]Callback),
	}

	for _, ops := range options {
		ops(result)
	}

	return result
}

type option func(*Manager)

func OnlyOnceByDynamoDB(table string) option {
	return func(manager *Manager) {
		manager.onceDynamoTable = table
	}
}

func (m *Manager) RegisterMentionCommand(reg string, cmd Command) {
	if _, has := m.commandMap[reg]; has {
		panic("重复注册了command: " + reg)
	}
	m.commandMap[reg] = cmd
}

func (m *Manager) RegisterCallback(reg string, callback Callback) {
	if _, has := m.callbackMap[reg]; has {
		panic("重复注册了callback: " + reg)
	}
	m.callbackMap[reg] = callback
}

func (m *Manager) HandleMessageEvent(c *gin.Context) {

	body, _ := ioutil.ReadAll(c.Request.Body)

	eventsAPIEvent, e := slackevents.ParseEvent(json.RawMessage(body), slackevents.OptionVerifyToken(slackevents.TokenComparator{VerificationToken: m.slackVerificationToken}))
	if e != nil {
		fmt.Printf("收到request, 但是作为event解析失败: %s\n", body)
		c.String(http.StatusBadRequest, "")
		return
	}

	if eventsAPIEvent.Type == slackevents.URLVerification {
		var r *slackevents.ChallengeResponse
		err := json.Unmarshal([]byte(body), &r)
		if err != nil {
			return
		}

		c.String(http.StatusOK, r.Challenge)
		return
	}

	defer func() {
		c.String(http.StatusOK, "")
	}()

	if eventsAPIEvent.Type == slackevents.CallbackEvent {
		innerEvent := eventsAPIEvent.InnerEvent

		fmt.Printf("inner event type: %s\n", innerEvent.Type)

		switch innerEvent.Type {
		case slackevents.AppMention:
			ev := innerEvent.Data.(*slackevents.AppMentionEvent)
			if m.onceDynamoTable != "" {
				if err := once.New(m.onceDynamoTable).Ensure(ev.TimeStamp); err != nil {
					return
				}
			}

			fmt.Printf("收到mention事件: %s: %s\n", ev.User, ev.Text)

			text := strings.TrimSpace(ev.Text)

			processed := false
			for cmd, process := range m.commandMap {
				if reg := regexp.MustCompile(cmd); reg.MatchString(text) {
					param := reg.FindStringSubmatch(text)
					process(c, ev, param)
					processed = true
					break
				}
			}

			if !processed {
				fmt.Printf("收到未知的命令: %s\n", text)
			}

			return

		default:
			fmt.Printf("收到未处理的event type: %s\n", innerEvent.Type)
		}
	}

	fmt.Printf("收到未知的事件: %s\n", body)

	return
}

func (m *Manager) HandleCallbackEvent(c *gin.Context) {
	defer func() {
		c.String(http.StatusOK, "")
	}()

	payload, has := c.GetPostForm("payload")

	if !has {
		fmt.Printf("callbackk没有payload的formValue\n")
		return
	}

	var action slack.InteractionCallback

	if err := json.Unmarshal([]byte(payload), &action); err != nil {
		fmt.Printf("unmarshal payload失败: %s\n", err)
		return
	}

	fmt.Printf("payload: %+v\n", action)

	if action.Token != m.slackVerificationToken {
		fmt.Printf("token验证失败\n")
		return
	}

	if m.onceDynamoTable != "" {
		if err := once.New(m.onceDynamoTable).Ensure(action.ActionTs); err != nil {
			return
		}
	}

	fmt.Printf("收到callback事件: %s: %s\n", action.User.Name, action.CallbackID)

	processed := false

	for callback, cmd := range m.callbackMap {
		if callback == action.CallbackID {
			cmd(c, action)
			processed = true
			break
		}
	}

	if !processed {
		fmt.Printf("未知callback id: %s\n", action.CallbackID)
	}

	return
}
