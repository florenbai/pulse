package model

import (
	"context"
	"darkhawk/app/biz/dal/mysql"
)

type LevelIntegrationMapping struct {
	Id            uint64 `gorm:"column:id;primary_key;comment:主键id" json:"id" form:"id"` // 主键
	LevelId       uint64 `gorm:"column:level_id;comment:系统告警等级编号" json:"levelId"`
	IntegrationId uint64 `gorm:"column:integration_id;comment:工作区告警源编号" json:"integrationId"`
	AlertLevel    string `gorm:"column:alert_level;comment:告警源等级" json:"alertLevel"`
}

type LevelIntegrationMappingItem struct {
	LevelIntegrationMapping
	LevelName string `gorm:"column:level_name;" json:"levelName"`
	Color     string `gorm:"column:color;" json:"color"`
}

func (m *LevelIntegrationMapping) TableName() string {
	return mysql.LevelIntegrationMappingTable
}

func (m *LevelIntegrationMapping) All(ctx context.Context, id uint64) ([]LevelIntegrationMappingItem, error) {
	var data []LevelIntegrationMappingItem
	err := mysql.DB.WithContext(ctx).Table(m.TableName()).
		Select("level_integration_mapping.*,alert_level.level_name,alert_level.color").
		Joins("LEFT JOIN alert_level ON alert_level.id = level_integration_mapping.level_id").
		Where("integration_id", id).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (m *LevelIntegrationMapping) UpdateMapping(ctx context.Context, mapping []LevelIntegrationMapping) error {
	for _, v := range mapping {
		err := mysql.DB.WithContext(ctx).Table(m.TableName()).Where("id", v.Id).Update("alert_level", v.AlertLevel).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *LevelIntegrationMapping) GetLevelSourceMapByIntegration(ctx context.Context, id uint64) (map[string]string, error) {
	var levelSource []LevelIntegrationMappingItem
	var data = make(map[string]string)
	err := mysql.DB.WithContext(ctx).Table(m.TableName()).Select("level_integration_mapping.*,alert_level.level_name").Joins("LEFT JOIN alert_level ON alert_level.id = level_integration_mapping.level_id").Where("integration_id", id).Find(&levelSource).Error
	if err != nil {
		return nil, err
	}
	for _, v := range levelSource {
		data[v.AlertLevel] = v.LevelName
	}
	return data, nil
}
