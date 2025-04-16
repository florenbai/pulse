package model

import (
	"context"
	"darkhawk/app/biz/dal/mysql"
	"errors"
	"gorm.io/gorm"
	"time"
)

type AlertStatus uint8

const (
	EventUnClaimed  AlertStatus = 1
	EventHasClaimed AlertStatus = 2
	EventClosed     AlertStatus = 3
)

const (
	AlertFiring   string = "firing"
	AlertResolved string = "resolved"
)

type AlertEvent struct {
	Model
	AlertTitle       string            `gorm:"column:alert_title" json:"alertTitle"`
	Description      string            `gorm:"column:description" json:"description"`
	SourceId         uint64            `gorm:"column:source_id" json:"sourceId"`
	Level            uint64            `gorm:"column:level" json:"level"`
	FirstTriggerTime *LocalTime        `gorm:"column:first_trigger_time" json:"firstTriggerTime"`
	FirstAckTime     *LocalTime        `gorm:"column:first_ack_time" json:"firstAckTime"`
	TriggerTime      *LocalTime        `gorm:"column:trigger_time" json:"triggerTime"`
	RecoverTime      *LocalTime        `gorm:"column:recover_time" json:"recoverTime"`
	Annotations      map[string]string `gorm:"column:annotations;serializer:json" json:"annotations"`
	IsRecovered      bool              `gorm:"column:is_recovered" json:"isRecovered"`
	NotifyCurNumber  uint64            `gorm:"column:notify_cur_number" json:"notifyCurNumber"`
	Progress         AlertStatus       `gorm:"column:progress" json:"progress"`
	WorkspaceId      uint64            `gorm:"column:workspace_id" json:"workspaceId"`
	IntegrationId    uint64            `gorm:"column:integration_id" json:"integrationId"`
	SourceIp         string            `gorm:"column:source_ip" json:"sourceIp"`
	Tags             map[string]string `gorm:"column:tags;serializer:json" json:"tags"`
	Uid              uint64            `json:"uid"`
	FingerPrint      string            `gorm:"column:finger_print" json:"fingerPrint"`
}

type WorkspaceAlertEvent struct {
	Id               uint64            `gorm:"column:id;primary_key;comment:主键id" json:"id" form:"id"` // 主键
	AlertTitle       string            `gorm:"column:alert_title" json:"alertTitle"`
	Description      string            `gorm:"column:description" json:"description"`
	SourceId         uint64            `gorm:"column:source_id" json:"sourceId"`
	Level            uint64            `gorm:"column:level" json:"level"`
	FirstTriggerTime *LocalTime        `gorm:"column:first_trigger_time" json:"firstTriggerTime"`
	FirstAckTime     *LocalTime        `gorm:"column:first_ack_time" json:"firstAckTime"`
	TriggerTime      *LocalTime        `gorm:"column:trigger_time" json:"triggerTime"`
	RecoverTime      *LocalTime        `gorm:"column:recover_time" json:"recoverTime"`
	Annotations      map[string]string `gorm:"column:annotations;serializer:json" json:"annotations"`
	IsRecovered      bool              `gorm:"column:is_recovered" json:"isRecovered"`
	NotifyCurNumber  uint64            `gorm:"column:notify_cur_number" json:"notifyCurNumber"`
	Progress         AlertStatus       `gorm:"column:progress" json:"progress"`
	WorkspaceId      uint64            `gorm:"column:workspace_id" json:"workspaceId"`
	SourceIp         string            `gorm:"column:source_ip" json:"sourceIp"`
	Tags             map[string]string `gorm:"column:tags;serializer:json" json:"tags"`
	Uid              uint64            `json:"uid"`
	LevelName        string            `gorm:"type:varchar(100);column:level_name;comment:告警等级" json:"levelName" form:"levelName"`
	Color            string            `gorm:"type:varchar(100);column:color;comment:图标" json:"color"`
	Icon             string            `gorm:"type:varchar(100);column:icon;comment:图标" json:"icon"`
	SourceName       string            `gorm:"type:varchar(100);column:source_name;comment:告警源名称" json:"sourceName" form:"sourceName"`
	Nickname         string            `json:"nickname"`
	CreatedAt        *LocalTime        `gorm:"column:created_at;not null;type:datetime;default:current_timestamp;comment:创建时间" json:"createdAt" form:"createdAt" copier:"must"`
	UpdatedAt        *LocalTime        `gorm:"column:updated_at;not null;type:datetime;default:current_timestamp on update CURRENT_TIMESTAMP;comment:更新时间" json:"updatedAt"  form:"updatedAt"`
}

type TopStatistic struct {
	AlertTitle string `gorm:"column:alert_title" json:"alertTitle"`
	Num        uint64 `json:"num"`
}

type LevelAlertStatistic struct {
	LevelName string `gorm:"type:varchar(100);column:level_name;comment:告警等级" json:"levelName" form:"levelName"`
	Color     string `gorm:"type:varchar(100);column:color;comment:图标" json:"color"`
	Count     uint64 `json:"count"`
}

type AlertCountStatistic struct {
	Date *LocalDate `json:"date"`
	Num  uint64     `json:"num"`
}

func (m *AlertEvent) TableName() string {
	return mysql.AlertEventTable
}

// GetWorkspaceEventList 获取工作区的告警事件
func (m *AlertEvent) GetWorkspaceEventList(ctx context.Context, workspaceId uint64, page uint64, pageSize uint64, scopes ...func(*gorm.DB) *gorm.DB) ([]WorkspaceAlertEvent, int64, error) {
	var i int64
	var data []WorkspaceAlertEvent
	scopes = append(scopes, WhereWithScope("alert_workspace.workspace_id", workspaceId))
	err := mysql.DB.WithContext(ctx).Table(m.TableName()).Joins("LEFT JOIN alert_workspace ON alert_event.id = alert_workspace.alert_id").Scopes(scopes...).Count(&i).Error
	if err != nil {
		return nil, 0, err
	}
	scopes = append(scopes, OrderScope("alert_event.trigger_time DESC"), Paginate(page, pageSize))
	err = mysql.DB.WithContext(ctx).Table(m.TableName()).
		Select("alert_event.*,t.level_name,t.color,system_integration.name as source_name,alert_source.icon,user.nickname").
		Joins("LEFT JOIN alert_workspace ON alert_event.id = alert_workspace.alert_id").
		Joins("LEFT JOIN alert_level AS t ON alert_event.level = t.id").
		Joins("LEFT JOIN user ON alert_event.uid = user.id").
		Joins("LEFT JOIN system_integration ON alert_event.integration_id = system_integration.id AND alert_event.source_id = system_integration.source_id").
		Joins("LEFT JOIN alert_source ON system_integration.source_id = alert_source.id").Scopes(scopes...).Find(&data).Error
	if err != nil {
		return nil, 0, err
	}

	return data, i, nil
}

// GetEventList 获取工作区的告警事件
func (m *AlertEvent) GetEventList(ctx context.Context, page uint64, pageSize uint64, scopes ...func(*gorm.DB) *gorm.DB) ([]WorkspaceAlertEvent, int64, error) {
	var i int64
	var data []WorkspaceAlertEvent
	err := mysql.DB.WithContext(ctx).Table(m.TableName()).Scopes(scopes...).Count(&i).Error
	if err != nil {
		return nil, 0, err
	}
	scopes = append(scopes, OrderScope("alert_event.trigger_time DESC"), Paginate(page, pageSize))
	err = mysql.DB.WithContext(ctx).Table(m.TableName()).
		Select("alert_event.*,t.level_name,t.color,system_integration.name as source_name,alert_source.icon,user.nickname").
		Joins("LEFT JOIN alert_level AS t ON alert_event.level = t.id").
		Joins("LEFT JOIN user ON alert_event.uid = user.id").
		Joins("LEFT JOIN system_integration ON alert_event.integration_id = system_integration.id").
		Joins("LEFT JOIN alert_source ON system_integration.source_id = alert_source.id").Scopes(scopes...).Find(&data).Error
	if err != nil {
		return nil, 0, err
	}

	return data, i, nil
}

// CheckCloseUser 检测是否可以关闭
func (m *AlertEvent) CheckCloseUser(ctx context.Context, id uint64, uid uint64) (bool, error) {
	var i int64
	err := mysql.DB.WithContext(ctx).Table(m.TableName()).Where("progress = ? AND uid = ? AND id = ?", EventHasClaimed, uid, id).Or("id = ? AND progress = ?", id, EventUnClaimed).Count(&i).Error
	if err != nil {
		return false, err
	}
	if i > 0 {
		return true, nil
	}
	return false, err
}

// CloseAlertEvent 关闭告警策略
func (m *AlertEvent) CloseAlertEvent(ctx context.Context, ids []uint64, IsRecovered bool, uid uint64) error {
	return Transaction(ctx, func(tx *gorm.DB) error {
		for _, id := range ids {
			if IsRecovered {
				err := tx.Table(m.TableName()).Where("id", id).Update("progress", EventClosed).Update("uid", uid).Update("is_recovered", true).Update("recover_time", time.Now()).Error
				if err != nil {
					return err
				}
			} else {
				err := tx.Table(m.TableName()).Where("id", id).Update("progress", EventClosed).Update("uid", uid).Error
				if err != nil {
					return err
				}
			}

			alertLog := AlertLog{
				AlertId: id,
				Action:  AlertClosed,
				Uid:     uid,
			}
			err := tx.Table(alertLog.TableName()).Create(&alertLog).Error
			if err != nil {
				return err
			}
		}
		return nil
	})

}

// ClaimAlertEvent 领取告警事件
func (m *AlertEvent) ClaimAlertEvent(ctx context.Context, ids []uint64, uid uint64) error {
	return Transaction(ctx, func(tx *gorm.DB) error {
		cuTime := LocalTime(time.Now())
		for _, id := range ids {
			err := tx.Table(m.TableName()).Where("id", id).First(&m).Error
			if err != nil {
				return err
			}
			m.Progress = EventHasClaimed
			m.Uid = uid
			m.FirstAckTime = &cuTime
			err = tx.Table(m.TableName()).Save(&m).Error
			if err != nil {
				return err
			}
			alertLog := AlertLog{
				AlertId: id,
				Action:  AlertClaim,
				Uid:     uid,
			}
			err = tx.Table(alertLog.TableName()).Create(&alertLog).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
}

// GetEnCloseEvent 获取未关闭的告警事件
func (m *AlertEvent) GetEnCloseEvent(ctx context.Context, id uint64) ([]AlertEvent, error) {
	var data []AlertEvent
	err := mysql.DB.WithContext(ctx).Table(m.TableName()).Joins("LEFT JOIN alert_workspace ON alert_workspace.alert_id = alert_event.id").Where("alert_workspace.workspace_id", id).Where("progress IN ?", []AlertStatus{EventHasClaimed, EventUnClaimed}).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (m *AlertEvent) IncrAlertEvent(ctx context.Context, id uint64) error {
	return mysql.DB.WithContext(ctx).Table(m.TableName()).Where("id", id).Update("notify_cur_number", gorm.Expr("notify_cur_number + ?", 1)).Error
}

// CreateAlertEvent 创建告警事件
func (m *AlertEvent) CreateAlertEvent(ctx context.Context, workspaces []uint64, IsRecovered bool) (err error) {
	// 获取前置告警事件未恢复的告警事件
	description := m.Description
	annotation := m.Annotations
	err = GetOneWithScope(ctx, m.TableName(), &m,
		WhereWithScope("finger_print", m.FingerPrint),
		InWithScope("progress", []AlertStatus{EventHasClaimed, EventUnClaimed}))
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	ct := LocalTime(time.Now())

	if m.Id > 0 {
		// 如果告警状态为恢复,设置状态并设置恢复时间
		if IsRecovered {
			m.RecoverTime = &ct
			m.Progress = EventClosed
			m.IsRecovered = IsRecovered
		} else {
			// 如果告警状态未恢复，则累加通知次数和最新告警时间
			m.NotifyCurNumber = m.NotifyCurNumber + 1
			m.TriggerTime = &ct
		}
		m.Description = description
		m.Annotations = annotation
		err = EditByScopes(ctx, m.TableName(), &m, WhereWithScope("id", m.Id))
		if err != nil {
			return err
		}
	} else {
		if !IsRecovered {
			m.Progress = EventUnClaimed
			// 创建新的告警事件
			m.NotifyCurNumber = 1
			m.TriggerTime = &ct
			m.FirstTriggerTime = &ct
			err := Transaction(ctx, func(tx *gorm.DB) error {
				err := tx.Table(m.TableName()).Create(&m).Error
				if err != nil {
					return err
				}
				alertWorkspaceModel := new(AlertWorkspace)
				for _, w := range workspaces {
					err := tx.Table(alertWorkspaceModel.TableName()).Create(&AlertWorkspace{
						WorkspaceId: w,
						AlertId:     m.Id,
					}).Error
					if err != nil {
						return err
					}
				}
				return nil
			})
			if err != nil {
				return err
			}
		} else {
			return errors.New("未发现告警中的告警信息")
		}
	}
	return
}

func (m *AlertEvent) GetAlertEventDetail(ctx context.Context, id uint64) (*WorkspaceAlertEvent, error) {
	var data WorkspaceAlertEvent
	err := mysql.DB.WithContext(ctx).Table(m.TableName()).
		Select("alert_event.*,alert_level.level_name,alert_source.source_name").
		Joins("LEFT JOIN alert_level ON alert_level.id = alert_event.level").
		Joins("LEFT JOIN alert_source ON alert_event.source_id = alert_source.id").
		Where("alert_event.id", id).First(&data).Error

	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (m *AlertEvent) GetTopTenStatistic(ctx context.Context, scopes []func(db *gorm.DB) *gorm.DB) ([]TopStatistic, error) {
	var data []TopStatistic
	err := mysql.DB.WithContext(ctx).Table(m.TableName()).Select("SUM(notify_cur_number) AS num,alert_title").Scopes(scopes...).Group("alert_title,finger_print").Order("num DESC").Limit(10).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (m *AlertEvent) GetLevelAlertTenStatistic(ctx context.Context, scopes []func(db *gorm.DB) *gorm.DB) ([]LevelAlertStatistic, error) {
	var data []LevelAlertStatistic
	err := mysql.DB.WithContext(ctx).Table(m.TableName()).Select("SUM(notify_cur_number) AS count,alert_level.level_name,alert_level.color").Joins("LEFT JOIN alert_level ON alert_level.id = alert_event.level").Scopes(scopes...).Group("level").Order("level").Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (m *AlertEvent) GetAlertCountStatistic(ctx context.Context, scopes []func(db *gorm.DB) *gorm.DB) ([]AlertCountStatistic, error) {
	var data []AlertCountStatistic
	err := mysql.DB.WithContext(ctx).Table(m.TableName()).Select("SUM(notify_cur_number) AS num,DATE(first_trigger_time) AS date").Scopes(scopes...).Group("DATE(first_trigger_time)").Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}
