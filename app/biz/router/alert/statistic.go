package alert

import (
	"darkhawk/app/biz/handler/alert"
	"darkhawk/app/biz/mw"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func RegisterAlertStatistic(r *server.Hertz) {
	root := r.Group("/")
	{
		_api := root.Group("/api")
		{
			_v1 := _api.Group("/v1")
			{
				_statistic := _v1.Group("/alert/statistic", mw.UserAuthMw())
				_statistic.GET("/mtta/:id", alert.Mtta)
				_statistic.GET("/mtta", alert.AllMtta)
				_statistic.GET("/mttr/:id", alert.Mttr)
				_statistic.GET("/mttr", alert.AllMttr)
				_statistic.GET("/count/:id", alert.Count)
				_statistic.GET("/count", alert.AllCount)
				_statistic.GET("/level/:id", alert.Level)
				_statistic.GET("/level", alert.AllLevel)
				_statistic.GET("/top10", alert.TopTen)
				_statistic.GET("/level-alert", alert.LevelAlert)
				_statistic.GET("/alert-count", alert.AlertCount)
			}
		}
	}
}
