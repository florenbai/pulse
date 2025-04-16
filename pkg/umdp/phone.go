package umdp

import (
	"darkhawk/conf"
	"darkhawk/pkg/zlog"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
)

func SendPhone(receiver []string, alertTitle, alertCount, duration string) error {
	if len(receiver) < 1 {
		return nil
	}
	for _, u := range receiver {
		ma := map[string]interface{}{
			"templateId": conf.GetConf().Umdp.Template,
			"channel":    conf.GetConf().Umdp.PhoneChannel,
			"parameters": map[string]interface{}{
				"receiver": []string{u},
				"cc":       []string{},
				"variable": map[string]interface{}{
					"alertTitle": alertTitle,
					"alertCount": alertCount,
					"duration":   duration,
				},
			},
		}

		aa, _ := json.Marshal(ma)
		client := resty.New()
		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetHeader("token", conf.GetConf().Umdp.Token).
			SetBody(aa).
			Post(conf.GetConf().Umdp.Host)
		if err != nil {
			return err
		}
		zlog.Debug(fmt.Sprintf("发送电话通知resp: %s", resp.String()))
	}
	return nil
}
