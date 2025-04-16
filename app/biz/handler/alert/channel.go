package alert

import (
	"context"
	"darkhawk/app/biz/service"
	pb "darkhawk/hertz_gen/alert_channel"
	basepb "darkhawk/hertz_gen/base"
	"darkhawk/pkg/response"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
)

// CreateChannel .
//
//	@Description	创建告警渠道
//	@Tags			channel
//	@router			/api/v1/alert/channel [POST]
func CreateChannel(ctx context.Context, c *app.RequestContext) {
	var req pb.AlertChannelRequest
	if err := c.Bind(&req); err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err := service.NewAlertChannelService(ctx, c).CreateChannel(req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, response.Success)
	return
}

// UpdateChannel .
//
//	@Description	编辑告警渠道
//	@Tags			channel
//	@router			/api/v1/alert/channel/:id [PUT]
func UpdateChannel(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	var req pb.AlertChannelRequest
	if err := c.Bind(&req); err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err = service.NewAlertChannelService(ctx, c).UpdateChannel(id, req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, response.Success)
	return
}

// ChannelList .
//
//	@Summary		channel list
//	@Description	告警渠道列表
//	@Tags			channel
//	@Param			page		query	int		false	"page"
//	@Param			pageSize	query	int		false	"pageSize"
//	@router			/api/v1/channel/list [GET]
func ChannelList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req basepb.PageInfoReq
	if err = c.Bind(&req); err != nil {
		response.SendBaseResp(c, err)
		return
	}
	resp, err := service.NewAlertChannelService(ctx, c).ChannelList(req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, resp)
}

// DeleteChannel .
//
//	@Description	删除告警渠道
//	@Tags			channel
//	@router			/api/v1/alert/channel/:id [DELETE]
func DeleteChannel(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err = service.NewAlertChannelService(ctx, c).DeleteChannel(id)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, response.Success)
}

// ChannelGroup .
//
//	@Summary		channel group
//	@Description	告警渠道组信息
//	@Tags			channel
//	@router			/api/v1/channel/group [GET]
func ChannelGroup(ctx context.Context, c *app.RequestContext) {
	data, err := service.NewAlertChannelService(ctx, c).GetChanelGroupData()
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, data)
}
