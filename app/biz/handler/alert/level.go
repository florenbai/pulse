package alert

import (
	"context"
	"darkhawk/app/biz/model"
	"darkhawk/app/biz/service"
	pb "darkhawk/hertz_gen/alert_level"
	"darkhawk/pkg/response"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
)

func CreateLevel(ctx context.Context, c *app.RequestContext) {
	var req pb.LevelRequest
	if err := c.Bind(&req); err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err := service.NewAlertLevelService(ctx, c).CreateAlertLevel(req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, response.Success)
}

func UpdateLevel(ctx context.Context, c *app.RequestContext) {
	var req pb.LevelRequest
	if err := c.Bind(&req); err != nil {
		response.SendBaseResp(c, err)
		return
	}
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err = service.NewAlertLevelService(ctx, c).UpdateAlertLevel(id, req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, response.Success)
}

func DeleteLevel(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err = service.NewAlertLevelService(ctx, c).DeleteAlertLevel(id)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, response.Success)
}

func LevelAll(ctx context.Context, c *app.RequestContext) {
	var err error
	alertLevelModel := new(model.AlertLevel)
	data, err := alertLevelModel.All(ctx)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, data)
}

func LevelMappingAll(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	levelIntegrationMapping := new(model.LevelIntegrationMapping)
	data, err := levelIntegrationMapping.All(ctx, id)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, data)
}

func UpdateLevelMapping(ctx context.Context, c *app.RequestContext) {
	var req []model.LevelIntegrationMapping
	if err := c.Bind(&req); err != nil {
		response.SendBaseResp(c, err)
		return
	}
	levelIntegrationMapping := new(model.LevelIntegrationMapping)
	err := levelIntegrationMapping.UpdateMapping(ctx, req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, response.Success)
}
