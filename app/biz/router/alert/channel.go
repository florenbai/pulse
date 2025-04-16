package alert

import (
	"darkhawk/app/biz/handler/alert"
	"darkhawk/app/biz/mw"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func RegisterChannel(r *server.Hertz) {
	root := r.Group("/")
	{
		_api := root.Group("/api")
		{
			_v1 := _api.Group("/v1")
			{
				_channel := _v1.Group("/alert/channel", mw.UserAuthMw())
				_channel.POST("", alert.CreateChannel)
				_channel.PUT("/:id", alert.UpdateChannel)
				_channel.GET("/list", alert.ChannelList)
				_channel.DELETE("/:id", alert.DeleteChannel)
				_channel.GET("/group", alert.ChannelGroup)
			}
		}
	}
}
