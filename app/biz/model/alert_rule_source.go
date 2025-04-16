package model

import (
	"context"
	"darkhawk/app/biz/dal/mysql"
	"gorm.io/gorm"
)

type AlertRuleSource struct {
	Model
	SourceName string `gorm:"type:varchar(100);column:source_name;comment:告警源名称" json:"sourceName" form:"sourceName"`
	SourceType string `gorm:"type:varchar(100);column:source_type;comment:告警源类型" json:"sourceType" form:"sourceType"`
	Address    string `gorm:"type:varchar(100);column:address;comment:地址" json:"address"`
	AutoSync   bool   `gorm:"type:tinyint(1);column:auto_sync;comment:是否自动同步" json:"autoSync"`
	Mark       string `gorm:"type:varchar(100);column:mark;comment:备注" json:"mark"`
	Sign       string `gorm:"type:varchar(100);column:sign;comment:数据源标识" json:"sign"`
}

func (m *AlertRuleSource) TableName() string {
	return mysql.AlertRuleSourceTable
}

func (m *AlertRuleSource) GetRuleSourceByGroupId(ctx context.Context, id uint64) (*AlertRuleSource, error) {
	err := mysql.DB.WithContext(ctx).Table(m.TableName()).Select("alert_rule_source.*").Joins("LEFT JOIN alert_rule_group ON alert_rule_group.rule_source = alert_rule_source.id").Where("alert_rule_group.id", id).Take(&m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (m *AlertRuleSource) DeleteSource(ctx context.Context, id uint64) error {
	return Transaction(ctx, func(tx *gorm.DB) error {
		err := tx.Table(m.TableName()).Where("id", id).Delete(&m).Error
		if err != nil {
			return err
		}
		alertRuleModel := new(AlertRule)
		err = tx.Table(alertRuleModel.TableName()).Where("gid", id).Delete(&alertRuleModel).Error
		if err != nil {
			return err
		}
		alertRuleGroupModel := new(AlertRuleGroup)
		err = tx.Table(alertRuleGroupModel.TableName()).Where("rule_source", id).Delete(&alertRuleGroupModel).Error
		if err != nil {
			return err
		}
		return nil
	})
}
