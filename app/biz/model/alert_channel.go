package model

import (
	"context"
	"darkhawk/app/biz/dal/mysql"
)

type AlertChannel struct {
	Model
	ChannelName  string `gorm:"column:channel_name" json:"channelName"`
	ChannelType  string `gorm:"column:channel_type" json:"channelType"`
	ChannelSign  string `gorm:"column:channel_sign" json:"channelSign"`
	ChannelGroup string `gorm:"column:channel_group" json:"channelGroup"`
}

func (m *AlertChannel) TableName() string {
	return mysql.AlertChannelTable
}

func (m *AlertChannel) GetAlertChannelName(ctx context.Context, ids []uint64) ([]string, error) {
	var data []string
	err := mysql.DB.WithContext(ctx).Table(m.TableName()).Where("id IN ?", ids).Pluck("channel_name", &data).Error
	if err != nil {
		return data, err
	}
	return data, nil
}
