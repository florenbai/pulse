package notify

import (
	"context"
	"darkhawk/app/biz/model"
	"darkhawk/pkg/umdp"
	"darkhawk/pkg/utils"
	"darkhawk/pkg/zlog"
	"fmt"
	"github.com/fatih/structs"
	"github.com/flosch/pongo2/v6"
	"math"
	"time"
)

// Send 发送消息
func (n *Notify) Send(ctx context.Context, templateId uint64, members []uint64, channels []uint64, event model.AlertEvent) error {

	// 获取消息模板信息
	channelTemplateModel := new(model.ChannelTemplate)
	channelDetail, err := channelTemplateModel.GetChannelTemplateDetail(ctx, templateId)
	if err != nil {
		return err
	}
	userModel := new(model.User)
	alertEvent, err := event.GetAlertEventDetail(ctx, event.Id)
	if err != nil {
		return err
	}

	for _, v := range channelDetail {
		if utils.InOfT(v.ChannelId, channels) && v.Finished {
			tmpl, err := pongo2.FromString(v.Content)
			if err != nil {
				zlog.Error(fmt.Sprintf("渠道：【%s】模板初始化异常，%s", v.ChannelName, err.Error()))
				continue
			}
			m := structs.Map(alertEvent)
			content, err := tmpl.Execute(m)
			if err != nil {
				zlog.Error(err.Error())
				continue
			}
			switch v.ChannelType {
			case "企业微信":
				wx, err := userModel.GetUserWxIds(ctx, members)
				if err != nil {
					zlog.Error(err.Error())
					continue
				}
				err = umdp.SendWechat(wx, content)
				if err != nil {
					zlog.Error(err.Error())
					continue
				}
			case "邮件":
				emails, err := userModel.GetUserEmails(ctx, members)
				if err != nil {
					zlog.Error(err.Error())
					continue
				}

				err = umdp.SendEmail(emails, content)
				if err != nil {
					zlog.Error(err.Error())
					continue
				}
			case "电话":
				err := model.GetOneById(context.TODO(), event.TableName(), event.Id, &event)
				if err != nil {
					zlog.Error(fmt.Sprintf("电话，获取告警信息异常：%s", err.Error()))
					continue
				}
				if event.Progress == model.EventUnClaimed {
					phones, err := userModel.GetUserPhone(ctx, members)
					if err != nil {
						zlog.Error(err.Error())
						continue
					}
					zlog.Debug(fmt.Sprintf("向%v发送电话语音，告警编号：%d", phones, event.Id))
					endTime := time.Now()
					startTime := event.FirstTriggerTime.ToTime()
					duration := endTime.Sub(startTime)
					minutes := duration.Minutes()
					err = umdp.SendPhone(phones, event.AlertTitle, fmt.Sprintf("%d", event.NotifyCurNumber), fmt.Sprintf("%d分钟", int(math.Round(minutes))))
					if err != nil {
						zlog.Error(err.Error())
						continue
					}
				}
			case "企微群机器人":
				err = umdp.SendWechatBot([]string{v.ChannelSign}, content)
				if err != nil {
					zlog.Error(err.Error())
					continue
				}
			}
		}
	}

	return nil
}
