package model

import "darkhawk/app/biz/dal/mysql"

type IntegrationRouter struct {
	Model
	IntegrationId uint64    `gorm:"type:int(11);column:integration_id;comment:集成编号" json:"integrationId" form:"integrationId"`
	Workspaces    []uint64  `gorm:"type:int(11);column:workspaces;comment:工作区编号;serializer:json" json:"workspaces" form:"workspaces"`
	Filters       []Filters `gorm:"json;column:filters;comment:匹配条件;serializer:json" json:"filters"`
	Sort          int64     `json:"sort"`
	Next          int64     `json:"next"`
	Uid           uint64    `json:"uid"`
}

type IntegrationRouterNames struct {
	IntegrationRouter
	WorkspaceNames []string `gorm:"-" json:"workspaceNames"`
}

type Filters struct {
	Values []TagFilter `json:"values"`
}
type TagFilter struct {
	Tag       string   `json:"tag"`
	Operation string   `json:"operation"`
	Value     []string `json:"value"`
}

func (m *IntegrationRouter) TableName() string {
	return mysql.IntegrationRouterTable
}
