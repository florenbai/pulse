package umdp

import (
	"darkhawk/conf"
	"encoding/json"
	"github.com/go-resty/resty/v2"
)

func SendEmail(receiver []string, content string) error {
	if len(receiver) < 1 {
		return nil
	}
	ma := map[string]interface{}{
		"templateId": conf.GetConf().Umdp.Template,
		"channel":    conf.GetConf().Umdp.EmailChannel,
		"parameters": map[string]interface{}{
			"receiver": receiver,
			"cc":       []string{},
			"variable": map[string]interface{}{
				"content": content,
			},
		},
	}
	aa, _ := json.Marshal(ma)
	client := resty.New()
	_, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("token", conf.GetConf().Umdp.Token).
		SetBody(aa).
		Post(conf.GetConf().Umdp.Host)
	if err != nil {
		return err
	}
	return err
}
