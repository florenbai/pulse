package model

import (
	"context"
	"darkhawk/app/biz/dal/mysql"
	"darkhawk/pkg/utils"
	"gorm.io/gorm"
)

type AlertTemplate struct {
	Model
	TemplateName string `gorm:"type:varchar(100);column:template_name;comment:模板名称" json:"templateName" form:"templateName"`
	TemplateDesc string `gorm:"type:text;column:template_desc;comment:模板描述" json:"templateDesc" form:"templateDesc"`
}

type AlertTemplateDetail struct {
	AlertTemplate
	Channels []ChannelTemplateDetail `gorm:"-" json:"channels"`
}

func (m *AlertTemplate) TableName() string {
	return mysql.AlertTemplateTable
}

func (m *AlertTemplate) CreateTemplate(ctx context.Context, channels []uint64) error {
	return Transaction(ctx, func(tx *gorm.DB) error {
		err := tx.Table(m.TableName()).Create(&m).Error
		if err != nil {
			return err
		}
		for _, v := range channels {
			item := ChannelTemplate{
				TemplateId: m.Id,
				ChannelId:  v,
				Content:    "",
				Finished:   false,
			}
			err := tx.Table(item.TableName()).Create(&item).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (m *AlertTemplate) UpdateTemplate(ctx context.Context, channels []uint64) error {
	var oldChannel []uint64
	err := GetPluck(ctx, mysql.ChannelTemplateTable, "channel_id", &oldChannel, WhereWithScope("template_id", m.Id))
	if err != nil {
		return err
	}
	var addChannel, deleteChannel []uint64
	for _, v := range channels {
		if !utils.InOfT(v, oldChannel) {
			addChannel = append(addChannel, v)
		}
	}

	for _, v := range oldChannel {
		if !utils.InOfT(v, channels) {
			deleteChannel = append(deleteChannel, v)
		}
	}
	return Transaction(ctx, func(tx *gorm.DB) error {
		err := tx.Table(m.TableName()).Save(&m).Error
		if err != nil {
			return err
		}
		for _, v := range addChannel {
			item := ChannelTemplate{
				TemplateId: m.Id,
				ChannelId:  v,
				Content:    "",
				Finished:   false,
			}
			err := tx.Table(item.TableName()).Create(&item).Error
			if err != nil {
				return err
			}
		}
		err = tx.Table(mysql.ChannelTemplateTable).Where("template_id", m.Id).Where("channel_id IN ?", deleteChannel).Delete(ChannelTemplate{}).Error
		if err != nil {
			return err
		}
		return nil
	})
}

// All 获取所有告警模板
func (m *AlertTemplate) All(ctx context.Context) ([]AlertTemplate, error) {
	var data []AlertTemplate
	err := GetAll(ctx, m.TableName(), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
