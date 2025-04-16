package model

import (
	"context"
	"darkhawk/app/biz/dal/mysql"
	"gorm.io/gorm"
)

type TeamRule struct {
	Id     uint64 `gorm:"column:id;primary_key;comment:主键id" json:"id" form:"id"` // 主键
	MenuId uint64 `json:"menuId"`
	TeamId uint64 `json:"teamId"`
}

func (m *TeamRule) TableName() string {
	return mysql.TeamsRuleTable
}

func (m *TeamRule) AuthorizationMenu(ctx context.Context, id uint64, menu []uint64) error {
	return Transaction(ctx, func(tx *gorm.DB) error {
		err := tx.Table(m.TableName()).Where("team_id", id).Delete(TeamRule{}).Error
		if err != nil {
			return err
		}
		var rules []TeamRule
		for _, v := range menu {
			rule := TeamRule{TeamId: id, MenuId: v}
			rules = append(rules, rule)
		}
		return tx.Table(m.TableName()).Create(&rules).Error
	})
}
