package team

import (
	"darkhawk/app/biz/handler/team"
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
				_team := _v1.Group("/team", mw.UserAuthMw())
				_team.POST("", team.CreateTeam)
				_team.PUT("/:id", team.UpdateTeam)
				_team.PUT("/permission/:id", team.Authorization)
				_team.PUT("/status/:id", team.ChangeTeamStatus)
				_team.GET("/list", team.List)
				_team.GET("/all", team.All)
				_team.GET("/user/all/:team", team.TeamUserAll)
				_team.GET("/:id", team.TeamDetail)
				_team.DELETE("/:id", team.TeamDelete)
				_team.GET("/authorization/:id", team.AuthorizationInfo)
			}
		}
	}
}
