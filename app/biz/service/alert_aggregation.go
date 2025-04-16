package service

import (
	"context"
	"darkhawk/app/biz/model"
	pb "darkhawk/hertz_gen/alert_aggregation"
	"darkhawk/pkg/response"
	"darkhawk/pkg/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

type AlertAggregationService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewAlertAggregationService(ctx context.Context, c *app.RequestContext) *AlertAggregationService {
	return &AlertAggregationService{
		ctx: ctx,
		c:   c,
	}
}

// CreateAlertAggregation 创建告警聚合
func (service *AlertAggregationService) CreateAlertAggregation(req pb.AggregationRequest) error {
	aggregationModel := new(model.AlertAggregation)
	ok, err := model.ExistName(service.ctx, aggregationModel.TableName(), "aggregation_name", req.GetAggregationName(), nil)
	if err != nil {
		return err
	}
	if ok {
		return response.NameAlreadyExistErr
	}
	if err = utils.AnyToAny(req, &aggregationModel); err != nil {
		return err
	}
	session := sessions.Default(service.c)
	uid := session.Get("uid").(uint64)
	aggregationModel.Uid = uid
	aggregationModel.Status = model.StrategyEnabled
	return model.Create(service.ctx, aggregationModel.TableName(), &aggregationModel)
}

// UpdateAlertAggregation 编辑告警聚合
func (service *AlertAggregationService) UpdateAlertAggregation(id uint64, req pb.AggregationRequest) error {
	aggregationModel := new(model.AlertAggregation)
	err := model.GetOneById(service.ctx, aggregationModel.TableName(), id, &aggregationModel)
	if err != nil {
		return err
	}
	ok, err := model.ExistName(service.ctx, aggregationModel.TableName(), "aggregation_name", req.GetAggregationName(), &id)
	if err != nil {
		return err
	}
	if ok {
		return response.NameAlreadyExistErr
	}
	status := aggregationModel.Status
	if err = utils.AnyToAny(req, &aggregationModel); err != nil {
		return err
	}
	session := sessions.Default(service.c)
	uid := session.Get("uid").(uint64)
	aggregationModel.Uid = uid
	aggregationModel.Status = status
	return model.EditOneById(service.ctx, aggregationModel.TableName(), id, &aggregationModel)
}

// ChangeAlertAggregationStatus 修改告警聚合状态
func (service *AlertAggregationService) ChangeAlertAggregationStatus(id uint64) error {
	alertAggregationModel := new(model.AlertAggregation)
	err := model.GetOneById(service.ctx, alertAggregationModel.TableName(), id, &alertAggregationModel)
	if err != nil {
		return err
	}
	if alertAggregationModel.Status == model.StrategyEnabled {
		alertAggregationModel.Status = model.StrategyDisabled
	} else {
		alertAggregationModel.Status = model.StrategyEnabled
	}
	return model.EditOneById(service.ctx, alertAggregationModel.TableName(), id, &alertAggregationModel)
}

// DeleteAlertAggregation 删除告警聚合
func (service *AlertAggregationService) DeleteAlertAggregation(id uint64) error {
	alertAggregationModel := new(model.AlertAggregation)
	err := model.GetOneById(service.ctx, alertAggregationModel.TableName(), id, &alertAggregationModel)
	if err != nil {
		return err
	}
	session := sessions.Default(service.c)
	uid := session.Get("uid").(uint64)
	alertAggregationModel.Status = model.StrategyDeleted
	alertAggregationModel.Uid = uid
	return model.EditOneById(service.ctx, alertAggregationModel.TableName(), id, &alertAggregationModel)
}
