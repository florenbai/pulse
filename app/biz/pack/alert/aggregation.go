package alert

import (
	"context"
	"darkhawk/app/biz/dal/redis"
	"darkhawk/app/biz/model"
	"darkhawk/pkg/zlog"
	"fmt"
	"time"
)

type Alert struct {
	Name  string
	Label map[string]string
	Level uint64
}

type AggregationEvent struct {
	Notify    bool
	OutWindow bool
}

// Aggregation 聚合窗口匹配
func (alert *EventRequest) Aggregation(ctx context.Context, workspaceId uint64) (*AggregationEvent, error) {
	// 获取可用的聚合窗口
	aggModel := new(model.AlertAggregation)
	aggs, err := aggModel.GetEnableAggregationByWorkspace(ctx, workspaceId)
	if err != nil {
		return nil, err
	}

	zlog.Debug(fmt.Sprintf("工作区：%d, 成功获取已开启聚合设置 %d 个", workspaceId, len(aggs)))
	// 获取未关闭告警事件
	eventModel := new(model.AlertEvent)
	events, err := eventModel.GetEnCloseEvent(ctx, workspaceId)
	zlog.Debug(fmt.Sprintf("工作区：%d, 成功获取未关闭告警事件 %v 个， 准备开始检测聚合设置", workspaceId, len(events)))
	if err != nil {
		return nil, err
	}
	for _, event := range events {
		for _, v := range aggs {
			// 检查未关闭事件是否符合聚合
			ae, err := alert.AggregationEvent(v, event)
			if err != nil {
				return nil, err
			}
			if ae != nil {
				return ae, nil
			}
		}
	}
	return nil, nil
}

// AggregationEvent 未关闭事件匹配，返回是否通知以及窗口
func (alert *EventRequest) AggregationEvent(agg model.AlertAggregation, event model.AlertEvent) (*AggregationEvent, error) {
	// 标题属性
	if agg.TitleDimension {
		if alert.AlertTitle != event.AlertTitle {
			zlog.Debug(fmt.Sprintf("聚合匹配标题不通过，告警标题：%s,匹配未关闭告警编号：%d", alert.AlertTitle, event.Id))
			return nil, nil
		}
	}
	// 等级属性
	if agg.LevelDimension {
		if alert.LevelId != event.Level {
			zlog.Debug(fmt.Sprintf("聚合匹配告警等级不通过，告警标题：%s,匹配未关闭告警编号：%d", alert.AlertTitle, event.Id))
			return nil, nil
		}
	}

	// 标签属性
	if len(agg.TagsDimension) > 0 {
		pass := false
		for _, v := range agg.TagsDimension {
			if HasTag(v.Values, event.Tags) {
				pass = true
				break
			}
		}
		if !pass {
			zlog.Debug(fmt.Sprintf("聚合匹配标签属性不通过，告警标题：%s,匹配未关闭告警编号：%d", alert.AlertTitle, event.Id))
			return nil, nil
		}
	}

	// 聚合窗口检测
	if agg.Windows > 0 {
		var ae AggregationEvent
		windowKey := fmt.Sprintf("window_%d", event.Id)
		ok, err := redis.Rdb.Exists(context.TODO(), windowKey).Result()
		if err != nil {
			return nil, err
		}
		// 当聚合窗口存在，则对告警进行聚合。当窗口不存在则创建新的告警，并生成窗口
		if ok > 0 {
			// 当存在窗口 检测风暴预警, 当聚合数小于风暴阈值时，则进行聚合累加，并返回聚合成功
			/*
				eventModel := new(model.AlertEvent)
				err = eventModel.IncrAlertEvent(context.TODO(), event.Id)
				if err != nil {
					return nil, err
				}

			*/
			// 当没有设置风暴预警，则不进行通知
			if agg.Storm == 0 {
				zlog.Debug(fmt.Sprintf("聚合匹配存在窗口,并且无风暴预警，告警标题：%s,匹配未关闭告警编号：%d", alert.AlertTitle, event.Id))
				ae.OutWindow = false
				ae.Notify = false
			} else {
				alertCount, err := redis.Rdb.Incr(context.TODO(), windowKey).Result()
				zlog.Debug(fmt.Sprintf("聚合匹配存在窗口,并且存在风暴预警，告警风暴阈值：%d， 当前聚合数量：%d，告警标题：%s,匹配未关闭告警编号：%d", agg.Storm, alertCount, alert.AlertTitle, event.Id))
				if err != nil {
					return nil, err
				}
				// 当聚合预警到达阈值时， 重新通知
				if alertCount%agg.Storm == 0 {
					ae.OutWindow = false
					ae.Notify = true
				} else {
					ae.OutWindow = false
					ae.Notify = false
				}
			}

			return &ae, nil

		} else {
			//当不存在窗口，则设置窗口
			zlog.Debug(fmt.Sprintf("聚合匹配不窗口,设置窗口，告警标题：%s,匹配未关闭告警编号：%d", alert.AlertTitle, event.Id))
			_, err := redis.Rdb.Set(context.TODO(), windowKey, 1, time.Minute*time.Duration(agg.Windows)).Result()
			if err != nil {
				return nil, err
			}
			ae.OutWindow = true
			ae.Notify = true
			return &ae, nil
		}

	}
	return nil, nil
}
