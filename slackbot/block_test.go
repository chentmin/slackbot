package slackbot

import "testing"

func TestBlockMessage(t *testing.T) {
	msg := `[
	{
		"type": "section",
		"text": {
			"type": "mrkdwn",
			"text": "You have a new request:\n*<fakeLink.toEmployeeProfile.com|Fred Enriquez - New device request>*"
		}
	},
	{
		"type": "section",
		"fields": [
			{
				"type": "mrkdwn",
				"text": "*Type:*\n{{.Type}}"
			},
			{
				"type": "mrkdwn",
				"text": "*When:*\n{{.When}}"
			}
		]
	},
	{
		"type": "actions",
		"elements": [
			{
				"type": "button",
				"text": {
					"type": "plain_text",
					"emoji": true,
					"text": "Approve"
				},
				"confirm": {
					"title": {
						"type": "plain_text",
						"text": "Are you sure?"
					},
					"text": {
						"type": "mrkdwn",
						"text": "Wouldn't you prefer a good game of _chess_?"
					},
					"confirm": {
						"type": "plain_text",
						"text": "Do it"
					},
					"deny": {
						"type": "plain_text",
						"text": "Stop, I've changed my mind!"
					}
				},
				"style": "primary",
				"value": "click_me_123"
			}
		]
	}
]`
	result, err := NewBlockMessage(msg, map[string]interface{}{"Type": "Computer", "When": "Today"})

	if err != nil{
		t.Error(err)
	}

	if result == nil{
		t.Fail()
	}
}

