package alert

import (
	"context"
	"darkhawk/app/biz/service"
	pb "darkhawk/hertz_gen/alert_event"
	"darkhawk/pkg/response"
	"darkhawk/pkg/zlog"
	"encoding/base64"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
	"strings"
)

// WorkspaceEventList .
//
//	@Description	告警事件列表
//	@Tags			team
//	@Param			workspaceName	query	string	false
//	@Param			workspaceName	query	string	false
//	@Param			page		query	int		false	"page"
//	@Param			pageSize	query	int		false	"pageSize"
//	@router			/api/v1/alert/event/workspace/list/{id} [GET]
func WorkspaceEventList(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	var req pb.WorkspaceEventListRequest
	if err := c.Bind(&req); err != nil {
		response.SendBaseResp(c, err)
		return
	}

	resp, err := service.NewAlertEventService(ctx, c).GetWorkspaceEventList(id, req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, resp)
}

// Push .
//
//	@Description	接收告警
//	@Tags			alert
//	@router			/api/v1/alert-event/push/{id} [GET]
func Push(ctx context.Context, c *app.RequestContext) {
	var integrationKey string
	authorization := c.Request.Header.Get("Authorization")
	if authorization != "" {
		// 验证认证类型是否为Basic
		if !strings.HasPrefix(authorization, "Basic ") {
			response.SendBaseResp(c, response.AuthorizeFailErr)
			return
		}
		// 去掉"Basic "前缀
		encoded := authorization[6:]
		token, err := base64.StdEncoding.DecodeString(encoded)
		if err != nil {
			response.SendBaseResp(c, response.AuthorizeFailErr)
			return
		}
		integrationKey = strings.ReplaceAll(string(token), ":", "")
	} else {
		integrationKey = c.Query("token")
	}

	// 处理事件
	err := service.NewAlertEventService(ctx, c).HandleEvent(integrationKey)
	if err != nil {
		zlog.Error(fmt.Sprintf("处理事件异常: %s", err.Error()))
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, response.Success)
	return
}

// EventDetail .
//
//	@Description	告警详情
//	@Tags			alert
//	@router			/api/v1/alert-event/{id} [GET]
func EventDetail(ctx context.Context, c *app.RequestContext) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	data, err := service.NewAlertEventService(ctx, c).GetWorkspaceDetail(id)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, data)
}

// EventList .
//
//	@Description	告警列表
//	@Tags			alert
//	@router			/api/v1/alert-event/list [GET]
func EventList(ctx context.Context, c *app.RequestContext) {
	var req pb.WorkspaceEventListRequest
	if err := c.Bind(&req); err != nil {
		response.SendBaseResp(c, err)
		return
	}
	data, err := service.NewAlertEventService(ctx, c).GetEventList(req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, data)
}

// ClaimEvent .
//
//	@Description	认领告警
//	@Tags			alert
//	@router			/api/v1/alert-event/claim [POST]
func ClaimEvent(ctx context.Context, c *app.RequestContext) {
	var req pb.ClaimEventRequest
	if err := c.Bind(&req); err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err := service.NewAlertEventService(ctx, c).ClaimAlertEvent(req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, response.Success)
	return
}

// CloseEvent .
//
//	@Description	关闭告警
//	@Tags			alert
//	@router			/api/v1/alert-event/close [POST]
func CloseEvent(ctx context.Context, c *app.RequestContext) {
	var req pb.CloseEventRequest
	if err := c.Bind(&req); err != nil {
		response.SendBaseResp(c, err)
		return
	}
	err := service.NewAlertEventService(ctx, c).CloseAlertEvent(req)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendBaseResp(c, response.Success)
	return
}
