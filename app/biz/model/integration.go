package model

import (
	"context"
	"darkhawk/app/biz/dal/mysql"
	"gorm.io/gorm"
)

type Integration struct {
	Model
	SourceId    uint64 `gorm:"type:int(11);column:source_id;comment:告警源编号" json:"sourceId" form:"sourceId"`
	Name        string `gorm:"column:name;comment:告警源名称" json:"name" form:"name"`
	Token       string `gorm:"type:varchar(100);column:token;comment:token" json:"token"`
	LevelField  string `gorm:"type:varchar(100);column:level_field;comment:告警等级字段" json:"levelField"`
	Description string `json:"description"`
	Disabled    bool   `json:"disabled"`
}

type IntegrationItem struct {
	Integration
	Icon       string `gorm:"type:varchar(100);column:icon;comment:图标" json:"icon"`
	EventCount int    `gorm:"type:int(11);column:eventCount;comment:告警事件数量" json:"eventCount"`
}

func (m *Integration) TableName() string {
	return mysql.SystemIntegrationTable
}

// CreateIntegration 创建集成
func (m *Integration) CreateIntegration(ctx context.Context) error {
	levelModel := new(AlertLevel)
	LevelIntegrationMappingModel := new(LevelIntegrationMapping)
	data, err := levelModel.All(ctx)
	if err != nil {
		return err
	}
	var lsm []LevelIntegrationMapping
	return Transaction(ctx, func(tx *gorm.DB) error {
		err := tx.Table(m.TableName()).Create(&m).Error
		if err != nil {
			return err
		}
		for _, v := range data {
			l := LevelIntegrationMapping{
				LevelId:       v.Id,
				IntegrationId: m.Id,
				AlertLevel:    v.LevelName,
			}
			lsm = append(lsm, l)
		}
		return tx.Table(LevelIntegrationMappingModel.TableName()).Create(&lsm).Error
	})
}

// IntegrationList 获取集成列表
func (m *Integration) IntegrationList(ctx context.Context) ([]IntegrationItem, error) {
	var ws []IntegrationItem
	err := mysql.DB.WithContext(ctx).Table(m.TableName()).
		Select("system_integration.*,alert_source.icon,tmp.c as eventCount").
		Joins("LEFT JOIN alert_source ON alert_source.id = system_integration.source_id").Joins("LEFT JOIN (SELECT count(*) as c,integration_id FROM alert_event GROUP BY integration_id) as tmp ON system_integration.id = tmp.integration_id").Find(&ws).Error
	if err != nil {
		return ws, err
	}
	return ws, nil
}

// GetIntegrationDetail 获取集成详情
func (m *Integration) GetIntegrationDetail(ctx context.Context, id uint64) (IntegrationItem, error) {
	var data IntegrationItem
	err := mysql.DB.WithContext(ctx).Table(m.TableName()).
		Select("system_integration.*,alert_source.icon").
		Joins("LEFT JOIN alert_source ON alert_source.id = system_integration.source_id").Where("system_integration.id", id).Find(&data).Error
	if err != nil {
		return data, err
	}
	return data, nil
}
