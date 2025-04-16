package alert

import (
	"context"
	"darkhawk/app/biz/service"
	"darkhawk/pkg/response"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
)

func Mtta(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	data, err := service.NewAlertStatisticService(ctx, c).MTTA(id)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, data)
}

func AllMtta(ctx context.Context, c *app.RequestContext) {
	data, err := service.NewAlertStatisticService(ctx, c).MTTA(0)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, data)
}

func Mttr(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	data, err := service.NewAlertStatisticService(ctx, c).MTTR(id)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, data)
}

func AllMttr(ctx context.Context, c *app.RequestContext) {
	data, err := service.NewAlertStatisticService(ctx, c).MTTR(0)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, data)
}

func Count(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	data, err := service.NewAlertStatisticService(ctx, c).Count(id)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, data)
}

func AllCount(ctx context.Context, c *app.RequestContext) {
	data, err := service.NewAlertStatisticService(ctx, c).Count(0)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, data)
}

func Level(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	data, err := service.NewAlertStatisticService(ctx, c).Level(id)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, data)
}

func AllLevel(ctx context.Context, c *app.RequestContext) {
	data, err := service.NewAlertStatisticService(ctx, c).Level(0)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, data)
}

func TopTen(ctx context.Context, c *app.RequestContext) {
	timeRange := c.Query("timeRange")
	data, err := service.NewAlertStatisticService(ctx, c).TopTen(timeRange)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, data)
}

func LevelAlert(ctx context.Context, c *app.RequestContext) {
	timeRange := c.Query("timeRange")
	data, err := service.NewAlertStatisticService(ctx, c).LevelAlert(timeRange)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, data)
}

func AlertCount(ctx context.Context, c *app.RequestContext) {
	timeRange := c.Query("timeRange")
	data, err := service.NewAlertStatisticService(ctx, c).AlertCount(timeRange)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, data)
}
