package service

import (
	"context"
	"darkhawk/app/biz/model"
	pb "darkhawk/hertz_gen/alert_rule"
	"darkhawk/pkg/http"
	"darkhawk/pkg/response"
	"darkhawk/rpc/agent"
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/tidwall/gjson"
)

type AlertRuleService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewAlertRuleService(ctx context.Context, c *app.RequestContext) *AlertRuleService {
	return &AlertRuleService{
		ctx: ctx,
		c:   c,
	}
}

// 创建告警规则
func (service *AlertRuleService) CreateRule(req pb.RuleRequest) error {
	ruleModel := new(model.AlertRule)
	groupModel := new(model.AlertRuleGroup)
	err := model.GetOneById(service.ctx, groupModel.TableName(), req.GetGid(), &groupModel)
	if err != nil {
		return err
	}
	ok, err := model.ExistName(service.ctx, ruleModel.TableName(), "name", req.GetName(), nil)
	if err != nil {
		return err
	}
	if ok {
		return response.NameAlreadyExistErr
	}
	var labels = make(map[string]string)
	var annotations = make(map[string]string)
	for _, v := range req.GetLabels() {
		labels[v.GetTag()] = v.GetValue()
	}

	for _, v := range req.GetAnnotation() {
		annotations[v.GetTag()] = v.GetValue()
	}
	ruleModel.Name = req.GetName()
	ruleModel.Query = req.GetQuery()
	ruleModel.Gid = req.GetGid()
	ruleModel.SourceId = groupModel.RuleSource
	ruleModel.Duration = req.GetDuration()
	ruleModel.Annotations = annotations
	ruleModel.Labels = labels
	ruleModel.Health = "local"
	err = agent.CreateAlertRule(service.ctx, groupModel.File, groupModel.GroupName, *ruleModel)
	if err != nil {
		return err
	}
	return model.Create(context.TODO(), ruleModel.TableName(), &ruleModel)
}

// 编辑告警规则
func (service *AlertRuleService) UpdateRule(id uint64, req pb.RuleRequest) error {
	ruleModel := new(model.AlertRule)
	err := model.GetOneById(service.ctx, ruleModel.TableName(), id, &ruleModel)
	if err != nil {
		return err
	}
	ok, err := model.ExistName(service.ctx, ruleModel.TableName(), "name", req.GetName(), &id)
	if err != nil {
		return err
	}
	if ok {
		return response.NameAlreadyExistErr
	}
	groupModel := new(model.AlertRuleGroup)
	err = model.GetOneById(service.ctx, groupModel.TableName(), ruleModel.Gid, &groupModel)
	if err != nil {
		return err
	}
	var labels = make(map[string]string)
	var annotations = make(map[string]string)
	for _, v := range req.GetLabels() {
		labels[v.GetTag()] = v.GetValue()
	}
	oldName := ruleModel.Name
	for _, v := range req.GetAnnotation() {
		annotations[v.GetTag()] = v.GetValue()
	}
	ruleModel.Name = req.GetName()
	ruleModel.Query = req.GetQuery()
	ruleModel.Gid = req.GetGid()
	ruleModel.Duration = req.GetDuration()
	ruleModel.Annotations = annotations
	ruleModel.Labels = labels
	ruleModel.Health = "local"

	err = agent.UpdateAlertRule(service.ctx, groupModel.File, groupModel.GroupName, oldName, *ruleModel)
	if err != nil {
		return err
	}
	return model.EditOneById(context.TODO(), ruleModel.TableName(), id, &ruleModel)
}

// 删除告警规则
func (service *AlertRuleService) DeleteRule(id uint64) error {
	ruleModel := new(model.AlertRule)
	err := model.GetOneById(service.ctx, ruleModel.TableName(), id, &ruleModel)
	if err != nil {
		return err
	}
	groupModel := new(model.AlertRuleGroup)
	err = model.GetOneById(service.ctx, groupModel.TableName(), ruleModel.Gid, &groupModel)
	if err != nil {
		return err
	}
	err = agent.DeleteRemoteGroupAlert(service.ctx, groupModel.File, groupModel.GroupName, *ruleModel)
	if err != nil {
		return err
	}
	return model.DeleteOneById(service.ctx, ruleModel.TableName(), id, &ruleModel)
}

// 获取所有告警规则组
func (service *AlertRuleService) GetRuleGroup() ([]model.AlertRuleGroup, error) {
	var data []model.AlertRuleGroup
	ruleGroupModel := new(model.AlertRuleGroup)
	err := model.GetAll(service.ctx, ruleGroupModel.TableName(), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (service *AlertRuleService) GetRuleListByGroup(id uint64) ([]model.AlertRuleList, error) {
	ruleModel := new(model.AlertRule)
	data, err := ruleModel.GetRuleListByGroup(service.ctx, id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (service *AlertRuleService) CreateRuleGroup(req pb.RuleGroupRequest) error {
	ruleGroupModel := new(model.AlertRuleGroup)
	ok, err := model.ExistName(service.ctx, ruleGroupModel.TableName(), "group_name", req.GetGroupName(), nil)
	if err != nil {
		return err
	}
	if ok {
		return response.NameAlreadyExistErr
	}
	err = agent.CreateRemoteGroupFile(service.ctx, req.GetFile())
	if err != nil {
		return err
	}
	ruleGroupModel.GroupName = req.GetGroupName()
	ruleGroupModel.RuleSource = req.GetRuleSource()
	ruleGroupModel.File = req.GetFile()
	return model.Create(service.ctx, ruleGroupModel.TableName(), &ruleGroupModel)
}

func (service *AlertRuleService) DeleteRuleGroup(id uint64) error {
	ruleGroupModel := new(model.AlertRuleGroup)
	err := model.GetOneById(service.ctx, ruleGroupModel.TableName(), id, &ruleGroupModel)
	if err != nil {
		return err
	}
	err = agent.DeleteRemoteGroupFile(service.ctx, ruleGroupModel.File)
	if err != nil {
		return err
	}
	return ruleGroupModel.DeleteGroupAndChildrenRule(service.ctx, id)
}

func (service *AlertRuleService) GetLabelsByRuleGroup(id uint64) ([]string, error) {
	var data []string
	ruleSourceModel := new(model.AlertRuleSource)
	source, err := ruleSourceModel.GetRuleSourceByGroupId(service.ctx, id)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/api/v1/label/__name__/values", source.Address)
	resp, err := http.Get(url, nil, nil)
	if err != nil {
		return nil, err
	}
	rjson := gjson.ParseBytes(resp.Body())
	if resp.StatusCode() == 200 && rjson.Get("status").String() == "success" {
		rjson.Get("data").ForEach(func(key, value gjson.Result) bool {
			data = append(data, value.String())
			return true
		})
	} else {
		return nil, errors.New("label数据获取异常")
	}
	return data, nil
}
