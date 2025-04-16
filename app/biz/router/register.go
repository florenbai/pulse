package router

import (
	"darkhawk/app/biz/router/alert"
	"darkhawk/app/biz/router/alert_rule"
	"darkhawk/app/biz/router/integration"
	"darkhawk/app/biz/router/menu"
	"darkhawk/app/biz/router/schedule"
	"darkhawk/app/biz/router/team"
	"darkhawk/app/biz/router/user"
	"darkhawk/app/biz/router/workspace"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func GeneratedRegister(r *server.Hertz) {
	workspace.Register(r)
	workspace.RegisterSource(r)
	alert.Register(r)
	alert.RegisterLevel(r)
	alert.RegisterStrategy(r)
	alert.RegisterAggregation(r)
	alert.RegisterSilence(r)
	alert.RegisterAlertLog(r)
	alert.RegisterAlertStatistic(r)
	alert.RegisterChannel(r)
	alert.RegisterRestrain(r)
	alert.RegisterWechat(r)
	alert_rule.Register(r)
	menu.Register(r)
	team.Register(r)
	user.Register(r)
	schedule.Register(r)
	integration.Register(r)
}
