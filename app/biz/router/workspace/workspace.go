package workspace

import (
	"darkhawk/app/biz/handler/workspace"
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
				_workspace := _v1.Group("/workspace", mw.UserAuthMw())
				_workspace.POST("", workspace.CreateWorkspace)
				_workspace.GET("/list", workspace.WorkspaceList)
				_workspace.GET("/all", workspace.All)
				_workspace.GET("/:id", workspace.BaseInfo)
				_workspace.PUT("/:id", workspace.UpdateBaseInfo)
				_workspace.DELETE("/:id", workspace.DeleteWorkspace)
				_workspace.PUT("/status/:id", workspace.ChangeStatus)
			}
		}
	}
}
