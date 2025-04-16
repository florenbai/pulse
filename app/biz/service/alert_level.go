package service

import (
	"context"
	"darkhawk/app/biz/dal/mysql"
	"darkhawk/app/biz/model"
	pb "darkhawk/hertz_gen/alert_level"
	"darkhawk/pkg/response"
	"github.com/cloudwego/hertz/pkg/app"
	"gorm.io/gorm"
)

type AlertLevelService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewAlertLevelService(ctx context.Context, c *app.RequestContext) *AlertLevelService {
	return &AlertLevelService{
		ctx: ctx,
		c:   c,
	}
}

func (service *AlertLevelService) CreateAlertLevel(req pb.LevelRequest) error {
	alertLevelModel := new(model.AlertLevel)
	ok, err := model.ExistName(service.ctx, alertLevelModel.TableName(), "level_name", req.GetLevelName(), nil)
	if err != nil {
		return err
	}
	if ok {
		return response.NameAlreadyExistErr
	}
	alertLevelModel.LevelName = req.GetLevelName()
	alertLevelModel.Color = req.GetColor()
	alertLevelModel.LevelDesc = req.GetLevelDesc()
	alertLevelModel.IsDefault = false

	var integrations []uint64
	levelIntegrationModel := new(model.LevelIntegrationMapping)
	err = model.GetPluck(service.ctx, levelIntegrationModel.TableName(), "integration_id", &integrations, model.GroupScope("integration_id"))
	if err != nil {
		return err
	}
	return model.Transaction(service.ctx, func(tx *gorm.DB) error {
		err := tx.Table(alertLevelModel.TableName()).Create(&alertLevelModel).Error
		if err != nil {
			return err
		}
		for _, v := range integrations {
			mapping := model.LevelIntegrationMapping{
				LevelId:       alertLevelModel.Id,
				IntegrationId: v,
				AlertLevel:    "",
			}
			err = tx.Table(mapping.TableName()).Create(&mapping).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (service *AlertLevelService) UpdateAlertLevel(id uint64, req pb.LevelRequest) error {
	alertLevelModel := new(model.AlertLevel)
	err := model.GetOneById(service.ctx, alertLevelModel.TableName(), id, &alertLevelModel)
	if err != nil {
		return err
	}
	ok, err := model.ExistName(service.ctx, alertLevelModel.TableName(), "level_name", req.GetLevelName(), &id)
	if err != nil {
		return err
	}
	if ok {
		return response.NameAlreadyExistErr
	}
	alertLevelModel.LevelName = req.GetLevelName()
	alertLevelModel.Color = req.GetColor()
	alertLevelModel.LevelDesc = req.GetLevelDesc()
	return model.EditOneById(service.ctx, alertLevelModel.TableName(), id, &alertLevelModel)
}

func (service *AlertLevelService) DeleteAlertLevel(id uint64) error {
	alertLevelModel := new(model.AlertLevel)
	err := model.GetOneById(service.ctx, alertLevelModel.TableName(), id, &alertLevelModel)
	if err != nil {
		return err
	}
	return model.Transaction(service.ctx, func(tx *gorm.DB) error {
		err := tx.Table(alertLevelModel.TableName()).Where("id", id).Delete(alertLevelModel).Error
		if err != nil {
			return err
		}
		err = tx.Table(mysql.LevelIntegrationMappingTable).Where("level_id", id).Delete(model.LevelIntegrationMapping{}).Error
		if err != nil {
			return err
		}
		return nil
	})
}
