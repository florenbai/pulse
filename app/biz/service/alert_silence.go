package service

import (
	"context"
	"darkhawk/app/biz/model"
	pb "darkhawk/hertz_gen/alert_silence"
	"darkhawk/pkg/response"
	"darkhawk/pkg/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

type AlertSilenceService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewAlertSilenceService(ctx context.Context, c *app.RequestContext) *AlertSilenceService {
	return &AlertSilenceService{
		ctx: ctx,
		c:   c,
	}
}
func (service *AlertSilenceService) CreateAlertSilence(req pb.SilenceRequest) error {
	alertSilenceModel := new(model.AlertSilence)
	ok, err := model.ExistName(service.ctx, alertSilenceModel.TableName(), "silence_name", req.GetSilenceName(), nil)
	if err != nil {
		return err
	}
	if ok {
		return response.NameAlreadyExistErr
	}
	if err = utils.AnyToAny(req, &alertSilenceModel); err != nil {
		return err
	}
	session := sessions.Default(service.c)
	uid := session.Get("uid").(uint64)
	alertSilenceModel.Uid = uid
	return model.Create(service.ctx, alertSilenceModel.TableName(), &alertSilenceModel)
}

func (service *AlertSilenceService) UpdateAlertSilence(id uint64, req pb.SilenceRequest) error {
	alertSilenceModel := new(model.AlertSilence)
	ok, err := model.ExistName(service.ctx, alertSilenceModel.TableName(), "silence_name", req.GetSilenceName(), &id)
	if err != nil {
		return err
	}
	if ok {
		return response.NameAlreadyExistErr
	}
	if err = utils.AnyToAny(req, &alertSilenceModel); err != nil {
		return err
	}
	session := sessions.Default(service.c)
	uid := session.Get("uid").(uint64)
	alertSilenceModel.Uid = uid
	return model.EditOneById(service.ctx, alertSilenceModel.TableName(), id, &alertSilenceModel)
}

func (service *AlertSilenceService) DeleteSilence(id uint64) error {
	alertSilenceModel := new(model.AlertSilence)
	err := model.GetOneById(service.ctx, alertSilenceModel.TableName(), id, &alertSilenceModel)
	if err != nil {
		return err
	}
	return model.DeleteOneById(service.ctx, alertSilenceModel.TableName(), id, &alertSilenceModel)
}
