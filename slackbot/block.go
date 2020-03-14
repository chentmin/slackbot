package slackbot

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/slack-go/slack"
	"text/template"
)

func NewBlockMessage(block string, params map[string]interface{}) (*slack.Blocks, error){
	tplt, err := template.New("n").Parse(block)
	if err != nil{
		return nil, errors.Wrap(err, "解析block失败")
	}

	buf := &bytes.Buffer{}
	if err :=tplt.Execute(buf, params); err != nil{
		return nil, errors.Wrap(err, "template执行失败")
	}

	result := &slack.Blocks{}

	if err := json.Unmarshal(buf.Bytes(), result); err != nil{
		return nil, errors.Wrap(err, "json解析为blocks出错")
	}

	return result, nil
}