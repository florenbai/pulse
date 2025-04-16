package cliconfig

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
	"go.uber.org/zap"
	"os"
	"regexp"
)

type Loader struct {
	// 命令行上下文
	CLI *cli.Context

	// 配置文件
	Config *viper.Viper

	// 日志
	Logger zap.Logger

	// 默认配置文件
	DefaultConfigFilePaths []string
}

var argCliNameRegexp = regexp.MustCompile(`arg:(\d+)`)

func (l *Loader) Load() error {
	configPath := l.CLI.String("config")
	if configPath != "" {
		_, err := os.Stat(configPath)
		if err != nil {
			return fmt.Errorf("a configuration file could not be found at: %q", configPath)
		}

	} else if len(l.DefaultConfigFilePaths) > 0 {
		for _, path := range l.DefaultConfigFilePaths {
			_, err := os.Stat(path)
			if err == nil {
				configPath = path
				break
			}
		}
	}
	configPath = "./config.yaml"

	vip := viper.New()
	vip.SetConfigType("yaml")
	vip.SetConfigFile(configPath)
	err := vip.ReadInConfig()
	if err != nil {
		return fmt.Errorf("Fatal error config file: %s \n", err)
	}

	l.Config = vip

	return nil
}
