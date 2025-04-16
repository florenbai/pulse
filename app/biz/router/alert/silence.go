package alert

import (
	"darkhawk/app/biz/handler/alert"
	"darkhawk/app/biz/mw"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func RegisterSilence(r *server.Hertz) {
	root := r.Group("/")
	{
		_api := root.Group("/api")
		{
			_v1 := _api.Group("/v1")
			{
				_silence := _v1.Group("/alert/silence", mw.UserAuthMw())
				_silence.POST("", alert.CreateSilence)
				_silence.GET("/all/:id", alert.SilenceAll)
				_silence.PUT("/:id", alert.EditSilence)
				_silence.DELETE("/:id", alert.DeleteSilence)
			}
		}
	}
}
