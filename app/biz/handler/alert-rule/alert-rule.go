package alert_rule

import (
	"context"
	"darkhawk/app/biz/service"
	pb "darkhawk/hertz_gen/alert_rule"
	"darkhawk/pkg/response"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
)

func Create(ctx context.Context, c *app.RequestContext) {
	var req pb.RuleRequest
	if err := c.Bind(&req); err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err := service.NewAlertRuleService(ctx, c).CreateRule(req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, response.Success)
}

func Update(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	var req pb.RuleRequest
	if err := c.Bind(&req); err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err = service.NewAlertRuleService(ctx, c).UpdateRule(id, req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, response.Success)
}

func DeleteRule(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err = service.NewAlertRuleService(ctx, c).DeleteRule(id)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, response.Success)
}

func RuleGroups(ctx context.Context, c *app.RequestContext) {
	data, err := service.NewAlertRuleService(ctx, c).GetRuleGroup()
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, data)
}

func CreateRuleGroup(ctx context.Context, c *app.RequestContext) {
	var req pb.RuleGroupRequest
	if err := c.Bind(&req); err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err := service.NewAlertRuleService(ctx, c).CreateRuleGroup(req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, nil)
}

func RuleList(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	data, err := service.NewAlertRuleService(ctx, c).GetRuleListByGroup(id)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, data)
}

func DeleteRuleGroup(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err = service.NewAlertRuleService(ctx, c).DeleteRuleGroup(id)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, response.Success)
}

func GetLabelsByRuleGroup(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	data, err := service.NewAlertRuleService(ctx, c).GetLabelsByRuleGroup(id)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, data)
}
