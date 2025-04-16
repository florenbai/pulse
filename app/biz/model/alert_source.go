package model

import (
	"context"
	"darkhawk/app/biz/dal/mysql"
)

type AlertSource struct {
	Id                uint64 `gorm:"column:id;primary_key;comment:主键id" json:"id" form:"id"` // 主键
	SourceName        string `gorm:"type:varchar(100);column:source_name;comment:告警源名称" json:"sourceName" form:"sourceName"`
	Icon              string `gorm:"type:varchar(100);column:icon;comment:图标" json:"icon"`
	Category          string `gorm:"type:varchar(100);column:category;comment:分类" json:"category"`
	Status            bool   `gorm:"type:tinyint(1);default:true;comment:状态" json:"status" form:"status"`
	DefaultLevelField string `gorm:"type:varchar(100);column:default_level_field;comment:默认告警等级字段" json:"defaultLevelField"`
}

func (m *AlertSource) TableName() string {
	return mysql.AlertSourceTable
}

func (m *AlertSource) All(ctx context.Context) ([]AlertSource, error) {
	var data []AlertSource
	err := GetAll(ctx, m.TableName(), &data, WhereWithScope("status", true))
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (m *AlertSource) SourceMap(ctx context.Context) (map[uint64]string, error) {
	var sm = make(map[uint64]string)
	data, err := m.All(ctx)
	if err != nil {
		return nil, err
	}
	for _, v := range data {
		sm[v.Id] = v.SourceName
	}
	return sm, nil
}
