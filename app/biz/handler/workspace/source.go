package workspace

import (
	"context"
	"darkhawk/app/biz/service"
	pb "darkhawk/hertz_gen/workspace_source"
	"darkhawk/pkg/response"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
)

func RelatedAlertSource(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	var req pb.RelatedWorkspaceSourceRequest
	if err := c.Bind(&req); err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err = service.NewWorkspaceSourceService(ctx, c).RelatedAlertSource(id, req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, nil)
}

func UpdateAlertSource(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	var req pb.EditWorkspaceSourceRequest
	if err := c.Bind(&req); err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err = service.NewWorkspaceSourceService(ctx, c).UpdateAlertSource(id, req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, nil)
}

func RefreshSourceToken(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	token, err := service.NewWorkspaceSourceService(ctx, c).RefreshToken(id)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, token)
}

func LevelMapping(ctx context.Context, c *app.RequestContext) {
	var req pb.LevelMappingRequest
	if err := c.Bind(&req); err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err := service.NewWorkspaceSourceService(ctx, c).UpdateLevelMapping(req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, nil)
}
