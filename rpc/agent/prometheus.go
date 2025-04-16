package agent

import (
	"context"
	"darkhawk/app/biz/model"
	"darkhawk/conf"
	"darkhawk/kitex_gen/prometheus"
	"darkhawk/kitex_gen/prometheus/prometheusagentservice"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/kitex-contrib/registry-nacos/resolver"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"strconv"
	"time"
)

var prometheusAgentClient prometheusagentservice.Client

const PrometheusAgentServiceName = "owl-prometheus-agent"

func InitPrometheusAgentRpc() {
	/*
		r, err := consul.NewConsulResolver(conf.GetConf().Consul.Address)
		if err != nil {
			klog.Error(err)
			return
		}

	*/
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(conf.GetConf().Nacos.Addr, conf.GetConf().Nacos.Port),
	}

	cc := constant.ClientConfig{
		NamespaceId:         conf.GetConf().Nacos.Namespace,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              conf.GetConf().Nacos.LogDir,
		CacheDir:            conf.GetConf().Nacos.CacheDir,
		LogLevel:            conf.GetConf().Nacos.LogLevel,
		Username:            conf.GetConf().Nacos.Username,
		Password:            conf.GetConf().Nacos.Password,
	}

	cli, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		klog.Fatal(err)
	}
	r := resolver.NewNacosResolver(cli)
	prometheusAgentClient = prometheusagentservice.MustNewClient(PrometheusAgentServiceName, client.WithResolver(r), client.WithRPCTimeout(time.Second*10))
}

func CreateRemoteGroupFile(ctx context.Context, filename string) error {
	req := prometheus.AlertGroupRequest{FileName: filename}
	_, err := prometheusAgentClient.CreateAlertGroup(ctx, &req)
	return err
}

func DeleteRemoteGroupFile(ctx context.Context, filename string) error {
	req := prometheus.AlertGroupRequest{FileName: filename}
	_, err := prometheusAgentClient.DeleteAlertGroup(ctx, &req)
	return err
}

func CreateAlertRule(ctx context.Context, filename string, groupName string, rule model.AlertRule) error {
	var req prometheus.AlertRules
	var ruleFor string
	f := float64(rule.Duration) / 60
	_, err := strconv.Atoi(fmt.Sprintf("%v", f))
	if err != nil {
		ruleFor = fmt.Sprintf("%ds", rule.Duration)
	} else {
		ruleFor = fmt.Sprintf("%vm", f)
	}
	alertRule := prometheus.AlertRule{
		Alert:       rule.Name,
		Expr:        rule.Query,
		For:         ruleFor,
		Labels:      rule.Labels,
		Annotations: rule.Annotations,
		Filename:    filename,
	}
	req.GroupName = groupName
	req.Filename = filename
	req.Rules = append(req.Rules, &alertRule)
	_, err = prometheusAgentClient.CreateAlertRules(ctx, &req)
	return err
}

func UpdateAlertRule(ctx context.Context, filename string, groupName string, ordName string, rule model.AlertRule) error {
	var req prometheus.UpdateAlertRuleRequest
	var ruleFor string
	f := float64(rule.Duration) / 60
	_, err := strconv.Atoi(fmt.Sprintf("%v", f))
	if err != nil {
		ruleFor = fmt.Sprintf("%ds", rule.Duration)
	} else {
		ruleFor = fmt.Sprintf("%vm", f)
	}
	req.AlertRule = &prometheus.AlertRule{
		Alert:       rule.Name,
		Expr:        rule.Query,
		For:         ruleFor,
		Labels:      rule.Labels,
		Annotations: rule.Annotations,
		Filename:    filename,
	}
	req.GroupName = groupName
	req.Filename = filename
	req.UpdateAlertName = ordName

	_, err = prometheusAgentClient.UpdateAlertRule(ctx, &req)
	return err
}

func DeleteRemoteGroupAlert(ctx context.Context, filename string, groupName string, rule model.AlertRule) error {
	var req prometheus.AlertRules
	alertRule := prometheus.AlertRule{
		Alert:       rule.Name,
		Expr:        rule.Query,
		For:         fmt.Sprintf("%d", rule.Duration),
		Labels:      rule.Labels,
		Annotations: rule.Annotations,
		Filename:    filename,
	}
	req.GroupName = groupName
	req.Filename = filename
	req.Rules = append(req.Rules, &alertRule)
	_, err := prometheusAgentClient.DeleteAlertRule(ctx, &req)
	return err
}

func PrometheusReload(ctx context.Context) error {
	_, err := prometheusAgentClient.Reload(ctx, nil)
	return err
}
