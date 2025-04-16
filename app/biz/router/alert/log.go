package alert

import (
	"darkhawk/app/biz/handler/alert"
	"darkhawk/app/biz/mw"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func RegisterAlertLog(r *server.Hertz) {
	root := r.Group("/")
	{
		_api := root.Group("/api")
		{
			_v1 := _api.Group("/v1")
			{
				_log := _v1.Group("/alert/log", mw.UserAuthMw())
				_log.GET("/:id", alert.TimelineLog)
			}
		}
	}
}
