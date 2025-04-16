package alert

import (
	"context"
	"darkhawk/app/biz/model"
	"darkhawk/app/biz/service"
	"darkhawk/kitex_gen/agent"
	"errors"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

type AgentServiceImpl struct{}

func (s *AgentServiceImpl) Ping(ctx context.Context, req *agent.PingRequest) (res *emptypb.Empty, err error) {
	ruleSource := new(model.AlertRuleSource)
	err = model.GetOneWithScope(ctx, ruleSource.TableName(), &ruleSource, model.WhereWithScope("sign", req.GetSourceSign()))
	if err != nil {
		return nil, err
	}
	if ruleSource == nil {
		return nil, errors.New("数据源不存在")
	}
	t := model.LocalTime(time.Now())

	st := true
	agentModel := new(model.Agents)
	agentModel.ClientIp = req.GetClient()
	agentModel.RuleSourceId = ruleSource.Id
	agentModel.HeartbeatTime = &t
	agentModel.Status = &st
	err = agentModel.HeartBeat(ctx)
	return nil, err
}

func (s *AgentServiceImpl) UpdateConfig(ctx context.Context, req *agent.PingRequest) (res *emptypb.Empty, err error) {
	ruleSource := new(model.AlertRuleSource)
	err = model.GetOneWithScope(ctx, ruleSource.TableName(), &ruleSource, model.WhereWithScope("sign", req.GetSourceSign()))
	if err != nil {
		return nil, err
	}
	err = service.NewAlertRuleSourceService(ctx, nil).SyncPrometheusRule(ruleSource.Id, ruleSource.Address)
	return nil, err
}
