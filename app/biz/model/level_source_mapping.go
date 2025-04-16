package model

import (
	"context"
	"darkhawk/app/biz/dal/mysql"
)

type LevelSourceMapping struct {
	Id               uint64 `gorm:"column:id;primary_key;comment:主键id" json:"id" form:"id"` // 主键
	LevelId          uint64 `gorm:"column:level_id;comment:系统告警等级编号" json:"levelId"`
	WsId             uint64 `gorm:"column:ws_id;comment:工作区告警源编号" json:"wsId"`
	WorkspaceId      uint64 `gorm:"type:int(11);column:workspace_id;comment:工作区编号" json:"workspaceId" form:"workspaceId"`
	AlertSourceLevel string `gorm:"column:alert_source_level;comment:告警源等级" json:"alertSourceLevel"`
}

type LevelSourceMappingItem struct {
	LevelSourceMapping
	LevelName string `gorm:"column:level_name;" json:"levelName"`
	Color     string `gorm:"column:color;" json:"color"`
}

func (m *LevelSourceMapping) TableName() string {
	return mysql.LevelSourceMappingTable
}

func (m *LevelSourceMapping) All(ctx context.Context, wsId uint64) ([]LevelSourceMappingItem, error) {
	var data []LevelSourceMappingItem
	err := mysql.DB.WithContext(ctx).Table(m.TableName()).
		Select("level_source_mapping.*,alert_level.level_name,alert_level.color").
		Joins("LEFT JOIN alert_level ON alert_level.id = level_source_mapping.level_id").
		Where("ws_id", wsId).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (m *LevelSourceMapping) ExistAlertSourceLevel(ctx context.Context, levelName string) (bool, error) {
	var i int64
	if m.AlertSourceLevel == levelName {
		return false, nil
	}
	err := mysql.DB.WithContext(ctx).Table(m.TableName()).Where("ws_id", m.WsId).Where("alert_source_level", levelName).Count(&i).Error
	if err != nil {
		return false, err
	}
	if i > 0 {
		return true, nil
	}
	return false, err
}

func (m *LevelSourceMapping) GetLevelSourceMap(ctx context.Context, wsId uint64) (map[string]string, error) {
	var levelSource []LevelSourceMappingItem
	var data = make(map[string]string)
	err := mysql.DB.WithContext(ctx).Table(m.TableName()).Select("level_source_mapping.*,alert_level.level_name").Joins("LEFT JOIN alert_level ON alert_level.id = level_source_mapping.level_id").Where("ws_id", wsId).Find(&levelSource).Error
	if err != nil {
		return nil, err
	}
	for _, v := range levelSource {
		data[v.AlertSourceLevel] = v.LevelName
	}
	return data, nil
}
