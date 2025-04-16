package mw

import (
	"context"
	"darkhawk/pkg/response"
	"darkhawk/pkg/zlog"
	"fmt"
	"strings"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
)

// NoMethodHandler 未找到请求方法的处理函数
func NoMethodHandler() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		var data = response.Response{Code: 405, Message: "方法不被允许"}
		ctx.JSON(405, data)
	}
}

// NoRouteHandler 未找到请求路由的处理函数
func NoRouteHandler() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		var data = response.Response{Code: 500, Message: "未找到请求路由的处理函数"}
		ctx.JSON(500, data)
	}
}

// SkipperFunc 定义中间件跳过函数
type SkipperFunc func(c context.Context, ctx *app.RequestContext) bool

// AllowPathPrefixSkipper 检查请求路径是否包含指定的前缀，如果包含则跳过
func AllowPathPrefixSkipper(prefixes ...string) SkipperFunc {
	return func(c context.Context, ctx *app.RequestContext) bool {
		path := string(ctx.Request.URI().RequestURI())
		//fmt.Println("path:", path)
		pathLen := len(path)

		if path == "/" {
			return true
		}

		for _, p := range prefixes {
			if pl := len(p); pathLen >= pl && path[:pl] == p {
				return true
			}
		}
		return false
	}
}

// AllowPathPrefixNoSkipper 检查请求路径是否包含指定的前缀，如果包含则不跳过
func AllowPathPrefixNoSkipper(prefixes ...string) SkipperFunc {
	return func(c context.Context, ctx *app.RequestContext) bool {
		path := string(ctx.Request.URI().RequestURI())
		pathLen := len(path)

		for _, p := range prefixes {
			if pl := len(p); pathLen >= pl && path[:pl] == p {
				return false
			}
		}
		return true
	}
}

// AllowMethodAndPathPrefixSkipper 检查请求方法和路径是否包含指定的前缀，如果包含则跳过
func AllowMethodAndPathPrefixSkipper(prefixes ...string) SkipperFunc {
	return func(c context.Context, ctx *app.RequestContext) bool {
		path := JoinRouter(string(ctx.Request.Method()), string(ctx.Request.URI().RequestURI()))
		pathLen := len(path)

		for _, p := range prefixes {
			if pl := len(p); pathLen >= pl && path[:pl] == p {
				return true
			}
		}
		return false
	}
}

// JoinRouter 拼接路由
func JoinRouter(method, path string) string {
	if len(path) > 0 && path[0] != '/' {
		path = "/" + path
	}
	return fmt.Sprintf("%s%s", strings.ToUpper(method), path)
}

func AccessLog() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		start := time.Now()
		ctx.Next(c)
		end := time.Now()
		latency := end.Sub(start).Microseconds
		zlog.AccessLog(fmt.Sprintf(`status=%d cost=%d method=%s full_path=%s client_ip=%s host=%s body: %s`,
			ctx.Response.StatusCode(), latency,
			ctx.Request.Header.Method(), ctx.Request.URI().PathOriginal(), ctx.ClientIP(), ctx.Request.Host(), string(ctx.Request.Body())))
	}
}
