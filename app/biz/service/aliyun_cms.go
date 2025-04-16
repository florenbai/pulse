package service

import (
	"context"
	"darkhawk/app/biz/model"
	alertpack "darkhawk/app/biz/pack/alert"
	"encoding/json"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
	"time"
)

type AliyunCMSHookRequest struct {
	AlertName        string `json:"alertName" form:"alertName"`
	AlertState       string `json:"alertState" form:"alertState"`
	CurValue         string `json:"curValue" form:"curValue"`
	Dimensions       string `json:"dimensions" form:"dimensions"`
	Expression       string `json:"expression" form:"expression"`
	GroupId          string `json:"groupId" form:"groupId"`
	InstanceName     string `json:"instanceName" form:"instanceName"`
	LastTime         string `json:"lastTime" form:"lastTime"`
	MetricName       string `json:"metricName" form:"metricName"`
	MetricProject    string `json:"metricProject" form:"metricProject"`
	Namespace        string `json:"namespace" form:"namespace"`
	PreTriggerLevel  string `json:"preTriggerLevel" form:"preTriggerLevel"`
	ProductGroupName string `json:"productGroupName" form:"productGroupName"`
	RawMetricName    string `json:"rawMetricName" form:"rawMetricName"`
	RegionId         string `json:"regionId" form:"regionId"`
	RegionName       string `json:"regionName" form:"regionName"`
	RuleId           string `json:"ruleId" form:"ruleId"`
	Timestamp        string `json:"timestamp" form:"timestamp"`
	TransId          string `json:"transId" form:"transId"`
	TriggerLevel     string `json:"triggerLevel" form:"triggerLevel"`
	Unit             string `json:"unit" form:"unit"`
	UserId           string `json:"userId" form:"userId"`
}

type AliyunCMSService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewAliyunCMSService(ctx context.Context, c *app.RequestContext) *AliyunCMSService {
	return &AliyunCMSService{
		ctx: ctx,
		c:   c,
	}
}

func (service *AliyunCMSService) GetEvent(alert AliyunCMSHookRequest, integration model.Integration) ([]alertpack.EventRequest, error) {
	var status string
	if alert.AlertState == "ALERT" || alert.AlertState == "INSUFFICIENT_DATA" {
		status = "firing"
	} else {
		status = "resolved"
	}
	levelId, levelName, err := service.GetLevel(alert, integration)
	if err != nil {
		return nil, err
	}
	labels := service.getLabels(alert)
	unixTimestampSec, _ := strconv.Atoi(alert.Timestamp)
	timestamp := time.Unix(int64(unixTimestampSec), 0)
	t := timestamp.Format("2006-01-02 15:04:05")
	ip := service.c.ClientIP()
	alertEvent := alertpack.EventRequest{
		AlertTitle:  alert.AlertName,
		Description: alert.Dimensions,
		Level:       levelName,
		Status:      status,
		AlertTime:   &t,
		Labels:      labels,
		Annotations: service.getAnnotations(alert),
		FingerPrint: alert.TransId,
		SourceIp:    ip,
		LevelId:     levelId,
	}
	return []alertpack.EventRequest{alertEvent}, nil
}

func (service *AliyunCMSService) GetLevel(alert AliyunCMSHookRequest, integration model.Integration) (uint64, string, error) {
	//告警等级转换
	alertJsonByte, err := json.Marshal(alert)
	if err != nil {
		return 0, "", err
	}
	return alertpack.GetAlertLevel(alertJsonByte, integration.LevelField, integration.Id)
}

func (service *AliyunCMSService) getLabels(alert AliyunCMSHookRequest) map[string]string {
	return map[string]string{
		"curValue":         alert.CurValue,
		"instanceName":     alert.InstanceName,
		"metricName":       alert.MetricName,
		"metricProject":    alert.MetricProject,
		"namespace":        alert.Namespace,
		"preTriggerLevel":  alert.PreTriggerLevel,
		"productGroupName": alert.ProductGroupName,
		"rawMetricName":    alert.RawMetricName,
		"regionId":         alert.RegionId,
		"regionName":       alert.RegionName,
		"ruleId":           alert.RuleId,
		"triggerLevel":     alert.TriggerLevel,
		"unit":             alert.Unit,
		"userId":           alert.UserId,
	}
}

func (service *AliyunCMSService) getAnnotations(alert AliyunCMSHookRequest) map[string]string {
	return map[string]string{
		"expression": alert.Expression,
		"dimensions": alert.Dimensions,
		"lastTime":   alert.LastTime,
	}
}
