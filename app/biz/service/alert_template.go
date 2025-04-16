package service

import (
	"context"
	"darkhawk/app/biz/model"
	"darkhawk/hertz_gen/base"
	pb "darkhawk/hertz_gen/template"
	"darkhawk/pkg/response"
	"encoding/json"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/flosch/pongo2/v6"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
	"gorm.io/gorm"
)

type AlertTemplateService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewAlertTemplateService(ctx context.Context, c *app.RequestContext) *AlertTemplateService {
	return &AlertTemplateService{
		ctx: ctx,
		c:   c,
	}
}

func (service *AlertTemplateService) CreateTemplate(req pb.TemplateRequest) error {
	alertTemplateModel := new(model.AlertTemplate)
	ok, err := model.ExistName(service.ctx, alertTemplateModel.TableName(), "template_name", req.GetTemplateName(), nil)
	if err != nil {
		return err
	}
	if ok {
		return response.NameAlreadyExistErr
	}
	alertTemplateModel.TemplateName = req.GetTemplateName()
	alertTemplateModel.TemplateDesc = req.GetTemplateDesc()
	return alertTemplateModel.CreateTemplate(service.ctx, req.GetChannels())
}

func (service *AlertTemplateService) UpdateTemplate(id uint64, req pb.TemplateRequest) error {
	alertTemplateModel := new(model.AlertTemplate)
	err := model.GetOneById(service.ctx, alertTemplateModel.TableName(), id, &alertTemplateModel)
	if err != nil {
		return err
	}
	ok, err := model.ExistName(service.ctx, alertTemplateModel.TableName(), "template_name", req.GetTemplateName(), &id)
	if err != nil {
		return err
	}
	if ok {
		return response.NameAlreadyExistErr
	}

	alertTemplateModel.TemplateName = req.GetTemplateName()
	alertTemplateModel.TemplateDesc = req.GetTemplateDesc()
	return alertTemplateModel.UpdateTemplate(service.ctx, req.GetChannels())
}

func (service *AlertTemplateService) UpdateTemplateChannels(id uint64, req pb.TemplateChannel) error {
	channelTemplateModel := new(model.ChannelTemplate)
	err := model.GetOneById(service.ctx, channelTemplateModel.TableName(), id, &channelTemplateModel)
	if err != nil {
		return err
	}
	_, err = pongo2.FromString(req.Content)
	if err != nil {
		return fmt.Errorf("模板语法错误：%s", err.Error())
	}
	channelTemplateModel.Content = req.GetContent()
	channelTemplateModel.Finished = true

	return model.EditOneById(service.ctx, channelTemplateModel.TableName(), id, &channelTemplateModel)
}

func (service *AlertTemplateService) GetTemplateDetail(id uint64) (model.AlertTemplateDetail, error) {
	var data model.AlertTemplateDetail
	alertTemplateModel := new(model.AlertTemplate)
	err := model.GetOneById(service.ctx, alertTemplateModel.TableName(), id, &data)
	if err != nil {
		return data, err
	}
	channelTemplateModel := new(model.ChannelTemplate)
	data.Channels, err = channelTemplateModel.GetChannelTemplateDetail(service.ctx, id)
	if err != nil {
		return data, err
	}
	return data, nil
}

// GetTemplateList 获取模板列表
func (service *AlertTemplateService) GetTemplateList(req pb.TemplateQuery) (*base.BaseListResp, error) {
	search := map[string]string{
		"template_name": req.GetTemplateName(),
	}
	associate := map[string]string{
		"template_name": "template_name LIKE ?",
	}
	scopes := model.ParamWithScope(search, associate, nil, false)

	var data []model.AlertTemplate
	templateModel := new(model.AlertTemplate)
	i, err := model.GetPageList(service.ctx, templateModel.TableName(), req.GetPage(), req.GetPageSize(), &data, scopes)
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

func (service *AlertTemplateService) DeleteTemplate(id uint64) error {
	alertTemplateModel := new(model.AlertTemplate)
	err := model.GetOneById(service.ctx, alertTemplateModel.TableName(), id, &alertTemplateModel)
	if err != nil {
		return err
	}
	return model.Transaction(service.ctx, func(tx *gorm.DB) error {
		err := tx.Table(alertTemplateModel.TableName()).Where("id", id).Delete(alertTemplateModel).Error
		if err != nil {
			return err
		}
		channelTemplateModel := new(model.ChannelTemplate)
		err = tx.Table(channelTemplateModel.TableName()).Where("template_id", id).Delete(&channelTemplateModel).Error
		if err != nil {
			return err
		}
		return nil
	})
}

func (service *AlertTemplateService) GetTemplateEnableChannels(id uint64) ([]model.AlertChannel, error) {
	var data []model.AlertChannel
	alertChannelModel := new(model.AlertChannel)
	err := model.GetAll(service.ctx, alertChannelModel.TableName(), &data,
		model.JoinScope("LEFT JOIN channel_template ON channel_template.channel_id = alert_channel.id"),
		model.WhereWithScope("channel_template.template_id", id), model.WhereWithScope("channel_template.finished", true), model.SelectScope("alert_channel.*"))
	if err != nil {
		return nil, err
	}

	return data, nil
}
