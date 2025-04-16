package model

import (
	"context"
	"darkhawk/app/biz/dal/mysql"
)

type AlertAggregation struct {
	Model
	AggregationName string                 `gorm:"column:aggregation_name;comment:聚合名称" json:"aggregationName"`
	AggregationDesc string                 `gorm:"column:aggregation_desc;comment:聚合描述" json:"aggregationDesc"`
	WorkspaceId     uint64                 `gorm:"column:workspace_id;comment:工作区编号" json:"workspaceId"`
	LevelDimension  bool                   `gorm:"column:level_dimension;comment:等级维度" json:"levelDimension"`
	TitleDimension  bool                   `gorm:"column:title_dimension;comment:标题维度" json:"titleDimension"`
	TagsDimension   []StrategyFilterValues `gorm:"column:tags_dimension;comment:标签维度;serializer:json" json:"tagsDimension"`
	Windows         int64                  `gorm:"column:windows;comment:聚合窗口" json:"windows"`
	Storm           int64                  `gorm:"column:storm;comment:风暴预警" json:"storm"`
	Uid             uint64                 `gorm:"column:uid;comment:用户编号;" json:"uid"`
	Status          StrategyStatus         `gorm:"column:status;comment:状态" json:"status"`
}

type AlertAggregationList struct {
	AlertAggregation
	Nickname string `json:"nickname"`
}

func (m *AlertAggregation) TableName() string {
	return mysql.AlertAggregationTable
}

// AggregationAll 获取工作区所有启用和禁用的告警聚合
func (m *AlertAggregation) AggregationAll(ctx context.Context, id uint64) ([]AlertAggregationList, error) {
	var data []AlertAggregationList
	err := mysql.DB.WithContext(ctx).Table(m.TableName()).
		Select("alert_aggregation.*,user.nickname").
		Joins("LEFT JOIN user ON user.id = alert_aggregation.uid").
		Where("workspace_id", id).Where("alert_aggregation.status IN ?", []string{"enabled", "disabled"}).Find(&data).Error

	if err != nil {
		return nil, err
	}
	return data, nil
}

// GetEnableAggregationByWorkspace 获取工作区可用的告警聚合
func (m *AlertAggregation) GetEnableAggregationByWorkspace(ctx context.Context, id uint64) ([]AlertAggregation, error) {
	var data []AlertAggregation
	err := mysql.DB.WithContext(ctx).Table(m.TableName()).Where("workspace_id", id).Where("status", "enabled").Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}
