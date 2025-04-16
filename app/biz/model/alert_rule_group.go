package model

import (
	"context"
	"darkhawk/app/biz/dal/mysql"
	"gorm.io/gorm"
)

type AlertRuleGroup struct {
	Id         uint64 `gorm:"column:id;primary_key;comment:主键id" json:"id" form:"id"` // 主键
	File       string `json:"file"`
	RuleSource uint64 `json:"ruleSource"`
	GroupName  string `json:"groupName"`
}

func (m *AlertRuleGroup) TableName() string {
	return mysql.AlertRuleGroupTable
}

func (m *AlertRuleGroup) GetGroupNameMap(ctx context.Context, ruleSourceId uint64) (map[string]uint64, error) {
	var groups []AlertRuleGroup
	var data = make(map[string]uint64)
	err := GetAll(ctx, m.TableName(), &groups, WhereWithScope("rule_source", ruleSourceId))
	if err != nil {
		return nil, err
	}
	for _, v := range groups {
		data[v.GroupName] = v.Id
	}
	return data, nil
}

func (m *AlertRuleGroup) DeleteGroupAndChildrenRule(ctx context.Context, id uint64) error {
	alertRuleModel := new(AlertRule)
	return Transaction(ctx, func(tx *gorm.DB) error {
		err := tx.Table(m.TableName()).Where("id", id).Delete(&m).Error
		if err != nil {
			return err
		}
		err = tx.Table(alertRuleModel.TableName()).Where("gid", id).Delete(AlertRule{}).Error
		if err != nil {
			return err
		}
		return nil
	})
}
