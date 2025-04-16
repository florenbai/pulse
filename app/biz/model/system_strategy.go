package model

import (
	"context"
	"darkhawk/app/biz/dal/mysql"
	"gorm.io/gorm"
)

type SystemStrategy struct {
	Model
	StrategyName string        `gorm:"column:strategy_name;comment:策略名称" json:"strategyName"`
	StrategyDesc string        `gorm:"column:strategy_desc;comment:策略描述" json:"strategyDesc"`
	TemplateId   uint64        `gorm:"column:template_id;comment:策略通知模板" json:"templateId"`
	Delay        int64         `gorm:"column:delay;comment:延迟通知" json:"delay"`
	StrategySet  []StrategySet `gorm:"column:strategy_set;comment:策略详情;serializer:json" json:"strategySet"`
	Uid          uint64        `gorm:"column:uid;comment:用户编号;" json:"uid"`
}

type SystemStrategyList struct {
	SystemStrategy
	Nickname string `json:"nickname"`
}

func (m *SystemStrategy) TableName() string {
	return mysql.SystemStrategyTable
}

// GetSystemStrategyList 获取默认告警策略列表
func (m *SystemStrategy) GetSystemStrategyList(ctx context.Context, page, pageSize uint64, scope func(db *gorm.DB) *gorm.DB) ([]SystemStrategyList, int64, error) {
	var i int64
	var data []SystemStrategyList
	err := mysql.DB.WithContext(ctx).Table(m.TableName()).Scopes(scope).Count(&i).Error
	if err != nil {
		return nil, 0, err
	}
	var scopes []func(db *gorm.DB) *gorm.DB
	scopes = append(scopes, Paginate(page, pageSize), scope)
	err = mysql.DB.WithContext(ctx).Table(m.TableName()).Select("system_strategy.*,user.nickname").Joins("LEFT JOIN user ON user.id = system_strategy.uid").Scopes(scopes...).Find(&data).Error
	if err != nil {
		return nil, i, err
	}
	return data, i, nil
}

// GetSystemStrategyByWorkspaceId 根据工作区编号获取默认告警策略
func (m *SystemStrategy) GetSystemStrategyByWorkspaceId(ctx context.Context, workspaceId uint64) error {
	return mysql.DB.WithContext(ctx).Table(mysql.WorkspaceTable).Select("system_strategy.*").Joins("LEFT JOIN system_strategy ON workspace.system_strategy = system_strategy.id").Where("workspace.id", workspaceId).Take(&m).Error
}

// GetStrategyMap 获取系统策略map
func (m *SystemStrategy) GetStrategyMap(ctx context.Context) (map[uint64]string, error) {
	var data []SystemStrategy
	var strategyMap = make(map[uint64]string)
	err := GetAll(ctx, m.TableName(), &data)
	if err != nil {
		return nil, err
	}
	for _, v := range data {
		strategyMap[v.Id] = v.StrategyName
	}
	return strategyMap, nil
}
