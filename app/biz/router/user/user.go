package user

import (
	"darkhawk/app/biz/handler/user"
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
				_v1.POST("/user/login", user.Login)
				_user := _v1.Group("/user", mw.UserAuthMw())
				_user.GET("/all", user.All)
				_user.POST("", user.CreateUser)
				_user.PUT("/:id", user.UpdateUser)
				_user.GET("/:id", user.UserDetail)
				_user.POST("/logout", user.Logout)
				_user.GET("/list", user.List)
				_user.GET("/menu", user.Menu)
				_user.DELETE("/:id", user.Delete)
			}
		}
	}
}
