package model

import (
	"context"
	"darkhawk/app/biz/dal/mysql"
	"gorm.io/gorm/clause"
)

type Agents struct {
	Id            uint64     `gorm:"column:id;primary_key;comment:主键id" json:"id" form:"id"` // 主键
	ClientIp      string     `gorm:"client_ip" json:"clientIp"`
	RuleSourceId  uint64     `gorm:"rule_source_id" json:"ruleSourceId"`
	HeartbeatTime *LocalTime `json:"heartbeatTime"`
	ErrorMessage  string     `json:"errorMessage"`
	Status        *bool      `json:"status"`
}

func (m *Agents) TableName() string {
	return mysql.Agents
}

func (m *Agents) HeartBeat(ctx context.Context) error {
	return mysql.DB.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "client_ip"}, {Name: "rule_source_id"}},
		UpdateAll: true,
	}).Create(&m).Error
}
