package model

import "darkhawk/app/biz/dal/mysql"

type AlertRestrain struct {
	Model
	RestrainType   string   `gorm:"column:restrain_type;comment:抑制类型" json:"restrainType"`
	Fields         []string `gorm:"column:fields;comment:匹配字段;serializer:json" json:"fields"`
	CumulativeTime int64    `gorm:"column:cumulative_time;comment:抑制时长" json:"cumulativeTime"`
	WorkspaceId    uint64   `gorm:"column:workspace_id;comment:工作区编号" json:"workspaceId"`
	Uid            uint64   `gorm:"column:uid;comment:用户编号;" json:"uid"`
}

func (m *AlertRestrain) TableName() string {
	return mysql.AlertRestrainTable
}
