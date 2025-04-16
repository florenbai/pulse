package service

import (
	"context"
	"darkhawk/app/biz/model"
	alertpack "darkhawk/app/biz/pack/alert"
	"encoding/json"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/google/uuid"
	"time"
)

type AliyunEventRequest struct {
	Alert AlertMeta `json:"alert"`
}

type AlertMeta struct {
	Meta            EventMeta         `json:"meta"`
	EventContentMap map[string]string `json:"eventContentMap"`
}
type EventMeta struct {
	SysEventMeta SysEventMeta `json:"sysEventMeta"`
}

type SysEventMeta struct {
	Product      string            `json:"product"`
	ResourceId   string            `json:"resourceId"`
	Level        string            `json:"level"`
	InstanceName string            `json:"instanceName"`
	RegionId     string            `json:"regionId"`
	Name         string            `json:"name"`
	Content      map[string]string `json:"content"`
	Status       string            `json:"status"`
	EventNameZh  string            `json:"eventNameZh"`
}

type AliyunEventService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewAliyunEventService(ctx context.Context, c *app.RequestContext) *AliyunEventService {
	return &AliyunEventService{
		ctx: ctx,
		c:   c,
	}
}

func (service *AliyunEventService) GetEvent(alert AliyunEventRequest, integration model.Integration) ([]alertpack.EventRequest, error) {

	levelId, levelName, err := service.GetLevel(alert, integration)
	if err != nil {
		return nil, err
	}
	t := time.Now().Format("2006-01-02 15:04:05")
	ip := service.c.ClientIP()
	alert.Alert.Meta.SysEventMeta.Content = alert.Alert.EventContentMap
	b, _ := json.Marshal(alert.Alert.Meta.SysEventMeta)
	alertEvent := alertpack.EventRequest{
		AlertTitle:  alert.Alert.Meta.SysEventMeta.EventNameZh,
		Description: string(b),
		Level:       levelName,
		Status:      "firing",
		Labels: map[string]string{
			"product":      alert.Alert.Meta.SysEventMeta.Product,
			"resourceId":   alert.Alert.Meta.SysEventMeta.ResourceId,
			"instanceName": alert.Alert.Meta.SysEventMeta.InstanceName,
			"regionId":     alert.Alert.Meta.SysEventMeta.RegionId,
			"eventNameZh":  alert.Alert.Meta.SysEventMeta.EventNameZh,
			"name":         alert.Alert.Meta.SysEventMeta.Name,
			"status":       alert.Alert.Meta.SysEventMeta.Status,
			"level":        alert.Alert.Meta.SysEventMeta.Level,
		},
		AlertTime:   &t,
		FingerPrint: uuid.New().String(),
		SourceIp:    ip,
		LevelId:     levelId,
	}
	return []alertpack.EventRequest{alertEvent}, nil
}

func (service *AliyunEventService) GetLevel(alert AliyunEventRequest, integration model.Integration) (uint64, string, error) {
	//告警等级转换
	alertJsonByte, err := json.Marshal(alert)
	if err != nil {
		return 0, "", err
	}
	return alertpack.GetAlertLevel(alertJsonByte, integration.LevelField, integration.Id)
}
