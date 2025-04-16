package model

import (
	"context"
	"darkhawk/app/biz/dal/mysql"
	"fmt"
	"gorm.io/gorm"
)

type WorkspaceSource struct {
	Model
	WorkspaceId   uint64 `gorm:"type:int(11);column:workspace_id;comment:工作区编号" json:"workspaceId" form:"workspaceId"`
	SourceId      uint64 `gorm:"type:int(11);column:source_id;comment:告警源编号" json:"sourceId" form:"sourceId"`
	SourceName    string `gorm:"column:source_name;comment:告警源名称" json:"sourceName" form:"sourceName"`
	Verified      bool   `gorm:"type:tinyint(1);default:true;comment:是否完成配置" json:"verified" form:"verified"`
	Configuration string `gorm:"type:json;column:configuration;comment:告警策略" json:"configuration" form:"configuration"`
	Token         string `gorm:"type:varchar(100);column:token;comment:token" json:"token"`
	LevelField    string `gorm:"type:varchar(100);column:level_field;comment:告警等级字段" json:"levelField"`
	Description   string `json:"description"`
}

type WorkspaceSourceItem struct {
	WorkspaceSource
	Icon       string `gorm:"type:varchar(100);column:icon;comment:图标" json:"icon"`
	EventCount int    `gorm:"type:int(11);column:eventCount;comment:告警事件数量" json:"eventCount"`
}

func (m *WorkspaceSource) TableName() string {
	return mysql.WorkspaceSourceTable
}

func (m *WorkspaceSource) Create(ctx context.Context, id uint64) error {
	var lsm []LevelSourceMapping
	levelModel := new(AlertLevel)
	levelSourceMappingModel := new(LevelSourceMapping)
	data, err := levelModel.All(ctx)

	if err != nil {
		return err
	}
	return mysql.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Table(m.TableName()).Create(&m).Error
		if err != nil {
			return err
		}
		for _, v := range data {
			l := LevelSourceMapping{
				LevelId:          v.Id,
				WorkspaceId:      id,
				WsId:             m.Id,
				AlertSourceLevel: v.LevelName,
			}
			lsm = append(lsm, l)
		}
		return tx.Table(levelSourceMappingModel.TableName()).Create(&lsm).Error
	})
}

// GetWorkspaceSource 获取工作区数据源
func (m *WorkspaceSource) GetWorkspaceSource(ctx context.Context, id uint64) ([]WorkspaceSourceItem, error) {
	var ws []WorkspaceSourceItem
	err := mysql.DB.WithContext(ctx).Table(m.TableName()).
		Select("workspace_source.*,alert_source.icon,t1.eventCount").
		Joins("LEFT JOIN alert_source ON alert_source.id = workspace_source.source_id").
		Joins(fmt.Sprintf("LEFT JOIN (SELECT COUNT(*) AS eventCount,source_id FROM `alert_event` WHERE workspace_id = %d GROUP BY source_id) AS t1 ON t1.source_id = workspace_source.id", id)).
		Where("workspace_source.workspace_id", id).Find(&ws).Error
	if err != nil {
		return ws, err
	}
	return ws, nil
}

func (m *WorkspaceSource) DeleteWorkspaceSource(ctx context.Context, id uint64) error {
	err := mysql.DB.WithContext(ctx).Table(m.TableName()).Where("id", id).Take(&m).Error
	if err != nil {
		return err
	}
	levelSourceMappingModel := new(LevelSourceMapping)
	return mysql.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Table(m.TableName()).Delete(&m).Error
		if err != nil {
			return err
		}
		return tx.Table(levelSourceMappingModel.TableName()).Where("ws_id", id).Delete(LevelSourceMapping{}).Error
	})

}
