package model

import "darkhawk/app/biz/dal/mysql"

type RewriteType string

const (
	TagTypeExtract RewriteType = "extract"
	TagTypeRewrite RewriteType = "rewrite"
	TagTypeDelete  RewriteType = "delete"
)

type TagRewrite struct {
	Model
	Filters       []TagFilter `gorm:"column:filters;comment:匹配条件;serializer:json" json:"filters"`
	IntegrationId uint64      `gorm:"type:int(11);column:integration_id;comment:集成编号" json:"integrationId" form:"integrationId"`
	RewriteItem   RewriteItem `gorm:"column:rewrite_item;comment:匹配条件;serializer:json" json:"rewriteItem"`
	RewriteType   RewriteType `gorm:"column:rewrite_type;comment:重写类型;" json:"rewriteType"`
	Uid           uint64      `json:"uid"`
}

type RewriteItem struct {
	OldTag     string   `json:"oldTag"`
	Expression string   `json:"expression"`
	NewTag     string   `json:"newTag"`
	Value      string   `json:"value"`
	DeleteTag  []string `json:"deleteTag"`
}

func (m *TagRewrite) TableName() string {
	return mysql.TagRewriteTable
}
