package service

import (
	"context"
	"darkhawk/app/biz/model"
	"darkhawk/pkg/utils"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
)

type AlertLogService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewAlertLogService(ctx context.Context, c *app.RequestContext) *AlertLogService {
	return &AlertLogService{
		ctx: ctx,
		c:   c,
	}
}

func (service *AlertLogService) GetAlertTimeLineLog(id uint64) ([]model.AlertLogList, error) {
	alertLogModel := new(model.AlertLog)
	return alertLogModel.GetAlertLogByEvent(service.ctx, id)
}

type StrategyLogData struct {
	StrategyTitle string   `json:"strategyTitle"`
	CreatedAt     string   `json:"createdAt"`
	Nicknames     []string `json:"nicknames"`
	Channels      []string `json:"channels"`
	NotifyType    int      `json:"notifyType"`
}

// GetStrategyLog 获取分派日志
func (service *AlertLogService) GetStrategyLog(id uint64) ([]StrategyLogData, error) {
	strategyLogModel := new(model.StrategyLog)
	list, err := strategyLogModel.GetStrategyLogByEvent(service.ctx, id)
	if err != nil {
		return nil, err
	}
	alertStrategyModel := new(model.AlertStrategy)
	systemStrategyModel := new(model.SystemStrategy)
	strategyMap, err := alertStrategyModel.GetStrategyMap(service.ctx)
	if err != nil {
		return nil, err
	}
	systemStrategyMap, err := systemStrategyModel.GetStrategyMap(service.ctx)
	if err != nil {
		return nil, err
	}
	for k, v := range list {
		if v.StrategyType == model.DefaultStrategy {
			list[k].StrategyTitle = systemStrategyMap[v.StrategyId]
		} else {
			list[k].StrategyTitle = strategyMap[v.StrategyId]
		}
	}
	var data []StrategyLogData
	var d = make(map[string]StrategyLogData)
	for _, v := range list {
		k := fmt.Sprintf("%s-%s-%d", v.StrategyTitle, v.CreatedAt, v.NotifyType)
		if dk, ok := d[k]; ok {
			dk.Nicknames = append(dk.Nicknames, v.Nickname)
			dk.Channels = append(dk.Channels, v.Channels...)
			dk.Channels = utils.RemoveDuplicate(dk.Channels)
			d[k] = dk
		} else {
			d[k] = StrategyLogData{
				StrategyTitle: v.StrategyTitle,
				CreatedAt:     v.CreatedAt.Format(),
				Nicknames:     []string{v.Nickname},
				Channels:      v.Channels,
				NotifyType:    v.NotifyType,
			}
		}
	}

	for _, v := range d {
		data = append(data, v)
	}
	return data, nil
}
