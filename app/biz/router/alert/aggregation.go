package alert

import (
	"darkhawk/app/biz/handler/alert"
	"darkhawk/app/biz/mw"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func RegisterAggregation(r *server.Hertz) {
	root := r.Group("/")
	{
		_api := root.Group("/api")
		{
			_v1 := _api.Group("/v1")
			{
				_aggregation := _v1.Group("/alert/aggregation", mw.UserAuthMw())
				_aggregation.POST("", alert.CreateAggregation)
				_aggregation.PUT("/:id", alert.EditAggregation)
				_aggregation.GET("/all/:id", alert.AggregationAll)
				_aggregation.DELETE("/:id", alert.DeleteAggregation)
				_aggregation.PUT("/status/:id", alert.ChangeAggregationStatus)
			}
		}
	}
}
