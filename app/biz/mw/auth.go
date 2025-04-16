package mw

import (
	"context"
	"darkhawk/conf"
	"darkhawk/pkg/jwt"
	"darkhawk/pkg/response"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
	"strings"
)

var (
	HeaderAuthorization = "Authorization"
	HeaderTag           = "userinfo"
)

// UserAuthMw 用户授权中间件
func UserAuthMw() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		token := string(c.GetHeader("Authorization"))
		if token == "" {
			response.SendBaseResp(c, response.UnauthorizedErr)
			c.Abort()
			return
		}
		// 移除Bearer前缀
		token = strings.Replace(token, "Bearer ", "", 1)
		claims, err := jwt.ParseToken(token, []byte(conf.GetConf().Authentication.AuthSecret))
		if err != nil {
			response.SendBaseResp(c, response.UnauthorizedErr)
			c.Abort()
			return
		}
		// 将用户信息存储到上下文中
		c.Set("uid", claims.UserID)
		c.Set("username", claims.Username)
		session := sessions.Default(c)
		session.Set("uid", claims.UserID)
		err = session.Save()
		if err != nil {
			response.SendBaseResp(c, err)
			return
		}
		c.Next(ctx)
	}
}
