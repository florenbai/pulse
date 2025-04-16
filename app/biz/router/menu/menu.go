package menu

import (
	"darkhawk/app/biz/handler/menu"
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
				_menu := _v1.Group("/menu", mw.UserAuthMw())
				_menu.GET("/list", menu.List)
				_menu.GET("/name-list", menu.NameList)
				_menu.POST("", menu.Create)
				_menu.PUT("/:id", menu.Edit)
				_menu.PUT("/status/:id", menu.ChangeStatus)
				_menu.DELETE("/:id", menu.Delete)
				_menu.PUT("/sort", menu.ChangeSort)
			}
		}
	}
}
