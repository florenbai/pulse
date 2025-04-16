package model

import (
	"context"
	"darkhawk/app/biz/dal/mysql"
)

const (
	NotifyTypeAlert  = 1
	NotifyTypeRECOVE = 2
)

type StrategyLog struct {
	Model
	AlertId         uint64      `gorm:"column:alert_id" json:"alertId"`
	Uid             uint64      `json:"uid"`
	StrategyContent StrategySet `gorm:"column:strategy_content;comment:策略详情;serializer:json" json:"strategyContent"`
	StrategyId      uint64      `gorm:"column:strategy_id;comment:策略编号;" json:"strategyId"`
	Channels        []string    `gorm:"column:channels;comment:推送渠道;serializer:json" json:"channels"`
	IsNotify        bool        `json:"isNotify"`
	ErrMessage      string      `json:"errMessage"`
	StrategyType    int         `json:"strategyType"`
	NotifyType      int         `json:"notifyType"`
}

type StrategyLogList struct {
	StrategyLog
	Nickname      string `json:"nickname"`
	StrategyTitle string `json:"strategyTitle"`
}

const (
	DefaultStrategy = 1
	CustomStrategy  = 2
)

func (m *StrategyLog) TableName() string {
	return mysql.StrategyLogTable
}

// CheckStrategyUser 是否是分配的用户
func (m *StrategyLog) CheckStrategyUser(ctx context.Context, alertId, uid uint64) (bool, error) {
	i, err := GetCountWithScope(ctx, m.TableName(), WhereWithScope("alert_id", alertId), WhereWithScope("uid", uid))
	if err != nil {
		return false, err
	}
	if i > 0 {
		return true, nil
	}
	return false, nil
}

// GetStrategyLogByEvent 获取分派日志
func (m *StrategyLog) GetStrategyLogByEvent(ctx context.Context, id uint64) ([]StrategyLogList, error) {
	var log []StrategyLogList
	err := mysql.DB.WithContext(ctx).Select("strategy_log.*,user.nickname as nickname").Table(m.TableName()).Joins("LEFT JOIN user ON user.id = strategy_log.uid").Where("alert_id", id).Find(&log).Error
	if err != nil {
		return nil, err
	}
	return log, nil
}

// GetDistinctStrategyByEvent 获取去重的分派
func (m *StrategyLog) GetDistinctStrategyByEvent(ctx context.Context, id uint64) ([]StrategyLog, error) {
	var data []StrategyLog
	err := mysql.DB.WithContext(ctx).Table(m.TableName()).Select("Distinct(strategy_id),strategy_type").Where("alert_id", id).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}
