package service

import (
	"context"
	"darkhawk/app/biz/model"
	pb "darkhawk/hertz_gen/workspace_source"
	"darkhawk/pkg/response"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/google/uuid"
)

type WorkspaceSourceService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewWorkspaceSourceService(ctx context.Context, c *app.RequestContext) *WorkspaceSourceService {
	return &WorkspaceSourceService{
		ctx: ctx,
		c:   c,
	}
}

func (service *WorkspaceSourceService) RelatedAlertSource(id uint64, req pb.RelatedWorkspaceSourceRequest) error {
	workspaceSourceModel := new(model.WorkspaceSource)
	alertSourceModel := new(model.AlertSource)

	err := model.GetOneById(service.ctx, alertSourceModel.TableName(), req.GetSourceId(), &alertSourceModel)
	if err != nil {
		return err
	}
	workspaceSourceModel.SourceId = req.GetSourceId()
	workspaceSourceModel.SourceName = req.GetSourceName()
	workspaceSourceModel.WorkspaceId = id
	workspaceSourceModel.Token = uuid.New().String()
	workspaceSourceModel.Configuration = "{}"
	workspaceSourceModel.LevelField = alertSourceModel.DefaultLevelField
	workspaceSourceModel.Description = req.GetDescription()

	return workspaceSourceModel.Create(service.ctx, id)
}

// UpdateAlertSource 更新数据源
func (service *WorkspaceSourceService) UpdateAlertSource(id uint64, req pb.EditWorkspaceSourceRequest) error {
	workspaceSourceModel := new(model.WorkspaceSource)
	err := model.GetOneById(service.ctx, workspaceSourceModel.TableName(), id, &workspaceSourceModel)
	if err != nil {
		return err
	}
	if req.LevelField != nil {
		workspaceSourceModel.LevelField = req.GetLevelField()
	}
	if req.SourceName != nil {
		workspaceSourceModel.SourceName = req.GetSourceName()
	}
	if req.Description != nil {
		workspaceSourceModel.Description = req.GetDescription()
	}

	return model.EditOneById(service.ctx, workspaceSourceModel.TableName(), id, workspaceSourceModel)
}

// RefreshToken 刷新token
func (service *WorkspaceSourceService) RefreshToken(id uint64) (string, error) {
	workspaceSourceModel := new(model.WorkspaceSource)
	err := model.GetOneById(service.ctx, workspaceSourceModel.TableName(), id, &workspaceSourceModel)
	if err != nil {
		return "", err
	}
	newToken := uuid.New().String()
	workspaceSourceModel.Token = newToken
	err = model.EditOneById(service.ctx, workspaceSourceModel.TableName(), id, &workspaceSourceModel)
	if err != nil {
		return "", err
	}
	return newToken, nil
}

func (service *WorkspaceSourceService) UpdateLevelMapping(req pb.LevelMappingRequest) error {
	levelMapping := new(model.LevelSourceMapping)
	err := model.GetOneById(service.ctx, levelMapping.TableName(), req.GetId(), &levelMapping)
	if err != nil {
		return err
	}
	ok, err := levelMapping.ExistAlertSourceLevel(service.ctx, req.GetAlertSourceLevel())
	if err != nil {
		return err
	}
	if ok {
		return response.AlertLevelAlreadyExistErr
	}
	levelMapping.AlertSourceLevel = req.GetAlertSourceLevel()
	return model.EditOneById(service.ctx, levelMapping.TableName(), req.GetId(), &levelMapping)
}
