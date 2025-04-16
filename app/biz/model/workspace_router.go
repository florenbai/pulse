package model

import "darkhawk/app/biz/dal/mysql"

type WorkspaceRouter struct {
	Id          uint64 `gorm:"column:id;primary_key;comment:主键id" json:"id" form:"id"` // 主键
	WorkspaceId uint64 `json:"workspaceId"`
	RouterId    uint64 `json:"routerId"`
}

func (m *WorkspaceRouter) TableName() string {
	return mysql.WorkspaceRouterTable
}
