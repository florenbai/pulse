package service

import (
	"context"
	"darkhawk/app/biz/model"
	rulepack "darkhawk/app/biz/pack/alert_rule"
	pb "darkhawk/hertz_gen/alert_rule"
	basepb "darkhawk/hertz_gen/base"
	"darkhawk/pkg/response"
	"darkhawk/rpc/agent"
	"encoding/json"
	"github.com/cloudwego/hertz/pkg/app"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
	"time"
)

type AlertRuleSourceService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewAlertRuleSourceService(ctx context.Context, c *app.RequestContext) *AlertRuleSourceService {
	return &AlertRuleSourceService{
		ctx: ctx,
		c:   c,
	}
}

// 创建规则数据源
func (service *AlertRuleSourceService) CreateSource(req pb.SourceRequest) error {
	alertRuleSourceModel := new(model.AlertRuleSource)
	ok, err := model.ExistName(service.ctx, alertRuleSourceModel.TableName(), "source_name", req.GetSourceName(), nil)
	if err != nil {
		return err
	}
	if ok {
		return response.NameAlreadyExistErr
	}

	alertRuleSourceModel.SourceName = req.GetSourceName()
	alertRuleSourceModel.SourceType = req.GetSourceType()
	alertRuleSourceModel.Mark = req.GetMark()
	alertRuleSourceModel.AutoSync = true
	alertRuleSourceModel.Address = req.GetAddress()
	alertRuleSourceModel.Sign = req.GetSign()

	err = model.Create(service.ctx, alertRuleSourceModel.TableName(), &alertRuleSourceModel)
	if err != nil {
		return err
	}
	return service.SyncPrometheusRule(alertRuleSourceModel.Id, req.GetAddress())
}

// 编辑规则数据源
func (service *AlertRuleSourceService) UpdateSource(id uint64, req pb.SourceRequest) error {
	alertRuleSourceModel := new(model.AlertRuleSource)
	ok, err := model.ExistName(service.ctx, alertRuleSourceModel.TableName(), "source_name", req.GetSourceName(), &id)
	if err != nil {
		return err
	}
	if ok {
		return response.NameAlreadyExistErr
	}

	alertRuleSourceModel.SourceName = req.GetSourceName()
	alertRuleSourceModel.SourceType = req.GetSourceType()
	alertRuleSourceModel.Mark = req.GetMark()
	alertRuleSourceModel.Address = req.GetAddress()

	err = model.EditOneById(service.ctx, alertRuleSourceModel.TableName(), id, &alertRuleSourceModel)
	if err != nil {
		return err
	}
	return service.SyncPrometheusRule(id, req.GetAddress())
}

// 获取数据源列表
func (service *AlertRuleSourceService) GetSourceList(req basepb.PageInfoReq) (*basepb.BaseListResp, error) {
	var data []model.AlertRuleSource
	alertRuleSourceModel := new(model.AlertRuleSource)
	i, err := model.GetPageList(service.ctx, alertRuleSourceModel.TableName(), req.GetPage(), req.GetPageSize(), &data)
	if err != nil {
		return nil, err
	}
	listValue := new(structpb.ListValue)
	b, _ := json.Marshal(data)
	err = protojson.Unmarshal(b, listValue)
	if err != nil {
		return nil, err
	}
	return &basepb.BaseListResp{
		Total: uint64(i),
		List:  listValue,
	}, nil
}

func (service *AlertRuleSourceService) GetAllSource() ([]model.AlertRuleSource, error) {
	var data []model.AlertRuleSource
	alertRuleSourceModel := new(model.AlertRuleSource)
	err := model.GetAll(service.ctx, alertRuleSourceModel.TableName(), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (service *AlertRuleSourceService) DeleteAlertRuleSource(id uint64) error {
	alertRuleSourceModel := new(model.AlertRuleSource)
	err := model.GetOneById(service.ctx, alertRuleSourceModel.TableName(), id, &alertRuleSourceModel)
	if err != nil {
		return err
	}
	return alertRuleSourceModel.DeleteSource(service.ctx, id)
}

// 同步告警数据源
func (service *AlertRuleSourceService) SyncAlertRuleSource() error {
	var data []model.AlertRuleSource
	alertRuleSourceModel := new(model.AlertRuleSource)
	err := model.GetAll(service.ctx, alertRuleSourceModel.TableName(), &data, model.WhereWithScope("auto_sync", 1))
	if err != nil {
		return err
	}
	for _, v := range data {
		if v.SourceType == "Prometheus" {
			service.SyncPrometheusRule(v.Id, v.Address)
		}
	}
	return nil
}

// 同步Prometheus告警规则
func (service *AlertRuleSourceService) SyncPrometheusRule(id uint64, addr string) error {
	resp, err := rulepack.GetRemotePrometheusRule(addr)
	if err != nil {
		return err
	}
	groupModel := new(model.AlertRuleGroup)
	groupMap, err := groupModel.GetGroupNameMap(service.ctx, id)
	if err != nil {
		return err
	}
	var rules []model.AlertRule
	for _, v := range resp.Groups {
		var gid uint64
		if _, ok := groupMap[v.Name]; !ok {
			ng := model.AlertRuleGroup{
				File:       v.File,
				RuleSource: id,
				GroupName:  v.Name,
			}
			err := model.Create(service.ctx, groupModel.TableName(), &ng)
			if err != nil {
				return err
			}
			gid = ng.Id
			groupMap[v.Name] = ng.Id
		} else {
			gid = groupMap[v.Name]
		}

		for _, r := range v.Rules {
			rule := model.AlertRule{
				Name:        r.Name,
				Health:      r.Health,
				Annotations: r.Annotations,
				Labels:      r.Labels,
				Duration:    r.Duration,
				Query:       r.Query,
				Gid:         gid,
				SourceId:    id,
			}
			rules = append(rules, rule)
		}
	}
	ruleModel := new(model.AlertRule)
	err = ruleModel.ClausesRule(service.ctx, rules)
	if err != nil {
		return err
	}
	return nil
}

func (service *AlertRuleSourceService) GetSourceAgent(id uint64) ([]model.Agents, error) {
	var data []model.Agents
	agentModel := new(model.Agents)
	atime := time.Now().Add(-3 * time.Minute).Format("2006-01-02 15:04:05")
	err := model.GetAll(service.ctx, agentModel.TableName(), &data, model.WhereWithScope("heartbeat_time > ?", atime), model.WhereWithScope("rule_source_id", id))
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (service *AlertRuleSourceService) ReloadSource(id uint64) error {
	ruleSourceModel := new(model.AlertRuleSource)
	err := model.GetOneById(service.ctx, ruleSourceModel.TableName(), id, &ruleSourceModel)
	if err != nil {
		return err
	}
	return agent.PrometheusReload(service.ctx)
}
