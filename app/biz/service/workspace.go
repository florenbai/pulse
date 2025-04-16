package service

import (
	"context"
	"darkhawk/app/biz/model"
	"darkhawk/hertz_gen/base"
	pb "darkhawk/hertz_gen/workspace"
	"darkhawk/pkg/response"
	"darkhawk/pkg/utils"
	"encoding/json"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
	"strconv"
	"strings"
)

type WorkspaceService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewWorkspaceService(ctx context.Context, c *app.RequestContext) *WorkspaceService {
	return &WorkspaceService{
		ctx: ctx,
		c:   c,
	}
}

// CreateWorkspace 创建工作区
func (service *WorkspaceService) CreateWorkspace(req pb.WorkspaceRequest) error {
	workspaceModel := new(model.Workspace)
	ok, err := model.ExistName(service.ctx, workspaceModel.TableName(), "workspace_name", req.GetWorkspaceName(), nil)
	if err != nil {
		return err
	}
	if ok {
		return response.WorkspaceExistErr
	}
	if err = utils.AnyToAny(req, &workspaceModel); err != nil {
		return err
	}
	workspaceModel.SystemStrategy = req.GetStrategy()
	return workspaceModel.Create(service.ctx, req.GetTeams())
}

// UpdateWorkspace 编辑工作区
func (service *WorkspaceService) UpdateWorkspace(id uint64, req pb.WorkspaceBaseResp) error {
	workspaceModel := new(model.Workspace)
	err := model.GetOneById(service.ctx, workspaceModel.TableName(), id, &workspaceModel)
	if err != nil {
		return err
	}
	ok, err := model.ExistName(service.ctx, workspaceModel.TableName(), "workspace_name", req.GetWorkspaceName(), &id)
	if err != nil {
		return err
	}
	if ok {
		return response.WorkspaceExistErr
	}
	workspaceModel.WorkspaceName = req.GetWorkspaceName()
	workspaceModel.WorkspaceDesc = req.GetWorkspaceDesc()
	return workspaceModel.Update(service.ctx, req.GetTeams())
}

// GetWorkspaceList 获取工作区列表
func (service *WorkspaceService) GetWorkspaceList(req pb.WorkspaceQuery) (*base.BaseListResp, error) {
	search := map[string]string{
		"workspace_name": req.GetWorkspaceName(),
	}
	associate := map[string]string{
		"workspace_name": "workspace_name LIKE ?",
	}

	session := sessions.Default(service.c)
	uid := session.Get("uid").(uint64)
	teamUserModel := new(model.TeamUser)
	ok := teamUserModel.IsAllWorkspace(service.ctx, uid)

	container := model.NewScopeContainer()
	if !ok {
		workspaces, err := teamUserModel.GetUserWorkspace(service.ctx, uid)
		if err != nil {
			return nil, err
		}
		container = append(container, model.InWithScope("id", workspaces))
	}

	scopes := model.ParamWithScope(search, associate, nil, false)
	container = append(container, scopes)
	workspaceModel := new(model.Workspace)
	data, i, err := workspaceModel.List(service.ctx, req.GetPage(), req.GetPageSize(), container...)
	if err != nil {
		return nil, err
	}
	listValue := new(structpb.ListValue)
	b, _ := json.Marshal(data)
	err = protojson.Unmarshal(b, listValue)
	if err != nil {
		return nil, err
	}
	return &base.BaseListResp{
		Total: uint64(i),
		List:  listValue,
	}, nil
}

// GetWorkspaceBaseInfo 工作区基础信息
func (service *WorkspaceService) GetWorkspaceBaseInfo(id uint64) (*pb.WorkspaceBaseResp, error) {
	workspaceModel := new(model.Workspace)
	bl, err := workspaceModel.BaseInfo(service.ctx, id)
	if err != nil {
		return nil, err
	}
	var teamsArr []uint64
	teams := strings.Split(bl.Teams, ",")
	for _, v := range teams {
		i, _ := strconv.Atoi(v)
		teamsArr = append(teamsArr, uint64(i))
	}
	return &pb.WorkspaceBaseResp{
		WorkspaceName: bl.WorkspaceName,
		WorkspaceDesc: bl.WorkspaceDesc,
		Teams:         teamsArr,
	}, nil
}

// DeleteWorkspace 删除工作区
func (service *WorkspaceService) DeleteWorkspace(id uint64) error {
	workspaceModel := new(model.Workspace)
	err := model.GetOneById(service.ctx, workspaceModel.TableName(), id, &workspaceModel)
	if err != nil {
		return err
	}
	return workspaceModel.Delete(service.ctx)
}

// ChangeStatus 改变状态
func (service *WorkspaceService) ChangeStatus(id uint64) error {
	workspaceModel := new(model.Workspace)
	err := model.GetOneById(service.ctx, workspaceModel.TableName(), id, &workspaceModel)
	if err != nil {
		return err
	}
	var data = make(map[string]interface{})
	if workspaceModel.Enable {
		data["enable"] = false
	} else {
		data["enable"] = true
	}
	return model.EditOneById(service.ctx, workspaceModel.TableName(), id, data)
}

// GetWorkspaceAll 获取所有工作区
func (service *WorkspaceService) GetWorkspaceAll() ([]model.Workspace, error) {
	var data []model.Workspace
	workspaceModel := new(model.Workspace)
	err := model.GetAll(service.ctx, workspaceModel.TableName(), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
