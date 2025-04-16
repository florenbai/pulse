package alert

import (
	"context"
	"darkhawk/app/biz/model"
	"darkhawk/pkg/zlog"
	"fmt"
)

type EventRequest struct {
	AlertTitle  string            `json:"alertTitle" vd:"len($)>0;msg:'请输入告警标题'"`
	Description string            `json:"description" vd:"len($)>0;msg:'请输入告警描述'"`
	Level       string            `json:"level" vd:"regexp('^A[0-5]$');msg:'请输入正确的告警等级'"`
	LevelId     uint64            `json:"levelId"`
	Status      string            `json:"status" vd:"$=='resolved'||$=='firing';msg:'请输入正确的告警状态'" `
	AlertTime   *string           `json:"alertTime"`
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
	FingerPrint string            `json:"fingerprint" vd:"len($)>0; msg:'请输入告警指纹'"`
	SourceIp    string            `json:"sourceIp"`
}

func (alert *EventRequest) HandleEvent(ctx context.Context, workspaces []uint64, integration model.Integration) error {

	zlog.Debug("---------------------------------------------")
	alertEventModel := new(model.AlertEvent)
	alertEventModel.AlertTitle = alert.AlertTitle
	alertEventModel.Description = alert.Description
	alertEventModel.SourceId = integration.SourceId
	alertEventModel.IntegrationId = integration.Id
	if alert.LevelId > 0 {
		alertEventModel.Level = alert.LevelId
	} else {
		alertLevelModel := new(model.AlertLevel)
		levelNameMap, err := alertLevelModel.LevelNameMap(ctx)
		if err != nil {
			return err
		}
		if lid, ok := levelNameMap[alert.Level]; ok {
			alertEventModel.Level = lid
		} else {
			alertEventModel.Level = 3
		}
	}

	alertEventModel.Annotations = alert.Annotations
	alertEventModel.SourceIp = alert.SourceIp
	alertEventModel.Tags = alert.Labels
	alertEventModel.FingerPrint = alert.FingerPrint
	alertEventModel.IsRecovered = false
	if alert.Status == model.AlertResolved {
		alertEventModel.IsRecovered = true
	}
	var notifyWorkspaces []uint64
	var slogs = make(map[uint64]model.AlertSilence)
	for _, w := range workspaces {
		// 处理告警静默
		zlog.Debug(fmt.Sprintf("告警指纹：%s ，告警状态: %s, 开始在工作区：%d 进行静默匹配", alert.FingerPrint, alert.Status, w))
		// 静默, 如果静默，则不进入分配阶段
		isSilence, alertSilence, err := alert.CheckSilence(ctx, w)
		if err != nil {
			return err
		}
		if isSilence {
			slogs[w] = alertSilence
			zlog.Debug(fmt.Sprintf("告警指纹：%s ，告警状态: %s, 在工作区: %d 已触发静默", alert.FingerPrint, alert.Status, w))
			continue
		}
		zlog.Debug(fmt.Sprintf("告警指纹：%s ，告警状态: %s, 开始在工作区：%d 进行聚合匹配", alert.FingerPrint, alert.Status, w))
		//告警聚合
		aggInfo, err := alert.Aggregation(ctx, w)
		if err != nil {
			return err
		}
		// 当存在聚合窗口，并且在聚合窗口内，则不进行通知和分派
		if aggInfo != nil && !aggInfo.Notify {
			zlog.Debug(fmt.Sprintf("告警指纹：%s ，告警状态: %s, 开始在工作区：%d 已触发告警聚合", alert.FingerPrint, alert.Status, w))
			continue
		}
		// 告警抑制
		zlog.Debug(fmt.Sprintf("告警指纹：%s ，告警状态: %s, 开始在工作区：%d 进行告警抑制匹配", alert.FingerPrint, alert.Status, w))
		isRestrain, err := alert.CheckRestrain(ctx, w)
		if err != nil {
			return err
		}
		if isRestrain {
			zlog.Debug(fmt.Sprintf("告警指纹：%s ，告警状态: %s, 开始在工作区：%d 已触发告警抑制", alert.FingerPrint, alert.Status, w))
			continue
		}
		zlog.Debug(fmt.Sprintf("告警指纹：%s ，告警状态: %s, 开始在工作区：%d 告警抑制未被触发", alert.FingerPrint, alert.Status, w))
		notifyWorkspaces = append(notifyWorkspaces, w)

	}
	// 写入告警事件
	err := alertEventModel.CreateAlertEvent(ctx, workspaces, alertEventModel.IsRecovered)
	if err != nil {
		zlog.Debug(fmt.Sprintf("告警指纹：%s ，告警状态: %s, 写入告警事件失败：%s", alert.FingerPrint, alert.Status, err.Error()))
		return err
	}

	zlog.Debug(fmt.Sprintf("告警指纹：%s ，告警状态: %s, 写入告警事件完成", alert.FingerPrint, alert.Status))
	if len(slogs) > 0 {
		slModel := new(model.SilenceLog)
		go func() {
			for k, v := range slogs {
				err := model.Create(ctx, slModel.TableName(), &model.SilenceLog{
					AlertId:     alertEventModel.Id,
					SilenceId:   v.Id,
					WorkspaceId: k,
				})
				if err != nil {
					zlog.Error(err.Error())
				}
			}
		}()
	}
	for _, w := range notifyWorkspaces {
		zlog.Debug(fmt.Sprintf("告警指纹：%s ，告警状态: %s, 开始在工作区：%d 匹配告警通知策略", alert.FingerPrint, alert.Status, w))
		// 获取已开启的告警策略
		alertStrategyModel := new(model.AlertStrategy)
		strategyData, err := alertStrategyModel.EnableStrategy(ctx, w)
		if err != nil {
			return err
		}
		// 获取默认告警策略
		systemStrategyModel := new(model.SystemStrategy)
		err = systemStrategyModel.GetSystemStrategyByWorkspaceId(ctx, w)
		if err != nil {
			return err
		}
		// 告警通知策略匹配
		err = StrategyMatch(strategyData, *systemStrategyModel, *alertEventModel, alert.Level)
		if err != nil {
			return err
		}
	}

	zlog.Debug("---------------------------------------------")
	return nil
	/*
		alertLevelModel := new(model.AlertLevel)
		levelNameMap, err := alertLevelModel.LevelNameMap(context.Background())
		if err != nil {
			return err
		}
		if _, ok := levelNameMap[alert.Level]; !ok {
			return response.LevelNotFoundErr
		}
		// 告警聚合
		alertAgg := Alert{
			Name:  alert.AlertTitle,
			Label: alert.Labels,
			Level: levelNameMap[alert.Level],
		}
		aggInfo, err := alertAgg.Aggregation(ctx, workspaceSource.WorkspaceId)
		if err != nil {
			return err
		}
		// 当存在聚合窗口，并且在聚合窗口内，则不进行通知和分派
		if aggInfo != nil && !aggInfo.OutWindow {
			return nil
		}
		// 写入告警事件
		ct := model.LocalTime(time.Now())
		alertEventModel := new(model.AlertEvent)
		alertEventModel.AlertTitle = alert.AlertTitle
		alertEventModel.Description = alert.Description
		alertEventModel.SourceId = workspaceSource.Id
		alertEventModel.Level = levelNameMap[alert.Level]
		alertEventModel.Annotations = alert.Annotations
		addr := c.RemoteAddr().String()
		alertEventModel.SourceIp = strings.Split(addr, ":")[0]
		alertEventModel.Tags = alert.Labels
		alertEventModel.WorkspaceId = workspaceSource.WorkspaceId
		alertEventModel.FirstTriggerTime = &ct
		alertEventModel.TriggerTime = &ct
		alertEventModel.FingerPrint = alert.FingerPrint
		if alert.AlertTime != nil {
			at, err := time.Parse("2006-01-02 15:04:05", *alert.AlertTime)
			if err != nil {
				return err
			}
			ct = model.LocalTime(at)
			alertEventModel.FirstTriggerTime = &ct
		}
		if alert.Status == model.AlertResolved {
			alertEventModel.IsRecovered = true
		}
		err = alertEventModel.CreateAlertEvent(ctx, alertEventModel.IsRecovered)
		if err != nil {
			return err
		}
		if aggInfo != nil && !aggInfo.Notify {
			return nil
		}
		// 静默, 如果静默，则不进入分配阶段
		isSilence, err := alertAgg.CheckSilence(ctx, workspaceSource.WorkspaceId)
		if err != nil {
			return err
		}
		if isSilence {
			return nil
		}
		// 获取已开启的告警策略
		alertStrategyModel := new(model.AlertStrategy)
		strategyData, err := alertStrategyModel.EnableStrategy(ctx, workspaceSource.WorkspaceId)
		// 获取默认告警策略
		systemStrategyModel := new(model.SystemStrategy)
		err = systemStrategyModel.GetSystemStrategyByWorkspaceId(ctx, workspaceSource.WorkspaceId)
		if err != nil {
			return err
		}
		// 告警通知策略匹配
		err = StrategyMatch(strategyData, *systemStrategyModel, *alertEventModel, alert.Level)
		if err != nil {
			return err
		}
		return nil

	*/
}
