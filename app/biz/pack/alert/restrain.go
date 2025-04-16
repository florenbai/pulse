package alert

import (
	"context"
	"darkhawk/app/biz/dal/mysql"
	"darkhawk/app/biz/model"
	"darkhawk/pkg/zlog"
	"fmt"
	"time"
)

func (alert *EventRequest) CheckRestrain(ctx context.Context, workspaceId uint64) (bool, error) {
	var err error
	if alert.Status == model.AlertResolved {
		return false, nil
	}
	var data []model.AlertRestrain
	alertRestrainModel := new(model.AlertRestrain)
	err = model.GetAll(ctx, alertRestrainModel.TableName(), &data, model.WhereWithScope("workspace_id", workspaceId))
	if err != nil {
		return false, err
	}
	ok := false

	for _, v := range data {
		zlog.Debug(fmt.Sprintf("告警指纹：%s ，告警状态: %s, 告警抑制编号：%d", alert.FingerPrint, alert.Status, v.Id))
		switch v.RestrainType {
		case "fingerprint":
			ok, err = alert.fingerprint(ctx, v)
		case "label":
			ok, err = alert.label(ctx, v)
		case "annotation":
			ok, err = alert.annotation(ctx, v)
		}
		if err != nil {
			return false, err
		}
		if ok {
			return ok, nil
		}
	}
	return ok, err
}

func (alert *EventRequest) annotation(ctx context.Context, restrain model.AlertRestrain) (bool, error) {
	cumulativeDateTime := time.Now().Add(-time.Minute * time.Duration(restrain.CumulativeTime)).Format("2006-01-02 15:04:05")
	eventModel := new(model.AlertEvent)
	var data []model.AlertEvent
	err := model.GetAll(ctx, eventModel.TableName(), &data, model.WhereWithScope("trigger_time > ?", cumulativeDateTime), model.WhereWithScope("progress IN ?", []model.AlertStatus{model.EventUnClaimed, model.EventHasClaimed}))
	if err != nil {
		return false, err
	}
	pass := false
	for _, event := range data {
		for _, field := range restrain.Fields {
			// 当，当前告警存在tag时
			if v, ok := alert.Labels[field]; ok {
				tagValue := event.Annotations[field]
				if v == tagValue {
					zlog.Debug(fmt.Sprintf("告警指纹：%s ，告警状态: %s, 告警注解匹配成功，匹配注解：%s ，值：%s", alert.FingerPrint, alert.Status, field, v))
					pass = true
				} else {
					pass = false
				}
			} else {
				pass = false
			}
		}
	}
	return pass, nil
}

func (alert *EventRequest) label(ctx context.Context, restrain model.AlertRestrain) (bool, error) {

	cumulativeDateTime := time.Now().Add(-time.Minute * time.Duration(restrain.CumulativeTime)).Format("2006-01-02 15:04:05")
	eventModel := new(model.AlertEvent)
	var data []model.AlertEvent
	err := model.GetAll(ctx, eventModel.TableName(), &data, model.WhereWithScope("trigger_time > ?", cumulativeDateTime), model.WhereWithScope("progress IN ?", []model.AlertStatus{model.EventUnClaimed, model.EventHasClaimed}))
	if err != nil {
		return false, err
	}
	pass := false
	for _, event := range data {
		for _, field := range restrain.Fields {
			// 当，当前告警存在tag时
			if v, ok := alert.Labels[field]; ok {
				tagValue := event.Tags[field]
				if v == tagValue {
					zlog.Debug(fmt.Sprintf("告警指纹：%s ，告警状态: %s, 告警标签匹配成功，匹配标签：%s ，值：%s", alert.FingerPrint, alert.Status, field, v))
					pass = true
				} else {
					pass = false
				}
			} else {
				pass = false
			}
		}
	}
	return pass, nil
}

func (alert *EventRequest) fingerprint(ctx context.Context, restrain model.AlertRestrain) (bool, error) {
	var c int64
	cumulativeDateTime := time.Now().Add(-time.Minute * time.Duration(restrain.CumulativeTime)).Format("2006-01-02 15:04:05")
	strategyLogModel := new(model.StrategyLog)
	subQuery := mysql.DB.WithContext(ctx).Table(strategyLogModel.TableName()).Select("alert_id,MAX(strategy_log.created_at) AS created_at").Joins("LEFT JOIN `alert_event` ON alert_event.id = strategy_log.alert_id").Where("alert_event.finger_print", alert.FingerPrint).Where("progress IN ?", []model.AlertStatus{model.EventUnClaimed, model.EventHasClaimed})
	err := mysql.DB.WithContext(ctx).Table("(?) as u", subQuery).Where("u.created_at > ?", cumulativeDateTime).Count(&c).Error
	if err != nil {
		return false, err
	}
	if c > 0 {
		zlog.Debug(fmt.Sprintf("告警指纹：%s ，告警状态: %s, 告警抑制编号：%d, 告警指纹匹配成功，", alert.FingerPrint, alert.Status, restrain.Id))
		return true, nil
	}
	return false, err
}
