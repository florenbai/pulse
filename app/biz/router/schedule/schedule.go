package schedule

import (
	"darkhawk/app/biz/handler/schedule"
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
				_schedule := _v1.Group("/schedule", mw.UserAuthMw())
				_schedule.POST("", schedule.CreatedSchedule)
				_schedule.GET("/list", schedule.List)
				_schedule.GET("/all", schedule.All)
				_schedule.DELETE("/:id", schedule.DeleteSchedule)

				_schedulePlan := _v1.Group("/schedule-plan", mw.UserAuthMw())
				_schedulePlan.GET("/list/:id", schedule.PlanList)
				_schedulePlan.PUT("/:id", schedule.PlanIn)
				_schedulePlan.PUT("/clear/:id", schedule.PlanOut)
			}
		}
	}
}
