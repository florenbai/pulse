package integration

import (
	"darkhawk/app/biz/handler/integration"
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
				_integration := _v1.Group("/integration", mw.UserAuthMw())
				_integration.GET("/list", integration.List)
				_integration.POST("", integration.Create)
				_integration.PUT("/:id", integration.Update)
				_integration.PUT("/status/:id", integration.ChangeStatus)
				_integration.DELETE("/:id", integration.Delete)
				_integration.GET("/detail/:id", integration.Detail)

				_integrationRouter := _v1.Group("/integration/router", mw.UserAuthMw())
				_integrationRouter.POST("/:id", integration.CreateRouter)
				_integrationRouter.DELETE("/:id", integration.DeleteRouter)
				_integrationRouter.PUT("/:id", integration.UpdateRouter)
				_integrationRouter.PUT("/sort", integration.ChangeSort)
				_integrationRouter.GET("/all/:id", integration.GetRouters)

				_tagRewrite := _v1.Group("/integration/tag-rewrite", mw.UserAuthMw())
				_tagRewrite.POST("/:id", integration.CreateTagRewrite)
				_tagRewrite.PUT("/:id", integration.UpdateTagRewrite)
				_tagRewrite.DELETE("/:id", integration.DeleteTagRewrite)
				_tagRewrite.GET("/all/:id", integration.TagRewriteList)
			}

		}
	}
}
