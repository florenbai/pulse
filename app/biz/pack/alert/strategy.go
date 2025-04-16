package alert

import (
	"context"
	"darkhawk/app/biz/model"
	"darkhawk/app/biz/pack/notify"
	"darkhawk/pkg/utils"
	"darkhawk/pkg/zlog"
	"fmt"
	"time"
)

type Loop struct {
	Interval int64 `json:"interval"`
	Count    int64 `json:"count"`
}

type StrategyInfo struct {
	Channel  []string `json:"channel"`
	Template uint64   `json:"template"`
	Loop     *Loop    `json:"loop"`
}

// StrategyMatch 策略匹配
func StrategyMatch(customStrategy []model.AlertStrategy, systemStrategy model.SystemStrategy, event model.AlertEvent, level string) error {

	currentTime := time.Now()
	customCount := 0
	open := false
	for _, sitem := range customStrategy {
		zlog.Debug(fmt.Sprintf("告警指纹：%s ，开始匹配策略：%s ", event.FingerPrint, sitem.StrategyName))
		// 时间段策略
		if sitem.TimeSlot.Enable {

			// 当存在周期并且周期不匹配,直接进入下一个匹配策略
			if len(sitem.TimeSlot.Weeks) > 0 && !utils.InOfT(int64(currentTime.Weekday()), sitem.TimeSlot.Weeks) {
				zlog.Debug(fmt.Sprintf("告警指纹：%s ，策略：%s, 周期不匹配", event.FingerPrint, sitem.StrategyName))
				continue
			}
			// 只有同时设置了，开始时间和结束时间才会进行时间匹配
			if len(sitem.TimeSlot.Times) == 2 {
				startTime, err := time.ParseInLocation("15:04:05", sitem.TimeSlot.Times[0], currentTime.Location())
				if err != nil {
					return err
				}
				endTime, err := time.ParseInLocation("15:04:05", sitem.TimeSlot.Times[1], currentTime.Location())
				if err != nil {
					return err
				}
				finalStartTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), startTime.Hour(), startTime.Minute(), startTime.Second(), 0, currentTime.Location())
				finalEndTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), endTime.Hour(), endTime.Minute(), endTime.Second(), 0, currentTime.Location())
				if currentTime.Before(finalStartTime) || currentTime.After(finalEndTime) {
					zlog.Debug(fmt.Sprintf("告警指纹：%s ，策略：%s, 开始时间和结束时间不匹配", event.FingerPrint, sitem.StrategyName))
					continue
				}
			}
		}

		// 标签匹配
		if len(sitem.Filters) > 0 {
			zlog.Debug(fmt.Sprintf("告警指纹：%s ，匹配策略：%s ，开始标签匹配", event.FingerPrint, sitem.StrategyName))
			pass := false
			for _, v := range sitem.Filters {
				if HasTag(v.Values, event.Tags) {
					zlog.Debug(fmt.Sprintf("告警指纹：%s ，匹配策略：%s ，标签匹配成功：匹配标签：%v ， 匹配条件：%v ", event.FingerPrint, sitem.StrategyName, event.Tags, v.Values))
					pass = true
				}
			}
			if !pass {
				zlog.Debug(fmt.Sprintf("告警指纹：%s ，匹配策略：%s ，标签匹配不成功", event.FingerPrint, sitem.StrategyName))
				continue
			}
		}

		if len(sitem.StrategySet) == 0 {
			continue
		}
		if *sitem.SystemStrategy {
			open = true
		}
		customCount++
		err := strategySetMember(sitem.StrategySet, 0, level, sitem.TemplateId, sitem.Delay, event, sitem.Id, model.CustomStrategy)
		if err != nil {
			return err
		}
		// 未开启接续匹配，则直接退出
		if !*sitem.Continuous {
			break
		}
	}
	if (customCount > 0 && open) || len(customStrategy) == 0 {
		err := SystemStrategy(systemStrategy, level, event)
		if err != nil {
			return err
		}
	}
	return nil
}

// SystemStrategy 系统默认策略
func SystemStrategy(systemStrategy model.SystemStrategy, level string, event model.AlertEvent) error {
	return strategySetMember(systemStrategy.StrategySet, 0, level, systemStrategy.TemplateId, systemStrategy.Delay, event, systemStrategy.Id, model.DefaultStrategy)
}

// 策略分配人员
func strategySetMember(strategySet []model.StrategySet, k int, level string, template uint64, delay int64, event model.AlertEvent, id uint64, st int) error {
	levelModel := new(model.AlertLevel)
	levelObject, err := levelModel.All(context.TODO())
	if err != nil {
		return err
	}
	var errMessage string
	var isNotify = false
	strategyItem := strategySet[k]
	//告警等级匹配以及确认告警通道
	var channels []uint64
	for levelKey, cs := range strategyItem.AlertChannels {
		if level == levelObject[levelKey].LevelName {
			channels = cs
		}
	}
	notify := notify.Notify{
		Teams:     strategyItem.Teams,
		Members:   strategyItem.Members,
		Schedules: strategyItem.Schedules,
	}

	members, err := notify.GetNotifyMembers(context.TODO())
	if err != nil {
		return err
	}

	// 写入分派log
	defer func() {
		alertChannelModel := new(model.AlertChannel)
		channelNames, _ := alertChannelModel.GetAlertChannelName(context.TODO(), channels)
		for _, user := range members {
			slog := model.StrategyLog{
				AlertId:         event.Id,
				Uid:             user,
				StrategyContent: strategyItem,
				StrategyId:      id,
				Channels:        channelNames,
				IsNotify:        isNotify,
				ErrMessage:      errMessage,
				StrategyType:    st,
			}
			if !event.IsRecovered {
				slog.NotifyType = model.NotifyTypeAlert
			} else {
				slog.NotifyType = model.NotifyTypeRECOVE
			}
			model.Create(context.TODO(), slog.TableName(), &slog)
		}

	}()
	err = notify.Send(context.TODO(), template, members, channels, event)
	isNotify = true
	if err != nil {
		isNotify = false
		errMessage = err.Error()
		return err
	}
	// 循环通知
	if strategyItem.AlertLoop != nil && strategyItem.AlertLoop.Enable {
		go func() {
			for i := 0; i < int(strategyItem.AlertLoop.Count); i++ {
				time.Sleep(time.Minute * time.Duration(strategyItem.AlertLoop.Interval))
				err := model.GetOneById(context.TODO(), event.TableName(), event.Id, &event)
				if err != nil {
					zlog.Error(fmt.Sprintf("循环通知，获取告警信息异常：%s", err.Error()))
					return
				}
				// 已恢复和已关闭状态，则不进行循环通知
				if event.IsRecovered || event.Progress == model.EventClosed {
					return
				}
				notify.Send(context.TODO(), template, members, channels, event)
			}
		}()
	}
	// 故障升级
	if strategyItem.Progression != nil && strategyItem.Progression.Enable {
		go func() {
			time.Sleep(time.Minute * time.Duration(strategyItem.Progression.Duration))
			err := model.GetOneById(context.TODO(), event.TableName(), event.Id, &event)
			if err != nil {
				zlog.Error(fmt.Sprintf("故障升级，获取告警信息异常：%s", err.Error()))
				return
			}
			// 检测触发条件
			// 当未认领时，必须事件未认领否则为false，当条件为未关闭时，则处理事件必须不能为已关闭
			if (strategyItem.Progression.Condition == 0 && event.Progress != model.EventUnClaimed) || (strategyItem.Progression.Condition == 1 && event.Progress == model.EventClosed) {
				return
			}
			err = strategySetMember(strategySet, k+1, level, template, 0, event, id, st)
			if err != nil {
				zlog.Error(fmt.Sprintf("故障升级异常：%s", err.Error()))
				return
			}
		}()
	}
	return nil
}
