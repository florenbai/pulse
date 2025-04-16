package model

import (
	"context"
	"darkhawk/app/biz/dal/mysql"
	"gorm.io/gorm"
)

type Schedule struct {
	Model
	ScheduleName string   `gorm:"column:schedule_name" json:"scheduleName"`
	ScheduleDesc string   `gorm:"column:schedule_desc" json:"scheduleDesc"`
	TeamId       uint64   `gorm:"column:team_id" json:"teamId"`
	StartRange   string   `gorm:"column:start_range" json:"startRange"`
	EndRange     string   `gorm:"column:end_range" json:"endRange"`
	PeriodType   string   `gorm:"column:period_type" json:"periodType"`
	Period       uint8    `gorm:"column:period" json:"period"`
	Users        []uint64 `gorm:"column:users;serializer:json" json:"users"`
}

type ScheduleList struct {
	Schedule
	TeamName string `gorm:"column:team_name" json:"teamName"`
}

const (
	Schedule_PeriodType_Day   = "day"
	Schedule_PeriodType_Week  = "week"
	Schedule_PeriodType_Month = "month"
)

func (m *Schedule) TableName() string {
	return mysql.ScheduleTable
}

// List 值班列表
func (m *Schedule) List(ctx context.Context, page uint64, pageSize uint64, scopes ...func(*gorm.DB) *gorm.DB) ([]ScheduleList, int64, error) {
	var err error
	var data []ScheduleList

	var i int64
	err = mysql.DB.WithContext(ctx).Table(m.TableName()).Scopes(scopes...).Count(&i).Error
	if err != nil {
		return nil, 0, err
	}
	scopes = append(scopes, OrderScope("id DESC"), Paginate(page, pageSize))
	err = mysql.DB.WithContext(ctx).Select("schedule.*,teams.team_name").Table(m.TableName()).Joins("LEFT JOIN teams ON teams.id = schedule.team_id").Scopes(scopes...).Find(&data).Error
	if err != nil {
		return nil, 0, err
	}
	return data, i, nil
}

func (m *Schedule) GetScheduleUserById(ctx context.Context, id uint64) ([]uint64, error) {
	var users []uint64
	err := mysql.DB.WithContext(ctx).Table(m.TableName()).Joins("LEFT JOIN user_teams ON user_teams.team_id = schedule.team_id").Where("id", id).Pluck("user_teams.user_id", &users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (m *Schedule) All(ctx context.Context) ([]Schedule, error) {
	var data []Schedule
	err := GetAll(ctx, m.TableName(), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
