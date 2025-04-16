package service

import (
	"context"
	"darkhawk/app/biz/model"
	"darkhawk/conf"
	"darkhawk/pkg/wework"
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
)

type WechatService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewWechatService(ctx context.Context, c *app.RequestContext) *WechatService {
	return &WechatService{
		ctx: ctx,
		c:   c,
	}
}

func (service *WechatService) GetWxUserInfo(code string) (string, error) {
	wxConfig := wework.WeWorkConfig{
		CorpId:              "ww02f67573112243a9",
		SuiteToken:          "Jcq7jOGgoBWJUUbAxJx5U2YWkbGwlxxXmTI9HZgOB3Y",
		SuiteEncodingAesKey: "",
		Dsn:                 "",
	}
	client := wework.NewWeWork(wxConfig)
	client.SetDebug(true)
	client.SetCache(wework.RedisOpts{
		Host:     conf.GetConf().Redis.Address,
		Password: conf.GetConf().Redis.Password,
		Database: conf.GetConf().Redis.Db,
		MaxIdle:  conf.GetConf().Redis.MaxIdle,
	})
	client.SetAppSecretFunc(func(corpId uint) (corpid string, secret string, customizedApp bool) {
		return wxConfig.CorpId, wxConfig.SuiteToken, true
	})
	resp := client.GetUserInfo(0, code)

	if resp.Error.ErrorMsg != "ok" {
		return "", fmt.Errorf(resp.Error.ErrorMsg)
	}
	if resp.UserId == "" {
		return "", errors.New("您没有访问权限")
	}
	userModel := new(model.User)
	ok, err := model.ExistName(service.ctx, userModel.TableName(), "userid", resp.UserId, nil)
	if err != nil {
		return "", err
	}
	if !ok {
		return "", errors.New("您没有访问权限")
	}
	return resp.UserId, nil
}
