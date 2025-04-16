package clicommand

import (
	"agent/agent"
	"agent/cliconfig"
	"agent/kitex_gen/prometheus/prometheusagentservice"
	"agent/rpc"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	kitexServer "github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/registry-nacos/registry"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/urfave/cli"
	"net"
	"os"
)

const startDescription = `Usage:

    owl-agent prometheus [options...]

Description:

The agent will run any jobs within a PTY (pseudo terminal) if available.

Example:

    $ owl-agent prometheus --config=config.yaml`

type OwlAgentConfig struct {
	Server     OwlServer  `yaml:"server"`
	Consul     Consul     `yaml:"consul"`
	Prometheus Prometheus `yaml:"prometheus"`
}

type OwlServer struct {
	LogLevel string `yaml:"logLevel"`
}

type Consul struct {
	Address string `yaml:"address"`
}

type Prometheus struct {
	Host string `yaml:"host"`
}

var OwlAgentCommand = cli.Command{
	Name:        "prometheus",
	Usage:       "Starts a prometheus agent",
	Description: startDescription,

	Flags: []cli.Flag{
		cli.StringFlag{
			Name:   "config",
			Value:  "",
			Usage:  "Path to a configuration file",
			EnvVar: "OWL_AGENT_CONFIG",
		},
	},

	Action: func(c *cli.Context) error {
		cliconfig.ConfigPath = c.String("config")
		f, err := os.OpenFile(cliconfig.GetConf().Server.LogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			klog.Error(err)
		}
		defer f.Close()
		klog.SetOutput(f)

		agent.HeartBeat()
		agent.ConfigWatch()
		rpc.InitOwlAgentRpc()

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
		addr, err := net.ResolveTCPAddr("tcp", cliconfig.GetConf().Server.RpcHost)
		if err != nil {
			klog.Fatal(err)
		}
		svr := prometheusagentservice.NewServer(new(agent.PrometheusServiceImpl),
			kitexServer.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "owl-prometheus-agent"}),
			kitexServer.WithRegistry(registry.NewNacosRegistry(cli)),
			kitexServer.WithServiceAddr(addr))

		err = svr.Run()
		if err != nil {
			klog.Error(err)
		}
		return nil
	},
}
