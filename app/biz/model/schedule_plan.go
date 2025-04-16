package model

import (
	"context"
	"darkhawk/app/biz/dal/mysql"
	"gorm.io/gorm"
	"time"
)

type SchedulePlan struct {
	Model
	ScheduleId uint64     `gorm:"column:schedule_id" json:"scheduleId"`
	Month      string     `gorm:"column:month" json:"month"`
	Plan       []PlanItem `gorm:"column:plan;serializer:json" json:"plan"`
	Uid        uint64     `gorm:"column:uid" json:"uid"`
}

type PlanItem struct {
	Day        int  `json:"day"`
	IsSchedule bool `json:"isSchedule"`
}

type SchedulePlanList struct {
	SchedulePlan
	Nickname string `json:"nickname"`
}

func (m *SchedulePlan) TableName() string {
	return mysql.SchedulePlanTable
}

func (m *SchedulePlan) GetSchedulePlanList(ctx context.Context, scopes ...func(*gorm.DB) *gorm.DB) ([]SchedulePlanList, error) {
	var data []SchedulePlanList
	err := mysql.DB.WithContext(ctx).Select("schedule_plan.*,user.nickname").Table(m.TableName()).Joins("LEFT JOIN user ON user.id = schedule_plan.uid").Scopes(scopes...).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (m *SchedulePlan) GetCurrentMonthSchedulePlanByIds(ctx context.Context, id []uint64) ([]SchedulePlan, error) {
	var data []SchedulePlan
	month := time.Now().Format("2006-01")
	err := mysql.DB.WithContext(ctx).Select("schedule_plan.*").Table(m.TableName()).Where("schedule_id IN ?", id).Where("month", month).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (m *SchedulePlan) GetUsersByIds(ctx context.Context, id []uint64, month string) ([]uint64, error) {
	var data []uint64
	err := mysql.DB.WithContext(ctx).Table(m.TableName()).Where("schedule_id IN ?", id).Where("month", month).Pluck("uid", &data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}
