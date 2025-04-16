package model

import (
	"context"
	"darkhawk/app/biz/dal/mysql"
)

type AlertAction string

const (
	AlertClaim       = "claim"
	AlertOpened      = "opened"
	AlertClosed      = "closed"
	AlertCancelClaim = "cancel_claim"
)

type AlertLog struct {
	Model
	AlertId uint64      `gorm:"column:alert_id" json:"alertId"`
	Action  AlertAction `json:"action"`
	Uid     uint64      `json:"uid"`
}

type AlertLogList struct {
	AlertLog
	Nickname string `json:"nickname"`
}

func (m *AlertLog) TableName() string {
	return mysql.AlertLog
}

// GetAlertLogByEvent 根据告警编号获取告警日志
func (m *AlertLog) GetAlertLogByEvent(ctx context.Context, id uint64) ([]AlertLogList, error) {
	var log []AlertLogList
	err := mysql.DB.WithContext(ctx).Select("alert_log.*,user.nickname as nickname").Table(m.TableName()).Joins("LEFT JOIN user ON user.id = alert_log.uid").Where("alert_id", id).Find(&log).Error
	if err != nil {
		return nil, err
	}
	return log, nil
}
