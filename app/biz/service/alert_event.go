package service

import (
	"context"
	"darkhawk/app/biz/dal/mysql"
	"darkhawk/app/biz/model"
	"darkhawk/app/biz/pack/alert"
	pb "darkhawk/hertz_gen/alert_event"
	"darkhawk/hertz_gen/base"
	"darkhawk/pkg/response"
	"darkhawk/pkg/utils"
	"darkhawk/pkg/zlog"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
	"gorm.io/gorm"
	"time"
)

type AlertEventService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewAlertEventService(ctx context.Context, c *app.RequestContext) *AlertEventService {
	return &AlertEventService{
		ctx: ctx,
		c:   c,
	}
}

func (service AlertEventService) GetWorkspaceEventList(id uint64, req pb.WorkspaceEventListRequest) (*base.BaseListResp, error) {
	search := map[string]string{
		"alert_title":  req.GetAlertTitle(),
		"trigger_time": utils.GetTimeAgo(req.GetTimeRange()).Format("2006-01-02 15:04:05"),
		"level":        "",
	}

	if req.GetAlertLevel() != 0 {
		search["level"] = fmt.Sprintf("%d", req.GetAlertLevel())
	}
	associate := map[string]string{
		"alert_title":  "alert_title LIKE ?",
		"trigger_time": "trigger_time >= ?",
		"level":        "alert_event.level = ?",
	}
	scopes := model.NewScopeContainer()
	scopes = append(scopes, model.ParamWithScope(search, associate, nil, false))
	if req.GetProgress() != 0 {
		if req.GetProgress() > 3 {
			scopes = append(scopes, model.InWithScope("progress", []model.AlertStatus{model.EventUnClaimed, model.EventHasClaimed}))
		} else {
			scopes = append(scopes, model.WhereWithScope("progress", req.GetProgress()))
		}
	}
	alertEventModel := new(model.AlertEvent)
	data, i, err := alertEventModel.GetWorkspaceEventList(service.ctx, id, req.GetPage(), req.GetPageSize(), scopes...)
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

func (service AlertEventService) GetEventList(req pb.WorkspaceEventListRequest) (*base.BaseListResp, error) {
	search := map[string]string{
		"alert_title":  req.GetAlertTitle(),
		"trigger_time": utils.GetTimeAgo(req.GetTimeRange()).Format("2006-01-02 15:04:05"),
		"level":        "",
	}

	if req.GetAlertLevel() != 0 {
		search["level"] = fmt.Sprintf("%d", req.GetAlertLevel())
	}
	associate := map[string]string{
		"alert_title":  "alert_title LIKE ?",
		"trigger_time": "trigger_time >= ?",
		"level":        "alert_event.level = ?",
	}
	scopes := model.NewScopeContainer()
	scopes = append(scopes, model.ParamWithScope(search, associate, nil, false))
	if req.GetProgress() != 0 {
		if req.GetProgress() > 3 {
			scopes = append(scopes, model.InWithScope("progress", []model.AlertStatus{model.EventUnClaimed, model.EventHasClaimed}))
		} else {
			scopes = append(scopes, model.WhereWithScope("progress", req.GetProgress()))
		}
	}
	alertEventModel := new(model.AlertEvent)
	data, i, err := alertEventModel.GetEventList(service.ctx, req.GetPage(), req.GetPageSize(), scopes...)
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

func (service AlertEventService) HandleEvent(token string) error {
	integrationModel := new(model.Integration)
	err := model.GetOneWithScope(service.ctx, integrationModel.TableName(), &integrationModel, model.WhereWithScope("token", token))
	if err != nil || integrationModel.Id == 0 {
		return response.AuthorizeFailErr
	}
	if integrationModel.Disabled {
		return response.ChannelStatusCloseErr
	}
	alertSourceModel := new(model.AlertSource)
	err = model.GetOneById(service.ctx, alertSourceModel.TableName(), integrationModel.SourceId, &alertSourceModel)
	if err != nil {
		return err
	}
	// 获取路由
	integrationRouterModel := new(model.IntegrationRouter)
	var routerData []model.IntegrationRouter
	err = model.GetAll(service.ctx, integrationRouterModel.TableName(), &routerData, model.WhereWithScope("integration_id", integrationModel.Id), model.OrderScope("sort ASC"))
	if err != nil {
		return err
	}
	// 标签重写
	tagModel := new(model.TagRewrite)
	var tagRewriteData []model.TagRewrite
	err = model.GetAll(service.ctx, tagModel.TableName(), &tagRewriteData, model.WhereWithScope("integration_id", integrationModel.Id))
	if err != nil {
		return err
	}

	var events []alert.EventRequest
	switch alertSourceModel.SourceName {
	case "Prometheus":
		var req PrometheusHookRequest
		if err := service.c.Bind(&req); err != nil {
			return err
		}
		prometheus := NewPrometheusAlertService(service.ctx, service.c)
		events, err = prometheus.GetEvent(req.Alerts, *integrationModel)
		if err != nil {
			return err
		}
	case "自定义事件":
		var req alert.EventRequest
		if err := service.c.BindAndValidate(&req); err != nil {
			return err
		}
		req.SourceIp = service.c.ClientIP()
		events = append(events, req)
	case "阿里云监控 CM":
		var req AliyunCMSHookRequest
		if err := service.c.Bind(&req); err != nil {
			return err
		}
		events, err = NewAliyunCMSService(service.ctx, service.c).GetEvent(req, *integrationModel)
	case "阿里云监控 事件":
		var req AliyunEventRequest
		if err := service.c.Bind(&req); err != nil {
			return err
		}
		events, err = NewAliyunEventService(service.ctx, service.c).GetEvent(req, *integrationModel)
	}
	workspaceModel := new(model.Workspace)
	for _, v := range events {
		// 标签重写
		v.Labels = alert.TagRewrite(v.Labels, tagRewriteData)
		// 遍历路由,获取要发送到的工作区
		var workspaceRouter []uint64
		for _, router := range routerData {
			if len(router.Filters) > 0 {
				intoRouter := false
				for _, rf := range router.Filters {
					if alert.HasTag(rf.Values, v.Labels) {
						intoRouter = true
					}
				}
				if !intoRouter {
					continue
				}
			}
			workspaceRouter = append(workspaceRouter, router.Workspaces...)
			if router.Next == 0 {
				break
			}
		}
		workspaceRouter = utils.RemoveDuplicate(workspaceRouter)
		// 获取已开启的工作区
		ws, err := workspaceModel.GetEnableWorkspace(service.ctx, workspaceRouter)
		if err != nil {
			zlog.Error(fmt.Sprintf("获取可用工作区失败，%v", err.Error()))
			continue
		}
		if len(ws) == 0 {
			zlog.Debug("---------------------------------------------")
			zlog.Debug(fmt.Sprintf("告警 %s : %s 未匹配到路由", v.AlertTitle, v.FingerPrint))
			zlog.Debug("---------------------------------------------")
			continue
		}
		err = v.HandleEvent(service.ctx, ws, *integrationModel)
		if err != nil {
			zlog.Error(fmt.Sprintf("处理告警事件失败，%v", err.Error()))
			continue
		}

	}
	return nil
}

// CreateAlertEvent 创建告警事件
func (service AlertEventService) CreateAlertEvent(alert model.AlertEvent) (alertEvent *model.AlertEvent, notify bool, err error) {
	// 获取前置告警事件未恢复的告警事件
	notify = true
	alertEvent = new(model.AlertEvent)
	err = model.GetOneWithScope(service.ctx, alertEvent.TableName(), &alertEvent, model.WhereWithScope("alert_title", alert.AlertTitle), model.WhereWithScope("level", alert.Level), model.WhereWithScope("is_recovered", false))
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, false, err
	}
	ct := model.LocalTime(time.Now())

	if alertEvent.Id > 0 && alertEvent.Progress != model.EventClosed {
		// 如果告警状态为恢复,设置状态并设置恢复时间
		if alert.IsRecovered {
			alertEvent.IsRecovered = true
			alertEvent.RecoverTime = &ct
		} else {
			// 如果告警状态未恢复，则累加通知次数和最新告警时间
			alertEvent.NotifyCurNumber = alertEvent.NotifyCurNumber + 1
			alertEvent.TriggerTime = &ct
		}
		err = model.EditOneById(service.ctx, alertEvent.TableName(), alertEvent.Id, &alertEvent)
		if err != nil {
			return nil, false, err
		}
	} else {
		// 创建新的告警事件

		err = model.Create(service.ctx, alertEvent.TableName(), &alert)
		if err != nil {
			return nil, false, err
		}
		alertEvent = &alert
	}

	return
}

// 获取告警详情
func (service AlertEventService) GetWorkspaceDetail(id uint64) (*model.AlertEvent, error) {
	alertEventModel := new(model.AlertEvent)
	err := model.GetOneById(service.ctx, alertEventModel.TableName(), id, &alertEventModel)
	if err != nil {
		return nil, err
	}
	return alertEventModel, nil
}

// ClaimAlertEvent 领取告警事件
func (service AlertEventService) ClaimAlertEvent(req pb.ClaimEventRequest) error {
	session := sessions.Default(service.c)
	uid := session.Get("uid").(uint64)
	// 校验告警编号是否可被领取
	err := service.checkAlertOwner(req.GetEvents(), uid)
	if err != nil {
		return err
	}
	// 检测告警是否被领取
	err = service.checkAlertEnable(req.GetEvents())
	if err != nil {
		return err
	}
	alertModel := new(model.AlertEvent)
	return alertModel.ClaimAlertEvent(service.ctx, req.GetEvents(), uid)
}

// CloseAlertEvent 关闭告警事件
func (service AlertEventService) CloseAlertEvent(req pb.CloseEventRequest) error {
	session := sessions.Default(service.c)
	uid := session.Get("uid").(uint64)
	// 检测告警是否被关闭
	err := service.checkAlertEnableClose(req.GetEvents(), uid)
	if err != nil {
		return err
	}
	err = service.checkAlertOwner(req.GetEvents(), uid)
	if err != nil {
		return err
	}

	alertModel := new(model.AlertEvent)
	return alertModel.CloseAlertEvent(service.ctx, req.GetEvents(), req.GetIsRecovered(), uid)
}

func (service AlertEventService) checkAlertEnableClose(event []uint64, uid uint64) error {
	alertModel := new(model.AlertEvent)
	for _, v := range event {
		ok, err := alertModel.CheckCloseUser(service.ctx, v, uid)
		if err != nil {
			return err
		}
		if !ok {
			return response.NewErrNo(response.Err_BadRequest, fmt.Sprintf("您没有权限关闭编号为【%d】的告警事件", v))
		}
	}
	return nil
}

// checkAlertOwner 检测告警是否可以被领取
func (service AlertEventService) checkAlertOwner(event []uint64, uid uint64) error {
	strategyLog := new(model.StrategyLog)

	var strategySets []model.StrategySet
	for _, v := range event {
		strategyList, err := strategyLog.GetDistinctStrategyByEvent(service.ctx, v)
		if err != nil {
			return err
		}
		// 获取所有分派
		for _, s := range strategyList {
			if s.StrategyType == model.DefaultStrategy {
				systemStrategyModel := new(model.SystemStrategy)
				err := model.GetOneById(service.ctx, systemStrategyModel.TableName(), s.StrategyId, &systemStrategyModel)
				if err != nil {
					return err
				}
				strategySets = append(strategySets, systemStrategyModel.StrategySet[0])
			} else {
				customStrategyModel := new(model.AlertStrategy)
				err := model.GetOneById(service.ctx, customStrategyModel.TableName(), s.StrategyId, &customStrategyModel)
				if err != nil {
					return err
				}
				strategySets = append(strategySets, customStrategyModel.StrategySet[0])
			}
		}
		// 获取分派人员
		var members []uint64
		for _, strategyItem := range strategySets {
			members = append(members, strategyItem.Members...)
			if len(strategyItem.Teams) > 0 {
				teamModel := new(model.Teams)
				teamMembers, err := teamModel.GetUsersByTeam(service.ctx, strategyItem.Teams)
				if err != nil {
					return err
				}
				members = append(members, teamMembers...)
			}
			// 当前月份的所有值班成员都可以处理
			if len(strategyItem.Schedules) > 0 {
				planModel := new(model.SchedulePlan)
				month := time.Now().Format("2006-01")
				users, err := planModel.GetUsersByIds(service.ctx, strategyItem.Schedules, month)
				if err != nil {
					return err
				}
				members = append(members, users...)
			}
		}
		// 工作区权限
		alertWorkspaceModel := new(model.AlertWorkspace)
		var ws []uint64
		err = model.GetPluck(service.ctx, alertWorkspaceModel.TableName(), "workspace_id", &ws, model.WhereWithScope("alert_id", v))
		if err != nil {
			return err
		}
		var teamUsers []uint64
		err = model.GetPluck(service.ctx, mysql.TeamUserTable, "user_teams.user_id", &teamUsers, model.JoinScope("LEFT JOIN workspace_team ON user_teams.team_id = workspace_team.team_id"), model.WhereWithScope("workspace_team.workspace_id IN ?", ws))
		if err != nil {
			return err
		}

		if !utils.InOfT(uid, members) && !utils.InOfT(uid, teamUsers) {
			return response.NewErrNo(response.Err_BadRequest, fmt.Sprintf("您没有权限领取编号为【%d】的告警事件", v))
		}
	}
	return nil
}

// checkAlertEnable 检测告警是否被领取
func (service AlertEventService) checkAlertEnable(event []uint64) error {
	var disables []uint64
	alertModel := new(model.AlertEvent)
	err := model.GetPluck(service.ctx, alertModel.TableName(), "id", &disables, model.InWithScope("id", event), model.InWithScope("progress", []model.AlertStatus{model.EventHasClaimed, model.EventClosed}))
	if err != nil {
		return err
	}
	if len(disables) > 0 {
		return response.NewErrNo(response.Err_BadRequest, fmt.Sprintf("编号为 %v 的告警事件已被领取", disables))
	}
	return nil
}
