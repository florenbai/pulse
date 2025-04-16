package model

import (
	"context"
	"darkhawk/app/biz/dal/mysql"
	"gorm.io/gorm/clause"
)

type AlertRule struct {
	Model
	Name        string            `gorm:"type:varchar(100);column:name;" json:"name" form:"name"`
	Health      string            `gorm:"type:varchar(100);column:health;" json:"health" form:"health"`
	Annotations map[string]string `gorm:"column:annotations;serializer:json" json:"annotations"`
	Labels      map[string]string `gorm:"column:labels;serializer:json" json:"labels"`
	Duration    int64             `gorm:"column:duration;" json:"duration"`
	Query       string            `gorm:"column:query;" json:"query"`
	Gid         uint64            `gorm:"column:gid;" json:"gid"`
	SourceId    uint64            `gorm:"column:source_id;" json:"sourceId"`
}

type AlertRuleList struct {
	AlertRule
	SourceType string `gorm:"type:varchar(100);column:source_type;comment:告警源类型" json:"sourceType" form:"sourceType"`
}

func (m *AlertRule) TableName() string {
	return mysql.AlertRuleTable
}

func (m *AlertRule) ClausesRule(ctx context.Context, rules []AlertRule) error {
	return mysql.DB.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "name"}, {Name: "gid"}},
		UpdateAll: true,
	}).Create(&rules).Error
}

func (m *AlertRule) GetRuleListByGroup(ctx context.Context, id uint64) ([]AlertRuleList, error) {
	var data []AlertRuleList
	err := mysql.DB.WithContext(ctx).Table(m.TableName()).Select("alert_rule.*,alert_rule_source.source_type").
		Joins("LEFT JOIN alert_rule_group ON alert_rule_group.id = alert_rule.gid").
		Joins("LEFT JOIN alert_rule_source ON alert_rule_group.rule_source = alert_rule_source.id").Where("alert_rule.gid", id).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil

}
