package service

import (
	"context"
	"darkhawk/app/biz/model"
	ipb "darkhawk/hertz_gen/integration"
	pb "darkhawk/hertz_gen/workspace_source"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/google/uuid"
	"github.com/hertz-contrib/sessions"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SystemIntegrationService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewSystemIntegrationService(ctx context.Context, c *app.RequestContext) *SystemIntegrationService {
	return &SystemIntegrationService{
		ctx: ctx,
		c:   c,
	}
}

// CreateSystemIntegration 创建集成
func (service *SystemIntegrationService) CreateSystemIntegration(req pb.RelatedWorkspaceSourceRequest) error {
	integrationModel := new(model.Integration)
	alertSourceModel := new(model.AlertSource)

	err := model.GetOneById(service.ctx, alertSourceModel.TableName(), req.GetSourceId(), &alertSourceModel)
	if err != nil {
		return err
	}
	integrationModel.SourceId = req.GetSourceId()
	integrationModel.Name = req.GetSourceName()
	integrationModel.Token = uuid.New().String()
	integrationModel.LevelField = alertSourceModel.DefaultLevelField
	integrationModel.Description = req.GetDescription()
	integrationModel.Disabled = false

	return integrationModel.CreateIntegration(service.ctx)
}

func (service *SystemIntegrationService) UpdateSystemIntegration(id uint64, req ipb.IntegrationRequest) error {
	integrationModel := new(model.Integration)
	err := model.GetOneById(service.ctx, integrationModel.TableName(), id, &integrationModel)
	if err != nil {
		return err
	}
	integrationModel.Name = req.GetName()
	integrationModel.Token = req.GetToken()
	integrationModel.LevelField = req.LevelField
	integrationModel.Description = req.GetDescription()
	return model.EditOneById(service.ctx, integrationModel.TableName(), id, &integrationModel)
}

// AllIntegration 获取所有集成
func (service *SystemIntegrationService) AllIntegration() ([]model.IntegrationItem, error) {
	integrationModel := new(model.Integration)
	data, err := integrationModel.IntegrationList(service.ctx)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// IntegrationDetail 集成详情
func (service *SystemIntegrationService) IntegrationDetail(id uint64) (model.IntegrationItem, error) {
	integrationModel := new(model.Integration)
	data, err := integrationModel.GetIntegrationDetail(service.ctx, id)
	if err != nil {
		return data, err
	}
	return data, nil
}

// ChangeSystemIntegrationStatus 改变状态
func (service *SystemIntegrationService) ChangeSystemIntegrationStatus(id uint64) error {
	integrationModel := new(model.Integration)
	err := model.GetOneById(service.ctx, integrationModel.TableName(), id, &integrationModel)
	if err != nil {
		return err
	}
	var u = make(map[string]interface{})
	if integrationModel.Disabled {
		u["disabled"] = false
	} else {
		u["disabled"] = true
	}
	return model.EditOneById(service.ctx, integrationModel.TableName(), id, u)
}

// DeleteSystemIntegration 删除集成
func (service *SystemIntegrationService) DeleteSystemIntegration(id uint64) error {
	integrationModel := new(model.Integration)
	err := model.GetOneById(service.ctx, integrationModel.TableName(), id, &integrationModel)
	if err != nil {
		return err
	}
	return model.Transaction(service.ctx, func(tx *gorm.DB) error {
		err := model.DeleteOneById(service.ctx, integrationModel.TableName(), id, &integrationModel)
		if err != nil {
			return err
		}
		irModel := new(model.IntegrationRouter)
		var routers []uint64
		err = model.GetPluck(service.ctx, irModel.TableName(), "id", &routers, model.WhereWithScope("integration_id", id))
		if err != nil {
			return err
		}
		err = model.DeleteByScope(service.ctx, irModel.TableName(), model.IntegrationRouter{}, model.InWithScope("id", routers))
		if err != nil {
			return err
		}
		wrModel := new(model.WorkspaceRouter)
		err = model.DeleteByScope(service.ctx, wrModel.TableName(), model.WorkspaceRouter{}, model.InWithScope("router_id", routers))
		if err != nil {
			return err
		}
		return nil
	})
}

// CreateIntegrationRouter 创建集成路由
func (service *SystemIntegrationService) CreateIntegrationRouter(id uint64, req ipb.IntegrationRouterRequest) error {
	integrationModel := new(model.Integration)
	err := model.GetOneById(service.ctx, integrationModel.TableName(), id, &integrationModel)
	if err != nil {
		return err
	}
	session := sessions.Default(service.c)
	uid := session.Get("uid")
	return model.Transaction(service.ctx, func(tx *gorm.DB) error {
		var tagFilters []model.Filters
		for _, filters := range req.GetFilters() {
			var tf model.Filters
			var tags []model.TagFilter
			for _, v := range filters.GetValues() {
				filter := model.TagFilter{
					Tag:       v.GetTag(),
					Operation: v.GetOperation(),
					Value:     v.GetValue(),
				}
				tags = append(tags, filter)
			}
			tf.Values = tags
			tagFilters = append(tagFilters, tf)
		}
		ir := model.IntegrationRouter{
			IntegrationId: id,
			Workspaces:    req.GetWorkspaces(),
			Filters:       tagFilters,
			Next:          req.GetNext(),
			Uid:           uid.(uint64),
		}
		var c int64
		err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Table(ir.TableName()).Count(&c).Error
		if err != nil {
			return err
		}
		ir.Sort = c + 1
		err = tx.Table(ir.TableName()).Create(&ir).Error
		if err != nil {
			return err
		}
		for _, v := range req.GetWorkspaces() {
			wr := model.WorkspaceRouter{
				WorkspaceId: v,
				RouterId:    ir.Id,
			}
			err = tx.Table(wr.TableName()).Create(&wr).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
}

// UpdateIntegrationRouter 编辑集成路由
func (service *SystemIntegrationService) UpdateIntegrationRouter(id uint64, req ipb.IntegrationRouterRequest) error {
	irModel := new(model.IntegrationRouter)
	err := model.GetOneById(service.ctx, irModel.TableName(), id, &irModel)
	if err != nil {
		return err
	}
	session := sessions.Default(service.c)
	uid := session.Get("uid")
	return model.Transaction(service.ctx, func(tx *gorm.DB) error {
		var tagFilters []model.Filters
		for _, filters := range req.GetFilters() {
			var tf model.Filters
			var tags []model.TagFilter
			for _, v := range filters.GetValues() {
				filter := model.TagFilter{
					Tag:       v.GetTag(),
					Operation: v.GetOperation(),
					Value:     v.GetValue(),
				}
				tags = append(tags, filter)
			}
			tf.Values = tags
			tagFilters = append(tagFilters, tf)
		}

		irModel.Workspaces = req.GetWorkspaces()
		irModel.Filters = tagFilters
		irModel.Next = req.GetNext()
		irModel.Uid = uid.(uint64)

		err = tx.Debug().Table(irModel.TableName()).Save(&irModel).Error
		if err != nil {
			return err
		}
		var wr model.WorkspaceRouter
		err = tx.Table(wr.TableName()).Where("router_id", id).Delete(model.WorkspaceRouter{}).Error
		if err != nil {
			return err
		}
		for _, v := range req.GetWorkspaces() {
			wr := model.WorkspaceRouter{
				WorkspaceId: v,
				RouterId:    id,
			}
			err = tx.Table(wr.TableName()).Create(&wr).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
}

// GetIntegrationRouterList 获取集成路由列表
func (service *SystemIntegrationService) GetIntegrationRouterList(id uint64) ([]model.IntegrationRouterNames, error) {
	var data []model.IntegrationRouterNames
	irModel := new(model.IntegrationRouter)
	err := model.GetAll(service.ctx, irModel.TableName(), &data, model.WhereWithScope("integration_id", id), model.OrderScope("sort"))
	if err != nil {
		return nil, err
	}
	workspaceModel := new(model.Workspace)
	workspaceMap, err := workspaceModel.IdMap(service.ctx)
	if err != nil {
		return nil, err
	}
	for k, v := range data {
		var names []string
		for _, i := range v.Workspaces {
			names = append(names, workspaceMap[i])
		}
		data[k].WorkspaceNames = names
	}
	return data, nil
}

// DeleteIntegrationRouter 删除集成路由
func (service *SystemIntegrationService) DeleteIntegrationRouter(id uint64) error {
	routerModel := new(model.IntegrationRouter)
	err := model.GetOneById(service.ctx, routerModel.TableName(), id, &routerModel)
	if err != nil {
		return err
	}
	return model.Transaction(service.ctx, func(tx *gorm.DB) error {
		err := tx.Table(routerModel.TableName()).Where("id", id).Delete(&routerModel).Error
		if err != nil {
			return err
		}
		var wr model.WorkspaceRouter
		return tx.Table(wr.TableName()).Where("router_id", id).Delete(model.WorkspaceRouter{}).Error
	})
}

// ChangeRouterSort 改变路由排序
func (service *SystemIntegrationService) ChangeRouterSort(req []uint64) error {
	routerModel := new(model.IntegrationRouter)
	for k, v := range req {
		err := model.EditOneById(service.ctx, routerModel.TableName(), v, &model.IntegrationRouter{Sort: int64(k) + 1})
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}
