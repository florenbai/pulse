package service

import (
	"context"
	"darkhawk/app/biz/model"
	alertpack "darkhawk/app/biz/pack/alert"
	"encoding/json"
	"github.com/cloudwego/hertz/pkg/app"
	"strings"
	"time"
)

type PrometheusHookRequest struct {
	Version           string            `json:"version" form:"version"`
	Receiver          string            `json:"receiver"`
	Status            string            `json:"status"`
	Alerts            []PrometheusAlert `json:"alerts"`
	GroupLabels       map[string]string `json:"groupLabels"`
	CommonLabels      map[string]string `json:"commonLabels"`
	CommonAnnotations map[string]string `json:"commonAnnotations"`
	ExternalURL       string            `json:"externalURL"`
}

type PrometheusAlert struct {
	Status       string            `json:"status"`
	Labels       map[string]string `json:"labels"`
	Annotations  map[string]string `json:"annotations"`
	StartsAt     time.Time         `json:"startsAt"`
	EndsAt       time.Time         `json:"endsAt"`
	GeneratorURL string            `json:"generatorURL"`
	Fingerprint  string            `json:"fingerprint"`
}

type PrometheusAlertService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewPrometheusAlertService(ctx context.Context, c *app.RequestContext) *PrometheusAlertService {
	return &PrometheusAlertService{
		ctx: ctx,
		c:   c,
	}
}

func (service *PrometheusAlertService) GetEvent(alerts []PrometheusAlert, integration model.Integration) ([]alertpack.EventRequest, error) {
	var event []alertpack.EventRequest
	ip := service.c.ClientIP()
	for _, v := range alerts {
		levelId, levelName, err := service.GetLevel(v, integration)
		if err != nil {
			return event, err
		}

		t := v.StartsAt.Format("2006-01-02 15:04:05")
		alertEvent := alertpack.EventRequest{
			AlertTitle:  service.getAlertName(v.Labels),
			Description: service.getAnnotationsDescription(v.Annotations),
			Level:       levelName,
			Status:      v.Status,
			AlertTime:   &t,
			Labels:      service.getLabels(v.Labels, integration.LevelField),
			Annotations: v.Annotations,
			FingerPrint: v.Fingerprint,
			SourceIp:    ip,
			LevelId:     levelId,
		}
		event = append(event, alertEvent)
	}
	return event, nil
}

// 获取Prometheus描述
// 依次匹配 description、message, 如果没有匹配到,则取map第一个值
func (service *PrometheusAlertService) getAnnotationsDescription(annotations map[string]string) string {
	if description, ok := annotations["description"]; ok {
		return description
	}
	if message, ok := annotations["message"]; ok {
		return message
	}
	for _, v := range annotations {
		return v
	}
	return ""
}

// 获取Prometheus Label
// 依次匹配 description、message, 如果没有匹配到,则取map第一个值
func (service *PrometheusAlertService) getLabels(labels map[string]string, levelFiled string) map[string]string {
	var data = make(map[string]string)
	for k, v := range labels {
		if k != "alertname" || strings.Contains(levelFiled, k) {
			data[k] = v
		}
	}
	return data
}

// 获取Prometheus 告警标题
func (service *PrometheusAlertService) getAlertName(labels map[string]string) string {
	if name, ok := labels["alertname"]; ok {
		return name
	}
	return ""
}

func (service *PrometheusAlertService) GetLevel(alert PrometheusAlert, integration model.Integration) (uint64, string, error) {
	//告警等级转换
	alertJsonByte, err := json.Marshal(alert)
	if err != nil {
		return 0, "", err
	}
	return alertpack.GetAlertLevel(alertJsonByte, integration.LevelField, integration.Id)
}
