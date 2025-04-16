package alert

import (
	"context"
	"darkhawk/app/biz/model"
	"github.com/tidwall/gjson"
)

// GetAlertLevel 获取告警等级转换
// return 告警等级编号， 告警等级名称， 异常信息
func GetAlertLevel(alert []byte, field string, integrationId uint64) (uint64, string, error) {
	// 获取告警等级映射
	alertLevelNo := model.DefaultLevel
	levelMapModel := new(model.LevelIntegrationMapping)
	levelMap, err := levelMapModel.GetLevelSourceMapByIntegration(context.TODO(), integrationId)
	if err != nil {
		return 0, alertLevelNo, err
	}

	alertJson := gjson.ParseBytes(alert)

	if alertJson.Get(field).Exists() {
		alertLevel := alertJson.Get(field).String()
		if l, ok := levelMap[alertLevel]; ok {
			alertLevelNo = l
		}
	}
	alertLevelModel := new(model.AlertLevel)
	levelNameMap, err := alertLevelModel.LevelNameMap(context.TODO())
	if lid, ok := levelNameMap[alertLevelNo]; ok {
		return lid, alertLevelNo, nil
	}
	return 3, alertLevelNo, nil
}
