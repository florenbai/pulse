package model

import (
	"context"
	"darkhawk/app/biz/dal/mysql"
)

type SilenceType string

const (
	SilenceOnce   SilenceType = "once"
	SilencePeriod SilenceType = "period"
)

type AlertSilence struct {
	Model
	SilenceName string      `gorm:"column:silence_name;comment:告警静默名称" json:"silenceName"`
	SilenceDesc string      `gorm:"column:silence_desc;comment:告警静默描述" json:"silenceDesc"`
	WorkspaceId uint64      `gorm:"column:workspace_id;comment:工作区编号" json:"workspaceId"`
	SilenceType SilenceType `gorm:"column:silence_type;comment:静默时间类型" json:"silenceType"`
	SilenceTime SilenceTime `gorm:"column:silence_time;comment:静默时间;serializer:json" json:"silenceTime"`
	Filters     []TagFilter `gorm:"column:filters;comment:标签匹配策略;serializer:json" json:"filters"`
	Uid         uint64      `gorm:"column:uid;comment:用户编号;" json:"uid"`
}

type SilenceTime struct {
	RangeTime []string `json:"rangeTime"`
	Weeks     []uint64 `json:"weeks"`
	Times     []string `json:"times"`
}

func (m *AlertSilence) TableName() string {
	return mysql.AlertSilenceTable
}

func (m *AlertSilence) WorkspaceAll(ctx context.Context, id uint64) ([]AlertSilence, error) {
	var data []AlertSilence
	err := mysql.DB.WithContext(ctx).Table(m.TableName()).
		Where("workspace_id", id).Find(&data).Error

	if err != nil {
		return nil, err
	}
	return data, nil
}
