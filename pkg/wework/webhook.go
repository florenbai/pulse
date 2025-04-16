package wework

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
)

type BotTextMessage struct {
	Text Text `json:"text" validate:"required"`
}

type BotMarkDownMessage struct {
	MarkDown Text `json:"markdown" validate:"required"`
}

// BotMessageSend 发送群机器人消息
// https://developer.work.weixin.qq.com/document/path/99110
func (ww *weWork) BotMessageSend(key string, msg interface{}) (resp MessageSendResponse) {
	if ok := validate.Validate(msg); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	h := H{}
	buf, _ := json.Marshal(msg)
	json.Unmarshal(buf, &h)
	switch msg.(type) {
	case BotTextMessage:
		h["msgtype"] = "text"
	case BotMarkDownMessage:
		h["msgtype"] = "markdown"
	}
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%s", key)
	_, err := resty.New().R().SetBody(h).SetResult(&resp).
		Post(url)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}
