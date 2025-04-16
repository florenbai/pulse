package alert_rule

import (
	alert_rule "darkhawk/app/biz/handler/alert-rule"
	"darkhawk/app/biz/mw"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func Register(r *server.Hertz) {
	root := r.Group("/")
	{
		_api := root.Group("/api")
		{
			_v1 := _api.Group("/v1")
			{
				_alert_rule := _v1.Group("/alert-rule", mw.UserAuthMw())
				_alert_rule.POST("/source", alert_rule.CreateSource)
				_alert_rule.PUT("/source/:id", alert_rule.UpdateSource)
				_alert_rule.GET("/source/list", alert_rule.List)
				_alert_rule.GET("/source/all", alert_rule.AllSource)
				_alert_rule.DELETE("/source/:id", alert_rule.DeleteSource)
				_alert_rule.GET("/source/agent/:id", alert_rule.AgentList)
				_alert_rule.POST("/source/reload/:id", alert_rule.Reload)
				_alert_rule.GET("/list/gid/:id", alert_rule.RuleList)
				_alert_rule.POST("", alert_rule.Create)
				_alert_rule.PUT("/:id", alert_rule.Update)
				_alert_rule.DELETE("/:id", alert_rule.DeleteRule)

				_alert_rule_group := _v1.Group("/rule-group", mw.UserAuthMw())
				_alert_rule_group.GET("/all", alert_rule.RuleGroups)
				_alert_rule_group.POST("", alert_rule.CreateRuleGroup)
				_alert_rule_group.DELETE("/:id", alert_rule.DeleteRuleGroup)
				_alert_rule_group.GET("/label/:id", alert_rule.GetLabelsByRuleGroup)
			}
		}
	}
}
