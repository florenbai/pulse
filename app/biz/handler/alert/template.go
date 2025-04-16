package alert

import (
	"context"
	"darkhawk/app/biz/model"
	"darkhawk/app/biz/service"
	"darkhawk/conf"
	pb "darkhawk/hertz_gen/template"
	"darkhawk/pkg/response"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
)

func CreateTemplate(ctx context.Context, c *app.RequestContext) {
	var req pb.TemplateRequest
	if err := c.Bind(&req); err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err := service.NewAlertTemplateService(ctx, c).CreateTemplate(req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, response.Success)
	return
}

func UpdateTemplate(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	var req pb.TemplateRequest
	if err := c.Bind(&req); err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err = service.NewAlertTemplateService(ctx, c).UpdateTemplate(id, req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, response.Success)
	return
}

func UpdateTemplateChannel(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	var req pb.TemplateChannel
	if err := c.Bind(&req); err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err = service.NewAlertTemplateService(ctx, c).UpdateTemplateChannels(id, req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, response.Success)
}

func TemplateAll(ctx context.Context, c *app.RequestContext) {
	var err error
	alertTemplateModel := new(model.AlertTemplate)
	data, err := alertTemplateModel.All(ctx)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, data)
}

// TemplateList .
//
//	@Summary		template list
//	@Description	告警模板列表
//	@Tags			team
//	@Param			templateName	query	string	false
//	@Param			page		query	int		false	"page"
//	@Param			pageSize	query	int		false	"pageSize"
//	@router			/api/v1/template/list [GET]
func TemplateList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req pb.TemplateQuery
	if err = c.Bind(&req); err != nil {
		response.SendBaseResp(c, err)
		return
	}
	resp, err := service.NewAlertTemplateService(ctx, c).GetTemplateList(req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, resp)
}

func DeleteTemplate(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err = service.NewAlertTemplateService(ctx, c).DeleteTemplate(id)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, response.Success)
}

func GetTemplateDetail(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	data, err := service.NewAlertTemplateService(ctx, c).GetTemplateDetail(id)
	response.SendDataResp(c, response.Success, data)
	return
}

func WechatBot(ctx context.Context, c *app.RequestContext) {
	response.SendDataResp(c, response.Success, conf.GetConf().Umdp.WechatBotKeys)
	return
}

func GetChannelsByTemplate(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	var data []uint64
	channelTemplateModel := new(model.ChannelTemplate)
	err = model.GetPluck(ctx, channelTemplateModel.TableName(), "channel_id", &data, model.WhereWithScope("template_id", id))
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, data)
	return
}

func ChannelEnableGroup(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	data, err := service.NewAlertTemplateService(ctx, c).GetTemplateEnableChannels(id)
	response.SendDataResp(c, response.Success, data)
	return
}
