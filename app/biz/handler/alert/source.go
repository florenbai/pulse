package alert

import (
	"context"
	"darkhawk/app/biz/model"
	"darkhawk/pkg/response"
	"github.com/cloudwego/hertz/pkg/app"
)

// SourceAll 获取告警源
func SourceAll(ctx context.Context, c *app.RequestContext) {
	var err error
	alertSourceModel := new(model.AlertSource)
	data, err := alertSourceModel.All(ctx)
	if err != nil {
		response.SendBaseResp(c, err)
		return
	}
	response.SendDataResp(c, response.Success, data)
}
