package alert

import (
	"context"
	"darkhawk/app/biz/model"
	"darkhawk/app/biz/service"
	pb "darkhawk/hertz_gen/alert_strategy"
	"darkhawk/pkg/response"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
)

func CreateStrategy(ctx context.Context, c *app.RequestContext) {
	var req pb.StrategyRequest
	if err := c.Bind(&req); err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err := service.NewAlertStrategyService(ctx, c).CreateAlertStrategy(req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, response.Success)
}

func EditStrategy(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	var req pb.StrategyRequest
	if err := c.Bind(&req); err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err = service.NewAlertStrategyService(ctx, c).UpdateAlertStrategy(id, req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, response.Success)
}

func StrategyAll(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	strategyModel := new(model.AlertStrategy)
	data, err := strategyModel.WorkspaceAll(ctx, id)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, data)
}

func ChangeStrategyStatus(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err = service.NewAlertStrategyService(ctx, c).ChangeAlertStrategyStatus(id)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, response.Success)
}

func ChangeStrategySort(ctx context.Context, c *app.RequestContext) {
	var req pb.ChangeStrategySortRequest
	if err := c.Bind(&req); err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err := service.NewAlertStrategyService(ctx, c).ChangeSort(req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, response.Success)
}

func DeleteStrategy(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err = service.NewAlertStrategyService(ctx, c).DeleteStrategy(id)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, response.Success)
}

func CreateSystemStrategy(ctx context.Context, c *app.RequestContext) {
	var req pb.SystemStrategyRequest
	if err := c.Bind(&req); err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err := service.NewAlertStrategyService(ctx, c).CreateSystemStrategy(req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, response.Success)
}

func EditSystemStrategy(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	var req pb.SystemStrategyRequest
	if err := c.Bind(&req); err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err = service.NewAlertStrategyService(ctx, c).UpdateSystemStrategy(id, req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, response.Success)
}

func SystemStrategyList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req pb.SystemStrategyQuery
	if err = c.Bind(&req); err != nil {
		response.SendBaseResp(c, err)
		return
	}
	resp, err := service.NewAlertStrategyService(ctx, c).GetSystemStrategyList(req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, resp)
}

func SystemStrategyAll(ctx context.Context, c *app.RequestContext) {
	var err error
	resp, err := service.NewAlertStrategyService(ctx, c).GetSystemStrategyAll()
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, resp)
}

func DeleteSystemStrategy(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err = service.NewAlertStrategyService(ctx, c).DeleteSystemStrategy(id)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, response.Success)
}

func SystemStrategyInfoByWorkspace(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	data, err := service.NewAlertStrategyService(ctx, c).GetSystemStrategyInfoByWorkspace(id)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, data)
}

func StrategyLog(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	data, err := service.NewAlertLogService(ctx, c).GetStrategyLog(id)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, data)
}
