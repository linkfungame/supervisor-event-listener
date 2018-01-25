package notify

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ouqiang/supervisor-event-listener/event"
	"github.com/ouqiang/supervisor-event-listener/utils/httpclient"
)

type DingTalk struct{}

type DingTalkMsg struct {
	MsgType string `json:"msgtype"`
	Text struct {
		Content string `json:"content"`
	} `json:"text"`
	At struct {
		IsAtAll bool `json:"isAtAll"`
	} `json:"at"`
}

func (d *DingTalkMsg) fromMessage(message event.Message) {
	d.MsgType = "text"
	d.At.IsAtAll = true
	d.Text.Content = message.String()
}

func (hook *DingTalk) Send(message event.Message) error {
	d := new(DingTalkMsg)
	d.fromMessage(message)
	encodeMessage, err := json.Marshal(d)
	if err != nil {
		return err
	}
	timeout := 60
	response := httpclient.PostJson(Conf.WebHook.Url, string(encodeMessage), timeout)

	if response.StatusCode == 200 {
		return nil
	}
	errorMessage := fmt.Sprintf("webhook执行失败#HTTP状态码-%d#HTTP-BODY-%s", response.StatusCode, response.Body)
	return errors.New(errorMessage)
}
