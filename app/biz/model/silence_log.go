package model

import "darkhawk/app/biz/dal/mysql"

type SilenceLog struct {
	Model
	AlertId     uint64 `gorm:"column:alert_id" json:"alertId"`
	SilenceId   uint64 `gorm:"column:silence_id" json:"silenceId"`
	WorkspaceId uint64 `gorm:"column:workspace_id" json:"workspaceId"`
}

func (m *SilenceLog) TableName() string {
	return mysql.SilenceLogTable
}
