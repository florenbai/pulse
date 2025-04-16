package model

import (
	"context"
	"darkhawk/app/biz/dal/mysql"
	"gorm.io/gorm"
)

type StrategyStatus string

const StrategyEnabled StrategyStatus = "enabled"
const StrategyDisabled StrategyStatus = "disabled"
const StrategyDeleted StrategyStatus = "deleted"

type AlertStrategy struct {
	Model
	StrategyName   string                 `gorm:"column:strategy_name;comment:策略名称" json:"strategyName"`
	StrategyDesc   string                 `gorm:"column:strategy_desc;comment:策略描述" json:"strategyDesc"`
	WorkspaceId    uint64                 `gorm:"column:workspace_id;comment:策略描述" json:"workspaceId"`
	TemplateId     uint64                 `gorm:"column:template_id;comment:策略通知模板" json:"templateId"`
	Status         StrategyStatus         `gorm:"column:status;comment:状态" json:"status"`
	Delay          int64                  `gorm:"column:delay;comment:延迟通知" json:"delay"`
	SystemStrategy *bool                  `gorm:"column:system_strategy;comment:全局策略" json:"systemStrategy"`
	Continuous     *bool                  `gorm:"column:continuous;comment:接续匹配" json:"continuous"`
	Weight         int64                  `gorm:"column:weight;comment:权重" json:"weight"`
	TimeSlot       TimeSlot               `gorm:"column:time_slot;comment:通知时间段策略;serializer:json" json:"timeSlot"`
	Filters        []StrategyFilterValues `gorm:"column:filters;comment:标签匹配策略;serializer:json" json:"filters"`
	StrategySet    []StrategySet          `gorm:"column:strategy_set;comment:策略详情;serializer:json" json:"strategySet"`
	Uid            uint64                 `gorm:"column:uid;comment:用户编号;" json:"uid"`
}

type TimeSlot struct {
	Enable   bool     `json:"enable"`
	Type     *string  `json:"type"`
	Weeks    []int64  `json:"weeks"`
	Times    []string `json:"times"`
	Calendar *int64   `json:"calendar"`
}

type StrategyFilterValues struct {
	Values []TagFilter `json:"values"`
}

type StrategyFilter struct {
	Tag       string   `json:"tag"`
	Operation string   `json:"operation"`
	Value     []string `json:"value"`
}

type AlertLoop struct {
	Enable   bool  `json:"enable"`
	Interval int64 `json:"interval"`
	Count    int64 `json:"count"`
}

type Progression struct {
	Enable    bool  `json:"enable"`
	Condition int64 `json:"condition"`
	Duration  int64 `json:"duration"`
}

type StrategySet struct {
	Members       []uint64     `json:"members"`
	Teams         []uint64     `json:"teams"`
	Schedules     []uint64     `json:"schedules"`
	AlertChannels [][]uint64   `json:"alertChannels"`
	AlertLoop     *AlertLoop   `json:"alertLoop"`
	Progression   *Progression `json:"progression"`
}

type AlertStrategyList struct {
	AlertStrategy
	Nickname string `json:"nickname"`
}

func (m *AlertStrategy) TableName() string {
	return mysql.AlertStrategyTable
}

func (m *AlertStrategy) Create(ctx context.Context) error {
	var i int64
	return Transaction(ctx, func(tx *gorm.DB) error {
		err := tx.Table(m.TableName()).Where("workspace_id", m.WorkspaceId).Count(&i).Error
		if err != nil {
			return err
		}
		m.Weight = i + 1
		m.Status = StrategyEnabled
		return tx.Table(m.TableName()).Create(&m).Error
	})
}

func (m *AlertStrategy) WorkspaceAll(ctx context.Context, id uint64) ([]AlertStrategyList, error) {
	var data []AlertStrategyList
	err := mysql.DB.WithContext(ctx).Table(m.TableName()).
		Select("alert_strategy.*,user.nickname").
		Joins("LEFT JOIN user ON user.id = alert_strategy.uid").
		Where("workspace_id", id).Where("alert_strategy.status IN ?", []string{"enabled", "disabled"}).Order("alert_strategy.weight ASC").Find(&data).Error

	if err != nil {
		return nil, err
	}
	return data, nil
}

// EnableStrategy 获取已开启的告警通知策略
func (m *AlertStrategy) EnableStrategy(ctx context.Context, id uint64) ([]AlertStrategy, error) {
	var data []AlertStrategy
	err := mysql.DB.WithContext(ctx).Table(m.TableName()).
		Where("workspace_id", id).Where("alert_strategy.status", "enabled").Order("alert_strategy.weight ASC").Find(&data).Error

	if err != nil {
		return nil, err
	}
	return data, nil
}

// GetStrategyMap 获取策略map
func (m *AlertStrategy) GetStrategyMap(ctx context.Context) (map[uint64]string, error) {
	var data []AlertStrategy
	var strategyMap = make(map[uint64]string)
	err := mysql.DB.WithContext(ctx).Table(m.TableName()).Select("id,strategy_name").Find(&data).Error
	if err != nil {
		return nil, err
	}
	for _, v := range data {
		strategyMap[v.Id] = v.StrategyName
	}
	return strategyMap, nil
}
