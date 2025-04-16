package alert

import (
	"darkhawk/app/biz/handler/alert"
	"darkhawk/app/biz/mw"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func RegisterRestrain(r *server.Hertz) {
	root := r.Group("/")
	{
		_api := root.Group("/api")
		{
			_v1 := _api.Group("/v1")
			{
				_restrain := _v1.Group("/alert/restrain", mw.UserAuthMw())
				_restrain.POST("/:id", alert.CreateRestrain)
				_restrain.GET("/all/:id", alert.RestrainAll)
				_restrain.PUT("/:id", alert.UpdateRestrain)
				_restrain.DELETE("/:id", alert.DeleteRestrain)
			}
		}
	}
}
