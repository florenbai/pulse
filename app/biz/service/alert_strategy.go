package service

import (
	"context"
	"darkhawk/app/biz/model"
	pb "darkhawk/hertz_gen/alert_strategy"
	"darkhawk/hertz_gen/base"
	"darkhawk/pkg/response"
	"darkhawk/pkg/utils"
	"encoding/json"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
)

type AlertStrategyService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewAlertStrategyService(ctx context.Context, c *app.RequestContext) *AlertStrategyService {
	return &AlertStrategyService{
		ctx: ctx,
		c:   c,
	}
}

// CreateAlertStrategy 新增告警策略
func (service *AlertStrategyService) CreateAlertStrategy(req pb.StrategyRequest) error {
	alertStrategyModel := new(model.AlertStrategy)
	ok, err := model.ExistName(service.ctx, alertStrategyModel.TableName(), "strategy_name", req.GetStrategyName(), nil)
	if err != nil {
		return err
	}
	if ok {
		return response.NameAlreadyExistErr
	}
	if err = utils.AnyToAny(req, &alertStrategyModel); err != nil {
		return err
	}
	session := sessions.Default(service.c)
	uid := session.Get("uid").(uint64)
	alertStrategyModel.Uid = uid
	return alertStrategyModel.Create(service.ctx)
}

// CreateSystemStrategy 新增系统告警策略
func (service *AlertStrategyService) CreateSystemStrategy(req pb.SystemStrategyRequest) error {
	systemStrategyModel := new(model.SystemStrategy)
	ok, err := model.ExistName(service.ctx, systemStrategyModel.TableName(), "strategy_name", req.GetStrategyName(), nil)
	if err != nil {
		return err
	}
	if ok {
		return response.NameAlreadyExistErr
	}
	if err = utils.AnyToAny(req, &systemStrategyModel); err != nil {
		return err
	}
	session := sessions.Default(service.c)
	uid := session.Get("uid").(uint64)
	systemStrategyModel.Uid = uid
	return model.Create(service.ctx, systemStrategyModel.TableName(), &systemStrategyModel)
}

// UpdateAlertStrategy 编辑告警策略
func (service *AlertStrategyService) UpdateAlertStrategy(id uint64, req pb.StrategyRequest) error {
	alertStrategyModel := new(model.AlertStrategy)
	err := model.GetOneById(service.ctx, alertStrategyModel.TableName(), id, &alertStrategyModel)
	if err != nil {
		return err
	}
	ok, err := model.ExistName(service.ctx, alertStrategyModel.TableName(), "strategy_name", req.GetStrategyName(), &id)
	if err != nil {
		return err
	}
	if ok {
		return response.NameAlreadyExistErr
	}
	status := alertStrategyModel.Status
	if err = utils.AnyToAny(req, &alertStrategyModel); err != nil {
		return err
	}
	session := sessions.Default(service.c)
	uid := session.Get("uid").(uint64)
	alertStrategyModel.Uid = uid
	alertStrategyModel.Status = status
	return model.EditOneById(service.ctx, alertStrategyModel.TableName(), id, &alertStrategyModel)
}

// UpdateSystemStrategy 编辑系统告警策略
func (service *AlertStrategyService) UpdateSystemStrategy(id uint64, req pb.SystemStrategyRequest) error {
	systemStrategyModel := new(model.SystemStrategy)
	err := model.GetOneById(service.ctx, systemStrategyModel.TableName(), id, &systemStrategyModel)
	if err != nil {
		return err
	}
	ok, err := model.ExistName(service.ctx, systemStrategyModel.TableName(), "strategy_name", req.GetStrategyName(), &id)
	if err != nil {
		return err
	}
	if ok {
		return response.NameAlreadyExistErr
	}
	if err = utils.AnyToAny(req, &systemStrategyModel); err != nil {
		return err
	}
	session := sessions.Default(service.c)
	uid := session.Get("uid").(uint64)
	systemStrategyModel.Uid = uid
	return model.EditOneById(service.ctx, systemStrategyModel.TableName(), id, &systemStrategyModel)
}

// ChangeAlertStrategyStatus 设置告警状态
func (service *AlertStrategyService) ChangeAlertStrategyStatus(id uint64) error {
	alertStrategyModel := new(model.AlertStrategy)
	err := model.GetOneById(service.ctx, alertStrategyModel.TableName(), id, &alertStrategyModel)
	if err != nil {
		return err
	}
	if alertStrategyModel.Status == model.StrategyEnabled {
		alertStrategyModel.Status = model.StrategyDisabled
	} else {
		alertStrategyModel.Status = model.StrategyEnabled
	}
	return model.EditOneById(service.ctx, alertStrategyModel.TableName(), id, &alertStrategyModel)
}

// ChangeSort 改变告警排序
func (service *AlertStrategyService) ChangeSort(req pb.ChangeStrategySortRequest) error {
	m := new(model.AlertStrategy)
	for _, v := range req.NewWeight {
		err := model.EditOneById(service.ctx, m.TableName(), v.Id, map[string]interface{}{"weight": v.Weight})
		if err != nil {
			return err
		}
	}
	return nil
}

// DeleteStrategy 删除告警
func (service *AlertStrategyService) DeleteStrategy(id uint64) error {
	alertStrategyModel := new(model.AlertStrategy)
	err := model.GetOneById(service.ctx, alertStrategyModel.TableName(), id, &alertStrategyModel)
	if err != nil {
		return err
	}
	alertStrategyModel.Status = model.StrategyDeleted
	return model.EditOneById(service.ctx, alertStrategyModel.TableName(), id, &alertStrategyModel)
}

// DeleteSystemStrategy 删除系统告警
func (service *AlertStrategyService) DeleteSystemStrategy(id uint64) error {
	systemStrategyModel := new(model.SystemStrategy)
	err := model.GetOneById(service.ctx, systemStrategyModel.TableName(), id, &systemStrategyModel)
	if err != nil {
		return err
	}

	return model.DeleteOneById(service.ctx, systemStrategyModel.TableName(), id, &systemStrategyModel)
}

// GetSystemStrategyList 获取系统告警列表
func (service *AlertStrategyService) GetSystemStrategyList(req pb.SystemStrategyQuery) (*base.BaseListResp, error) {
	search := map[string]string{
		"strategy_name": req.GetStrategyName(),
	}
	associate := map[string]string{
		"strategy_name": "strategy_name LIKE ?",
	}
	scopes := model.ParamWithScope(search, associate, nil, false)
	strategyModel := new(model.SystemStrategy)
	data, i, err := strategyModel.GetSystemStrategyList(service.ctx, req.GetPage(), req.GetPageSize(), scopes)
	if err != nil {
		return nil, err
	}
	listValue := new(structpb.ListValue)
	b, _ := json.Marshal(data)
	err = protojson.Unmarshal(b, listValue)
	if err != nil {
		return nil, err
	}
	return &base.BaseListResp{
		Total: uint64(i),
		List:  listValue,
	}, nil
}

func (service *AlertStrategyService) GetSystemStrategyAll() ([]model.SystemStrategy, error) {
	var data []model.SystemStrategy
	strategyModel := new(model.SystemStrategy)
	err := model.GetAll(service.ctx, strategyModel.TableName(), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (service *AlertStrategyService) GetSystemStrategyInfoByWorkspace(id uint64) (*model.SystemStrategy, error) {
	workspaceModel := new(model.Workspace)
	err := model.GetOneById(service.ctx, workspaceModel.TableName(), id, &workspaceModel)
	if err != nil {
		return nil, err
	}

	systemStrategyModel := new(model.SystemStrategy)
	err = model.GetOneById(service.ctx, systemStrategyModel.TableName(), workspaceModel.SystemStrategy, &systemStrategyModel)
	if err != nil {
		return nil, err
	}
	return systemStrategyModel, nil
}
