package alert

import (
	"darkhawk/app/biz/handler/alert"
	"darkhawk/app/biz/mw"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func RegisterLevel(r *server.Hertz) {
	root := r.Group("/")
	{
		_api := root.Group("/api")
		{
			_v1 := _api.Group("/v1")
			{
				_level := _v1.Group("/alert/level", mw.UserAuthMw())
				_level.GET("/all", alert.LevelAll)
				_level.GET("/all-mapping/:id", alert.LevelMappingAll)
				_level.PUT("/mapping", alert.UpdateLevelMapping)
				_level.POST("", alert.CreateLevel)
				_level.PUT("/:id", alert.UpdateLevel)
				_level.DELETE("/:id", alert.DeleteLevel)
			}
		}
	}
}
