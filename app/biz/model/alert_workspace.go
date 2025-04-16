package model

import (
	"darkhawk/app/biz/dal/mysql"
)

type AlertWorkspace struct {
	Id          uint64 `gorm:"column:id;primary_key;comment:主键id" json:"id" form:"id"` // 主键
	AlertId     uint64 `gorm:"column:alert_id;comment:告警编号" json:"alertId"`
	WorkspaceId uint64 `gorm:"column:workspace_id;comment:工作区编号" json:"workspaceId"`
}

func (m *AlertWorkspace) TableName() string {
	return mysql.AlertWorkspaceTable
}
