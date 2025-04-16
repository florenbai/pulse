package workspace

import (
	"context"
	"darkhawk/app/biz/model"
	"darkhawk/app/biz/service"
	"darkhawk/hertz_gen/workspace"
	"darkhawk/pkg/response"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
)

// CreateWorkspace 创建工作区
func CreateWorkspace(ctx context.Context, c *app.RequestContext) {
	var req workspace.WorkspaceRequest
	if err := c.Bind(&req); err != nil {
		response.SendBaseResp(c, err)
		return
	}

	err := service.NewWorkspaceService(ctx, c).CreateWorkspace(req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, nil)
	return
}

// UpdateBaseInfo 编辑工作区
func UpdateBaseInfo(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	var req workspace.WorkspaceBaseResp
	if err := c.Bind(&req); err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err = service.NewWorkspaceService(ctx, c).UpdateWorkspace(id, req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, nil)
}

// WorkspaceList .
func WorkspaceList(ctx context.Context, c *app.RequestContext) {
	var req workspace.WorkspaceQuery
	if err := c.Bind(&req); err != nil {
		response.SendBaseResp(c, err)
		return
	}

	data, err := service.NewWorkspaceService(ctx, c).GetWorkspaceList(req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, data)
}

func WorkspaceSource(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	wsModel := new(model.WorkspaceSource)
	data, err := wsModel.GetWorkspaceSource(ctx, id)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, data)
}

func DeleteSource(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	wsModel := new(model.WorkspaceSource)
	err = wsModel.DeleteWorkspaceSource(ctx, id)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, nil)
}

func BaseInfo(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	resp, err := service.NewWorkspaceService(ctx, c).GetWorkspaceBaseInfo(id)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, resp)
}

func DeleteWorkspace(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err = service.NewWorkspaceService(ctx, c).DeleteWorkspace(id)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, nil)
}

func ChangeStatus(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err = service.NewWorkspaceService(ctx, c).ChangeStatus(id)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, nil)
}

func All(ctx context.Context, c *app.RequestContext) {
	resp, err := service.NewWorkspaceService(ctx, c).GetWorkspaceAll()
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, resp)
}
