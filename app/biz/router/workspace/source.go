package workspace

import (
	"darkhawk/app/biz/handler/workspace"
	"darkhawk/app/biz/mw"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func RegisterSource(r *server.Hertz) {
	root := r.Group("/")
	{
		_api := root.Group("/api")
		{
			_v1 := _api.Group("/v1")
			{
				_workspace := _v1.Group("/workspace/alert-source", mw.UserAuthMw())
				_workspace.POST("/:id", workspace.RelatedAlertSource)
				_workspace.PUT("/:id", workspace.UpdateAlertSource)
				_workspace.PUT("/refresh-token/:id", workspace.RefreshSourceToken)
				_workspace.PUT("/level-mapping", workspace.LevelMapping)
				_workspace.DELETE("/:id", workspace.DeleteSource)
				_workspace.GET("/:id", workspace.WorkspaceSource)
			}
		}
	}
}
