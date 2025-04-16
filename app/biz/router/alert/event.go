package alert

import (
	"darkhawk/app/biz/handler/alert"
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
				_source := _v1.Group("/alert/source", mw.UserAuthMw())
				_source.GET("/all", alert.SourceAll)

				_template := _v1.Group("/alert/template", mw.UserAuthMw())
				_template.GET("/all", alert.TemplateAll)
				_template.GET("/list", alert.TemplateList)
				_template.POST("", alert.CreateTemplate)
				_template.PUT("/:id", alert.UpdateTemplate)
				_template.PUT("/channel/:id", alert.UpdateTemplateChannel)
				_template.DELETE("/:id", alert.DeleteTemplate)
				_template.GET("/detail/:id", alert.GetTemplateDetail)
				_template.GET("/channels/:id", alert.GetChannelsByTemplate)
				_template.GET("/wechat-bot", alert.WechatBot)
				_template.GET("/channel-group/:id", alert.ChannelEnableGroup)

				_event := _v1.Group("/alert/event", mw.UserAuthMw())
				_event.GET("/workspace-list/:id", alert.WorkspaceEventList)
				_event.GET("/list", alert.EventList)
				_event.GET("/:id", alert.EventDetail)
				_event.POST("/claim", alert.ClaimEvent)
				_event.POST("/closed", alert.CloseEvent)

			}
			_v1.POST("/alert-event/push", alert.Push, mw.AccessLog())
		}
	}
}
