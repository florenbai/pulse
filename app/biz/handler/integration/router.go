package integration

import (
	"context"
	"darkhawk/app/biz/service"
	pb "darkhawk/hertz_gen/integration"
	"darkhawk/pkg/response"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
)

func CreateRouter(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	var req pb.IntegrationRouterRequest
	if err := c.Bind(&req); err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err = service.NewSystemIntegrationService(ctx, c).CreateIntegrationRouter(id, req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, nil)
}

func UpdateRouter(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	var req pb.IntegrationRouterRequest
	if err := c.Bind(&req); err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err = service.NewSystemIntegrationService(ctx, c).UpdateIntegrationRouter(id, req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, nil)
}

func GetRouters(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	data, err := service.NewSystemIntegrationService(ctx, c).GetIntegrationRouterList(id)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, data)
}

func DeleteRouter(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err = service.NewSystemIntegrationService(ctx, c).DeleteIntegrationRouter(id)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}

	response.SendBaseResp(c, nil)
}

func ChangeSort(ctx context.Context, c *app.RequestContext) {
	var req []uint64
	err := c.BindJSON(&req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err = service.NewSystemIntegrationService(ctx, c).ChangeRouterSort(req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, nil)
}
