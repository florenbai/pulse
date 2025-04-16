package service

import (
	"context"
	"darkhawk/app/biz/model"
	pb "darkhawk/hertz_gen/alert_restrain"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

type AlertRestrainService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewAlertRestrainService(ctx context.Context, c *app.RequestContext) *AlertRestrainService {
	return &AlertRestrainService{
		ctx: ctx,
		c:   c,
	}
}

// CreateAlertRestrain 创建告警抑制
func (service *AlertRestrainService) CreateAlertRestrain(id uint64, req pb.RestrainRequest) error {
	workspaceModel := new(model.Workspace)
	err := model.GetOneById(service.ctx, workspaceModel.TableName(), id, &workspaceModel)
	if err != nil {
		return err
	}
	session := sessions.Default(service.c)
	uid := session.Get("uid").(uint64)

	restrainModel := new(model.AlertRestrain)
	restrainModel.RestrainType = req.GetRestrainType()
	restrainModel.Fields = req.GetFields()
	restrainModel.CumulativeTime = req.GetCumulativeTime()
	restrainModel.WorkspaceId = id
	restrainModel.Uid = uid
	return model.Create(service.ctx, restrainModel.TableName(), &restrainModel)
}

// UpdateAlertRestrain 编辑告警抑制
func (service *AlertRestrainService) UpdateAlertRestrain(id uint64, req pb.RestrainRequest) error {
	restrainModel := new(model.AlertRestrain)
	err := model.GetOneById(service.ctx, restrainModel.TableName(), id, &restrainModel)
	if err != nil {
		return err
	}
	session := sessions.Default(service.c)
	uid := session.Get("uid").(uint64)

	restrainModel.RestrainType = req.GetRestrainType()
	fmt.Println(req.GetFields())
	restrainModel.Fields = req.GetFields()
	restrainModel.CumulativeTime = req.GetCumulativeTime()
	restrainModel.Uid = uid
	return model.EditOneById(service.ctx, restrainModel.TableName(), id, &restrainModel)
}

// DeleteAlertRestrain 删除告警抑制
func (service *AlertRestrainService) DeleteAlertRestrain(id uint64) error {
	restrainModel := new(model.AlertRestrain)
	err := model.GetOneById(service.ctx, restrainModel.TableName(), id, &restrainModel)
	if err != nil {
		return err
	}
	return model.DeleteOneById(service.ctx, restrainModel.TableName(), id, restrainModel)
}

func (service *AlertRestrainService) AlertRestrainAll(id uint64) ([]model.AlertRestrain, error) {
	var data []model.AlertRestrain
	restrainModel := new(model.AlertRestrain)
	err := model.GetAll(service.ctx, restrainModel.TableName(), &data, model.WhereWithScope("workspace_id", id))
	if err != nil {
		return nil, err
	}
	return data, nil
}
