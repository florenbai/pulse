package service

import (
	"context"
	"darkhawk/app/biz/model"
	pb "darkhawk/hertz_gen/alert_channel"
	basepb "darkhawk/hertz_gen/base"
	"darkhawk/pkg/response"
	"encoding/json"
	"github.com/cloudwego/hertz/pkg/app"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
	"gorm.io/gorm"
)

type AlertChannelService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewAlertChannelService(ctx context.Context, c *app.RequestContext) *AlertChannelService {
	return &AlertChannelService{
		ctx: ctx,
		c:   c,
	}
}

func (service *AlertChannelService) CreateChannel(req pb.AlertChannelRequest) error {
	alertChannelModel := new(model.AlertChannel)
	ok, err := model.ExistName(service.ctx, alertChannelModel.TableName(), "channel_name", req.GetChannelName(), nil)
	if err != nil {
		return err
	}
	if ok {
		return response.NameAlreadyExistErr
	}
	alertChannelModel.ChannelName = req.GetChannelName()
	alertChannelModel.ChannelSign = req.GetChannelSign()
	alertChannelModel.ChannelType = req.GetChannelType()
	alertChannelModel.ChannelGroup = req.GetChannelGroup()
	return model.Create(service.ctx, alertChannelModel.TableName(), &alertChannelModel)
}

func (service *AlertChannelService) UpdateChannel(id uint64, req pb.AlertChannelRequest) error {
	alertChannelModel := new(model.AlertChannel)
	err := model.GetOneById(service.ctx, alertChannelModel.TableName(), id, &alertChannelModel)
	if err != nil {
		return err
	}

	ok, err := model.ExistName(service.ctx, alertChannelModel.TableName(), "channel_name", req.GetChannelName(), &id)
	if err != nil {
		return err
	}
	if ok {
		return response.NameAlreadyExistErr
	}
	alertChannelModel.ChannelName = req.GetChannelName()
	alertChannelModel.ChannelSign = req.GetChannelSign()
	alertChannelModel.ChannelType = req.GetChannelType()
	alertChannelModel.ChannelGroup = req.GetChannelGroup()
	return model.EditOneById(service.ctx, alertChannelModel.TableName(), id, &alertChannelModel)
}

func (service *AlertChannelService) DeleteChannel(id uint64) error {
	alertChannelModel := new(model.AlertChannel)
	err := model.GetOneById(service.ctx, alertChannelModel.TableName(), id, &alertChannelModel)
	if err != nil {
		return err
	}
	return model.Transaction(service.ctx, func(tx *gorm.DB) error {
		err := tx.Table(alertChannelModel.TableName()).Where("id", id).Delete(&alertChannelModel).Error
		if err != nil {
			return err
		}
		channelTemplateModel := new(model.ChannelTemplate)
		err = tx.Table(channelTemplateModel.TableName()).Where("channel_id", id).Delete(&channelTemplateModel).Error
		if err != nil {
			return err
		}
		return nil
	})
}

func (service *AlertChannelService) ChannelList(p basepb.PageInfoReq) (*basepb.BaseListResp, error) {
	var data []model.AlertChannel
	alertChannelModel := new(model.AlertChannel)
	i, err := model.GetPageList(service.ctx, alertChannelModel.TableName(), p.GetPage(), p.GetPageSize(), &data)
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

func (service *AlertChannelService) GetChanelGroupData() ([]model.AlertChannel, error) {

	var data []model.AlertChannel
	alertChannelModel := new(model.AlertChannel)
	err := model.GetAll(service.ctx, alertChannelModel.TableName(), &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
