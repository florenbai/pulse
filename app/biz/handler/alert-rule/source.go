package alert_rule

import (
	"context"
	"darkhawk/app/biz/service"
	pb "darkhawk/hertz_gen/alert_rule"
	basepb "darkhawk/hertz_gen/base"
	"darkhawk/pkg/response"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
)

func CreateSource(ctx context.Context, c *app.RequestContext) {
	var req pb.SourceRequest
	if err := c.Bind(&req); err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err := service.NewAlertRuleSourceService(ctx, c).CreateSource(req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, response.Success)
}

func UpdateSource(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	var req pb.SourceRequest
	if err := c.Bind(&req); err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err = service.NewAlertRuleSourceService(ctx, c).UpdateSource(id, req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, response.Success)
}

func List(ctx context.Context, c *app.RequestContext) {
	var err error
	var req basepb.PageInfoReq
	if err = c.Bind(&req); err != nil {
		response.SendBaseResp(c, err)
		return
	}

	resp, err := service.NewAlertRuleSourceService(ctx, c).GetSourceList(req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}

	response.SendDataResp(c, response.Success, resp)
}

func AllSource(ctx context.Context, c *app.RequestContext) {
	resp, err := service.NewAlertRuleSourceService(ctx, c).GetAllSource()
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}

	response.SendDataResp(c, response.Success, resp)
}

func DeleteSource(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err = service.NewAlertRuleSourceService(ctx, c).DeleteAlertRuleSource(id)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}

	response.SendBaseResp(c, response.Success)
}

func AgentList(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	data, err := service.NewAlertRuleSourceService(ctx, c).GetSourceAgent(id)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, data)
}

func Reload(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err = service.NewAlertRuleSourceService(ctx, c).ReloadSource(id)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, response.Success)
}
