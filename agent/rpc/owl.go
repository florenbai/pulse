package rpc

import (
	"agent/cliconfig"
	agentpb "agent/kitex_gen/agent"
	"agent/kitex_gen/agent/agentservice"
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/kitex-contrib/registry-nacos/resolver"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"net"
	"time"
)

var owlClient agentservice.Client

const OwlAgentServiceName = "owl"

func InitOwlAgentRpc() {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(cliconfig.GetConf().Nacos.Address, cliconfig.GetConf().Nacos.Port),
	}

	cc := constant.ClientConfig{
		NamespaceId:         cliconfig.GetConf().Nacos.NamespaceId,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "info",
		Username:            cliconfig.GetConf().Nacos.Username,
		Password:            cliconfig.GetConf().Nacos.Password,
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
	r := resolver.NewNacosResolver(cli)
	owlClient = agentservice.MustNewClient(
		OwlAgentServiceName,
		client.WithResolver(r),
		client.WithRPCTimeout(time.Second*10),
	)
}

func Ping(ctx context.Context) error {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return err
	}
	var clientIp string
	for _, address := range addrs {
		// 检查地址是否为IPv4类型
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				clientIp = ipnet.IP.String()
			}
		}
	}
	req := agentpb.PingRequest{
		Client:     clientIp,
		SourceSign: cliconfig.GetConf().Prometheus.SourceSign,
	}
	_, err = owlClient.Ping(ctx, &req)
	if err != nil {
		return err
	}

	return nil
}

func UpdateConfig(ctx context.Context) error {
	req := agentpb.PingRequest{
		SourceSign: cliconfig.GetConf().Prometheus.SourceSign,
	}
	_, err := owlClient.UpdateConfig(ctx, &req)
	if err != nil {
		return err
	}

	return nil
}
