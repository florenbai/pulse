package model

import (
	"context"
	"darkhawk/app/biz/dal/mysql"
)

type ChannelTemplate struct {
	Id         uint64 `gorm:"column:id;primary_key;comment:主键id" json:"id" form:"id"` // 主键
	TemplateId uint64 `gorm:"column:template_id;comment:模板编号" json:"templateId" form:"templateId"`
	ChannelId  uint64 `gorm:"column:channel_id;comment:模板编号" json:"channelId" form:"channelId"`
	Content    string `json:"content"`
	Finished   bool   `json:"finished"`
}

type ChannelTemplateDetail struct {
	Id          uint64 `gorm:"column:id;primary_key;comment:主键id" json:"id" form:"id"` // 主键
	TemplateId  uint64 `gorm:"column:template_id;comment:模板编号" json:"templateId" form:"templateId"`
	ChannelId   uint64 `gorm:"column:channel_id;comment:模板编号" json:"channelId" form:"channelId"`
	Content     string `json:"content"`
	Finished    bool   `json:"finished"`
	ChannelName string `json:"channelName"`
	ChannelType string `json:"channelType"`
	ChannelSign string `json:"channelSign"`
}

func (m *ChannelTemplate) TableName() string {
	return mysql.ChannelTemplateTable
}

func (m *ChannelTemplate) GetChannelTemplateDetail(ctx context.Context, id uint64) ([]ChannelTemplateDetail, error) {
	var channels []ChannelTemplateDetail
	err := mysql.DB.WithContext(ctx).Table(m.TableName()).
		Select("alert_channel.channel_name,alert_channel.channel_type,alert_channel.channel_sign,channel_template.*").
		Joins("LEFT JOIN alert_channel ON alert_channel.id = channel_template.channel_id").
		Where("channel_template.template_id", id).Find(&channels).Error
	if err != nil {
		return nil, err
	}
	return channels, nil
}
