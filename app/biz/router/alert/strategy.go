package alert

import (
	"darkhawk/app/biz/handler/alert"
	"darkhawk/app/biz/mw"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func RegisterStrategy(r *server.Hertz) {
	root := r.Group("/")
	{
		_api := root.Group("/api")
		{
			_v1 := _api.Group("/v1")
			{
				_strategy := _v1.Group("/alert/strategy", mw.UserAuthMw())
				_strategy.POST("", alert.CreateStrategy)
				_strategy.PUT("/:id", alert.EditStrategy)
				_strategy.DELETE("/:id", alert.DeleteStrategy)
				_strategy.PUT("/sort", alert.ChangeStrategySort)
				_strategy.GET("/all/:id", alert.StrategyAll)
				_strategy.PUT("/status/:id", alert.ChangeStrategyStatus)

				_systemStrategy := _v1.Group("/system/strategy", mw.UserAuthMw())
				_systemStrategy.POST("", alert.CreateSystemStrategy)
				_systemStrategy.PUT("/:id", alert.EditSystemStrategy)
				_systemStrategy.DELETE("/:id", alert.DeleteSystemStrategy)
				_systemStrategy.GET("/:id", alert.SystemStrategyInfoByWorkspace)
				_systemStrategy.GET("/list", alert.SystemStrategyList)
				_systemStrategy.GET("/all", alert.SystemStrategyAll)

				_strategyLog := _v1.Group("/strategy/log", mw.UserAuthMw())
				_strategyLog.GET("/:id", alert.StrategyLog)
			}
		}
	}
}
