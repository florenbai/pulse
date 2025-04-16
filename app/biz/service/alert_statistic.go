package service

import (
	"context"
	"darkhawk/app/biz/model"
	"darkhawk/pkg/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"time"
)

type AlertStatisticService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewAlertStatisticService(ctx context.Context, c *app.RequestContext) *AlertStatisticService {
	return &AlertStatisticService{
		ctx: ctx,
		c:   c,
	}
}

type MttaResp struct {
	Today     float64 `json:"today"`
	Yesterday float64 `json:"yesterday"`
	Diff      float64 `json:"diff"`
	Direction string  `json:"direction"`
}

func (service *AlertStatisticService) MTTA(id uint64) (*MttaResp, error) {
	var yesterdayData []model.AlertEvent
	var todayData []model.AlertEvent
	alertEventModel := new(model.AlertEvent)
	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	today := time.Now().Format("2006-01-02")
	tomorrow := time.Now().AddDate(0, 0, +1).Format("2006-01-02")
	yesterdayScopes := model.NewScopeContainer()
	todayScopes := model.NewScopeContainer()
	yesterdayScopes = append(yesterdayScopes, model.WhereWithScope("progress IN ?", []model.AlertStatus{model.EventHasClaimed, model.EventClosed}), model.WhereWithScope("trigger_time >= ?", yesterday), model.WhereWithScope("trigger_time < ?", today))
	todayScopes = append(todayScopes, model.WhereWithScope("progress IN ?", []model.AlertStatus{model.EventHasClaimed, model.EventClosed}), model.WhereWithScope("trigger_time >= ?", today), model.WhereWithScope("trigger_time < ?", tomorrow))
	if id > 0 {
		yesterdayScopes = append(yesterdayScopes, model.JoinScope("LEFT JOIN alert_workspace ON alert_workspace.alert_id = alert_event.id"), model.WhereWithScope("alert_workspace.workspace_id", id))
		todayScopes = append(todayScopes, model.JoinScope("LEFT JOIN alert_workspace ON alert_workspace.alert_id = alert_event.id"), model.WhereWithScope("alert_workspace.workspace_id", id))
	}

	err := model.GetAll(service.ctx, alertEventModel.TableName(), &yesterdayData, yesterdayScopes...)
	if err != nil {
		return nil, err
	}
	err = model.GetAll(service.ctx, alertEventModel.TableName(), &todayData, todayScopes...)
	if err != nil {
		return nil, err
	}
	var todayMinutes, yesterdayMinutes int
	var todayCount, yesterdayCount = 0, 0

	for _, v := range todayData {
		if v.FirstAckTime != nil {
			duration := time.Time(*v.FirstAckTime).Sub(time.Time(*v.FirstTriggerTime))
			todayMinutes += int(duration.Minutes())
			todayCount++
		}
	}
	var data MttaResp
	for _, v := range yesterdayData {
		if v.FirstAckTime != nil {
			duration := time.Time(*v.FirstAckTime).Sub(time.Time(*v.FirstTriggerTime))
			yesterdayMinutes += int(duration.Minutes())
			yesterdayCount++
		}

	}

	if todayCount > 0 {
		data.Today = float64(todayMinutes / todayCount)
	}
	if yesterdayCount > 0 {
		data.Yesterday = float64(yesterdayMinutes / yesterdayCount)
	}
	if data.Today >= data.Yesterday {
		data.Direction = "up"
		data.Diff = data.Today - data.Yesterday
	} else {
		data.Direction = "down"
		data.Diff = data.Yesterday - data.Today
	}

	return &data, nil
}

func (service *AlertStatisticService) MTTR(id uint64) (*MttaResp, error) {
	var yesterdayData []model.AlertEvent
	var todayData []model.AlertEvent
	alertEventModel := new(model.AlertEvent)
	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	today := time.Now().Format("2006-01-02")
	tomorrow := time.Now().AddDate(0, 0, +1).Format("2006-01-02")
	yesterdayScopes := model.NewScopeContainer()
	todayScopes := model.NewScopeContainer()
	yesterdayScopes = append(yesterdayScopes, model.WhereWithScope("trigger_time >= ?", yesterday), model.WhereWithScope("trigger_time < ?", today))
	todayScopes = append(todayScopes, model.WhereWithScope("trigger_time >= ?", today), model.WhereWithScope("trigger_time < ?", tomorrow))
	if id > 0 {
		yesterdayScopes = append(yesterdayScopes, model.JoinScope("LEFT JOIN alert_workspace ON alert_workspace.alert_id = alert_event.id"), model.WhereWithScope("alert_workspace.workspace_id", id))
		todayScopes = append(todayScopes, model.JoinScope("LEFT JOIN alert_workspace ON alert_workspace.alert_id = alert_event.id"), model.WhereWithScope("alert_workspace.workspace_id", id))
	}
	err := model.GetAll(service.ctx, alertEventModel.TableName(), &yesterdayData, yesterdayScopes...)
	if err != nil {
		return nil, err
	}
	err = model.GetAll(service.ctx, alertEventModel.TableName(), &todayData, todayScopes...)
	if err != nil {
		return nil, err
	}
	var todayMinutes, yesterdayMinutes int
	var todayCount, yesterdayCount = 0, 0

	for _, v := range todayData {
		if v.IsRecovered {
			duration := time.Time(*v.RecoverTime).Sub(time.Time(*v.FirstTriggerTime))
			todayMinutes += int(duration.Minutes())
		} else {
			duration := time.Now().Sub(time.Time(*v.FirstTriggerTime))
			todayMinutes += int(duration.Minutes())
		}
		todayCount++
	}
	var data MttaResp
	for _, v := range yesterdayData {
		if v.IsRecovered {
			duration := time.Time(*v.RecoverTime).Sub(time.Time(*v.FirstTriggerTime))
			yesterdayMinutes += int(duration.Minutes())
		} else {
			duration := time.Now().Sub(time.Time(*v.FirstTriggerTime))
			yesterdayMinutes += int(duration.Minutes())
		}
		yesterdayCount++
	}

	if todayCount > 0 {
		data.Today = float64(todayMinutes / todayCount)
	}
	if yesterdayCount > 0 {
		data.Yesterday = float64(yesterdayMinutes / yesterdayCount)
	}
	if data.Today >= data.Yesterday {
		data.Direction = "up"
		data.Diff = data.Today - data.Yesterday
	} else {
		data.Direction = "down"
		data.Diff = data.Yesterday - data.Today
	}

	return &data, nil
}

func (service *AlertStatisticService) Count(id uint64) (*MttaResp, error) {
	var todayCount, yesterdayCount int64
	var data MttaResp
	alertEventModel := new(model.AlertEvent)
	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	today := time.Now().Format("2006-01-02")
	tomorrow := time.Now().AddDate(0, 0, +1).Format("2006-01-02")
	yesterdayScopes := model.NewScopeContainer()
	todayScopes := model.NewScopeContainer()
	yesterdayScopes = append(yesterdayScopes, model.WhereWithScope("trigger_time >= ?", yesterday), model.WhereWithScope("trigger_time < ?", today))
	todayScopes = append(todayScopes, model.WhereWithScope("trigger_time >= ?", today), model.WhereWithScope("trigger_time < ?", tomorrow))
	if id > 0 {
		yesterdayScopes = append(yesterdayScopes, model.JoinScope("LEFT JOIN alert_workspace ON alert_workspace.alert_id = alert_event.id"), model.WhereWithScope("alert_workspace.workspace_id", id))
		todayScopes = append(todayScopes, model.JoinScope("LEFT JOIN alert_workspace ON alert_workspace.alert_id = alert_event.id"), model.WhereWithScope("alert_workspace.workspace_id", id))
	}
	yesterdayCount, err := model.GetCountWithScope(service.ctx, alertEventModel.TableName(), yesterdayScopes...)
	if err != nil {
		return nil, err
	}
	todayCount, err = model.GetCountWithScope(service.ctx, alertEventModel.TableName(), todayScopes...)
	if err != nil {
		return nil, err
	}
	data.Today = float64(todayCount)
	data.Yesterday = float64(yesterdayCount)
	if data.Today >= data.Yesterday {
		data.Direction = "up"
		data.Diff = data.Today - data.Yesterday
	} else {
		data.Direction = "down"
		data.Diff = data.Yesterday - data.Today
	}
	return &data, nil
}

func (service *AlertStatisticService) Level(id uint64) ([]model.LevelStatistic, error) {
	alertLevel := new(model.AlertLevel)
	data, err := alertLevel.GetLevelStatistic(service.ctx, id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (service *AlertStatisticService) TopTen(timeRange string) ([]model.TopStatistic, error) {
	alertEventModel := new(model.AlertEvent)
	scopes := model.NewScopeContainer()
	if timeRange != "" {
		tg := utils.GetTimeAgo(timeRange).Format("2006-01-02 15:04:05")
		scopes = append(scopes, model.WhereWithScope("first_trigger_time >= ? ", tg))
	}
	return alertEventModel.GetTopTenStatistic(service.ctx, scopes)
}

func (service *AlertStatisticService) LevelAlert(timeRange string) ([]model.LevelAlertStatistic, error) {
	alertEventModel := new(model.AlertEvent)
	scopes := model.NewScopeContainer()
	if timeRange != "" {
		tg := utils.GetTimeAgo(timeRange).Format("2006-01-02 15:04:05")
		scopes = append(scopes, model.WhereWithScope("first_trigger_time >= ? ", tg))
	}
	return alertEventModel.GetLevelAlertTenStatistic(service.ctx, scopes)
}

func (service *AlertStatisticService) AlertCount(timeRange string) ([]model.AlertCountStatistic, error) {
	alertEventModel := new(model.AlertEvent)
	scopes := model.NewScopeContainer()
	if timeRange != "" {
		tg := utils.GetTimeAgo(timeRange).Format("2006-01-02 15:04:05")
		scopes = append(scopes, model.WhereWithScope("first_trigger_time >= ? ", tg))
	}
	return alertEventModel.GetAlertCountStatistic(service.ctx, scopes)
}
