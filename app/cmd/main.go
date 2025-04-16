package main

import (
	"darkhawk/app/biz/dal"
	"darkhawk/app/biz/handler/alert"
	"darkhawk/app/biz/router"
	"darkhawk/conf"
	"darkhawk/kitex_gen/agent/agentservice"
	"darkhawk/rpc/agent"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	kitexServer "github.com/cloudwego/kitex/server"
	"github.com/hertz-contrib/cors"
	"github.com/hertz-contrib/gzip"
	"github.com/hertz-contrib/logger/accesslog"
	"github.com/hertz-contrib/sessions"
	"github.com/hertz-contrib/sessions/cookie"
	"github.com/kitex-contrib/registry-nacos/registry"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"net"
	"time"
)

func init() {
	//初始化数据库
	dal.Init()
}

func main() {
	h := server.Default(server.WithHostPorts(conf.GetConf().Server.Addr))
	h.Use(gzip.Gzip(gzip.DefaultCompression))
	h.Use(accesslog.New(accesslog.WithFormat("[${time}] ${status} - ${latency} ${method} ${path} ${queryParams}")))
	h.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	store := cookie.NewStore([]byte(conf.GetConf().Authentication.AuthSecret))

	h.Use(sessions.New(fmt.Sprintf("%s_session", conf.Dom), store))

	router.GeneratedRegister(h)
	// 如果开启远程Prometheus 则使用krpc
	if conf.GetConf().Server.RemotePrometheus {
		agent.InitPrometheusAgentRpc()

		go func() {
			sc := []constant.ServerConfig{
				*constant.NewServerConfig(conf.GetConf().Nacos.Addr, conf.GetConf().Nacos.Port),
			}
			cc := constant.ClientConfig{
				NamespaceId:         conf.GetConf().Nacos.Namespace,
				TimeoutMs:           5000,
				NotLoadCacheAtStart: true,
				LogDir:              conf.GetConf().Nacos.LogDir,
				CacheDir:            conf.GetConf().Nacos.CacheDir,
				LogLevel:            conf.GetConf().Nacos.LogLevel,
				Username:            conf.GetConf().Nacos.Username,
				Password:            conf.GetConf().Nacos.Password,
			}

			cli, err := clients.NewNamingClient(
				vo.NacosClientParam{
					ClientConfig:  &cc,
					ServerConfigs: sc,
				},
			)
			if err != nil {
				klog.Fatal(err)
			}
			addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Server.GrpcAddr)
			if err != nil {
				klog.Fatal(err)
			}
			svr := agentservice.NewServer(new(alert.AgentServiceImpl),
				kitexServer.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.GetConf().Server.RpcName}),
				kitexServer.WithRegistry(registry.NewNacosRegistry(cli)),
				kitexServer.WithServiceAddr(addr),
			)
			if err := svr.Run(); err != nil {
				klog.Fatal(err)
			}
			/*
				addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Server.GrpcAddr)
				if err != nil {
					klog.Fatal(err)
				}
				svr := kitexServer.NewServer(kitexServer.WithServiceAddr(addr))
				err = agentservice.RegisterService(svr, new(alert.AgentServiceImpl))
				if err != nil {
					klog.Fatal(err)
				}
				err = svr.Run()
				if err != nil {
					klog.Error(err)
				}

			*/

		}()
	}

	h.Spin()
}
