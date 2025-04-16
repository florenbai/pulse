package alert

import (
	"context"
	"darkhawk/app/biz/model"
	"darkhawk/pkg/utils"
	"darkhawk/pkg/zlog"
	"fmt"
	"strings"
	"time"
)

// CheckSilence 检测静默
func (alert *EventRequest) CheckSilence(ctx context.Context, workspaceId uint64) (bool, model.AlertSilence, error) {
	var data []model.AlertSilence
	alertSilenceModel := new(model.AlertSilence)
	err := model.GetAll(ctx, alertSilenceModel.TableName(), &data, model.WhereWithScope("workspace_id", workspaceId))
	if err != nil {
		return false, model.AlertSilence{}, err
	}
	for _, v := range data {
		if v.SilenceType == model.SilenceOnce {
			if rangeTimeIn(v.SilenceTime.RangeTime) && TagAndAlertNameIn(alert.Labels, v.Filters, alert.AlertTitle) {
				return true, v, nil
			}
		} else {
			zlog.Debug(fmt.Sprintf("告警周期静默匹配结果：%v ，周期匹配: %v", weekTimeIn(v.SilenceTime.Weeks, v.SilenceTime.Times) && TagIn(alert.Labels, v.Filters), weekTimeIn(v.SilenceTime.Weeks, v.SilenceTime.Times)))
			if weekTimeIn(v.SilenceTime.Weeks, v.SilenceTime.Times) && TagAndAlertNameIn(alert.Labels, v.Filters, alert.AlertTitle) {
				return true, v, nil
			}
		}
	}
	return false, model.AlertSilence{}, nil
}

func rangeTimeIn(rangeTime []string) bool {
	if len(rangeTime) < 2 {
		return false
	}
	loc, _ := time.LoadLocation("Asia/Shanghai")
	start, _ := time.ParseInLocation("2006-01-02 15:04:05", rangeTime[0], loc)
	end, _ := time.ParseInLocation("2006-01-02 15:04:05", rangeTime[1], loc)
	now := time.Now()
	if now.After(start) && now.Before(end) {
		return true
	}
	return false
}

func weekTimeIn(weeks []uint64, rangeTime []string) bool {
	if len(weeks) < 1 {
		return false
	}
	if len(rangeTime) < 2 {
		return false
	}
	now := time.Now()

	weekDay := now.Weekday()
	if utils.InOfT(uint64(weekDay), weeks) {
		return rangeHourIn(rangeTime)
	}
	return false
}

func rangeHourIn(rangeTime []string) bool {
	now := time.Now()
	rTime := []string{
		fmt.Sprintf("%s %s", now.Format("2006-01-02"), rangeTime[0]),
		fmt.Sprintf("%s %s", now.Format("2006-01-02"), rangeTime[1]),
	}
	return rangeTimeIn(rTime)
}

func TagIn(labels map[string]string, filter []model.TagFilter) bool {
	if len(filter) < 1 {
		return false
	}
	pass := false
	if HasTag(filter, labels) {
		zlog.Debug(fmt.Sprintf("告警标签匹配成功，匹配条件：%v", filter))
		pass = true
	}
	return pass
}

func TagAndAlertNameIn(labels map[string]string, filter []model.TagFilter, alertName string) bool {
	if len(filter) < 1 {
		return false
	}
	pass := false
	if HasTag(filter, labels) || HasName(filter, alertName) {
		zlog.Debug(fmt.Sprintf("告警标签匹配成功，匹配条件：%v", filter))
		pass = true
	}
	return pass
}

func HasName(filter []model.TagFilter, alertName string) bool {
	for _, f := range filter {
		if f.Tag == "alertname" {
			if f.Operation == "IN" {
				for _, item := range f.Value {
					if strings.Contains(alertName, item) {
						return true
					}
				}
			} else if f.Operation != "IN" {
				for _, item := range f.Value {
					if !strings.Contains(alertName, item) {
						return true
					}
				}
			}
		}
	}
	return false
}
