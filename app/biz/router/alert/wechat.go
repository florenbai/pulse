package alert

import (
	"darkhawk/app/biz/handler/wechat"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func RegisterWechat(r *server.Hertz) {
	root := r.Group("/")
	{
		_api := root.Group("/api")
		{
			_v1 := _api.Group("/v1")
			{
				_wechat := _v1.Group("/wechat")
				_wechat.GET("/alert", wechat.AlertList)
				_wechat.GET("/user", wechat.WxUserInfo)
			}
		}
	}
}
