package model

import (
	"context"
	"darkhawk/app/biz/dal/mysql"
	"time"
)

type AlertLevel struct {
	Model
	LevelName string `gorm:"type:varchar(100);column:level_name;comment:告警源名称" json:"levelName" form:"levelName"`
	LevelDesc string `gorm:"type:varchar(100);column:level_desc;comment:告警源名称" json:"levelDesc" form:"levelDesc"`
	Color     string `gorm:"type:varchar(100);column:color;comment:图标" json:"color"`
	IsDefault bool   `gorm:"type:tinyint(1);column:is_default;comment:图标" json:"isDefault"`
}

const DefaultLevel string = "A3"

type LevelStatistic struct {
	Count     int64  `json:"count"`
	LevelName string `json:"levelName"`
	Color     string `json:"color"`
	Level     uint64 `json:"level"`
}

func (m *AlertLevel) TableName() string {
	return mysql.AlertLevelTable
}

func (m *AlertLevel) All(ctx context.Context) ([]AlertLevel, error) {
	var data []AlertLevel
	err := GetAll(ctx, m.TableName(), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (m *AlertLevel) LevelNameMap(ctx context.Context) (map[string]uint64, error) {
	list, err := m.All(ctx)
	if err != nil {
		return nil, err
	}
	var data = make(map[string]uint64)
	for _, v := range list {
		data[v.LevelName] = v.Id
	}
	return data, nil
}

func (m *AlertLevel) GetLevelStatistic(ctx context.Context, workspaceId uint64) ([]LevelStatistic, error) {
	var levelStatistic []LevelStatistic
	var data []LevelStatistic
	var levelData []AlertLevel
	var levelMap = make(map[uint64]int64)
	today := time.Now().Format("2006-01-02")
	tomorrow := time.Now().AddDate(0, 0, +1).Format("2006-01-02")
	sql := mysql.DB.WithContext(ctx).Table(mysql.AlertEventTable).Select("count(*) as count, level").Where("trigger_time >= ? AND trigger_time < ?", today, tomorrow)
	if workspaceId > 0 {
		sql.Joins("LEFT JOIN alert_workspace ON alert_workspace.alert_id = alert_event.id").Where("alert_workspace.workspace_id", workspaceId)
	}
	err := sql.Group("level").Find(&levelStatistic).Error
	if err != nil {
		return nil, err
	}
	for _, v := range levelStatistic {
		levelMap[v.Level] = v.Count
	}
	err = mysql.DB.WithContext(ctx).Table(m.TableName()).Find(&levelData).Error
	if err != nil {
		return nil, err
	}
	for _, v := range levelData {
		a := LevelStatistic{
			Color:     v.Color,
			Count:     0,
			Level:     v.Id,
			LevelName: v.LevelName,
		}
		if c, ok := levelMap[v.Id]; ok {
			a.Count = c
		}
		data = append(data, a)
	}

	return data, nil
}
