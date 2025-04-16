package wechat

import (
	"context"
	"darkhawk/app/biz/service"
	pb "darkhawk/hertz_gen/alert_event"
	"darkhawk/pkg/response"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
)

func AlertList(ctx context.Context, c *app.RequestContext) {
	req := pb.WorkspaceEventListRequest{
		AlertTitle: "",
		Progress:   4,
		TimeRange:  "6month",
		AlertLevel: nil,
		Page:       1,
		PageSize:   100,
	}
	data, err := service.NewAlertEventService(ctx, c).GetEventList(req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, data)
}

func WxUserInfo(ctx context.Context, c *app.RequestContext) {
	code := c.Query("code")
	if code == "" {
		response.SendBaseResp(c, errors.New("code 不能为空"))
		return
	}
	userid, err := service.NewWechatService(ctx, c).GetWxUserInfo(code)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, userid)
}
