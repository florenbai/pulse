package conf

import (
	"github.com/bytedance/go-tagexpr/v2/validator"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/klog"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"sync"
)

var (
	conf *Config
	once sync.Once
)

type Config struct {
	Server         Server         `yaml:"server"`         //服务器配置
	Authentication Authentication `yaml:"authentication"` //登录验证配置
	MySQL          MySQL          `yaml:"mysql"`          //MySQL配置
	Redis          Redis          `yaml:"redis"`          //Redis配置
	Umdp           Umdp           `yaml:"umdp"`           //umdp配置
	Nacos          NacosConfig    `yaml:"nacos"`
}

type Server struct {
	LogLevel         string `yaml:"log_level"`
	LogPath          string `yaml:"log_path"`
	ErrorLog         string `yaml:"error_log"`
	DebugLog         string `yaml:"debug_log"`
	AccessLog        string `yaml:"access_log"`
	Addr             string `yaml:"addr"`
	GrpcAddr         string `yaml:"grpc_addr"`
	RpcName          string `yaml:"rpc_name"`
	RemotePrometheus bool   `yaml:"remote_prometheus"`
}

type Authentication struct {
	AuthSecret string `yaml:"auth_secret"`
}

type MySQL struct {
	DSN string `yaml:"dsn"` // MySQL dsn配置
}

type Redis struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
	MaxIdle  int    `yaml:"max_idle"`
}

type WechatBotKey struct {
	Name string `yaml:"name" json:"name"`
	Key  string `yaml:"key" json:"key"`
}

type Umdp struct {
	Host             string         `yaml:"host"`
	Token            string         `yaml:"token"`
	EmailChannel     string         `yaml:"email_channel"`
	WechatChannel    string         `yaml:"wechat_channel"`
	WechatBotChannel string         `yaml:"wechat_bot_channel"`
	WechatBotKeys    []WechatBotKey `yaml:"wechat_bot_keys"`
	PhoneChannel     string         `yaml:"phone_channel"`
	Template         int            `yaml:"template"`
}

type NacosConfig struct {
	Namespace string `yaml:"namespace"`
	Addr      string `yaml:"addr"`
	Port      uint64 `yaml:"port"`
	LogDir    string `yaml:"log_dir"`
	CacheDir  string `yaml:"cache_dir"`
	LogLevel  string `yaml:"log_level"`
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
}

// GetConf gets configuration instance
func GetConf() *Config {
	once.Do(initConf)
	return conf
}

func initConf() {
	prefix := "conf"
	confFileRelPath := filepath.Join(prefix, "config.yaml")
	content, err := os.ReadFile(confFileRelPath)
	if err != nil {
		panic(err)
	}

	conf = new(Config)
	err = yaml.Unmarshal(content, conf)
	if err != nil {
		hlog.Error("parse yaml error - %v", err)
		panic(err)
	}
	if err := validator.Validate(conf); err != nil {
		hlog.Error("validate config error - %v", err)
		panic(err)
	}
}

func HLogLevel() hlog.Level {
	level := GetConf().Server.LogLevel
	switch level {
	case "trace":
		return hlog.LevelTrace
	case "debug":
		return hlog.LevelDebug
	case "info":
		return hlog.LevelInfo
	case "notice":
		return hlog.LevelNotice
	case "warn":
		return hlog.LevelWarn
	case "error":
		return hlog.LevelError
	case "fatal":
		return hlog.LevelFatal
	default:
		return hlog.LevelInfo
	}
}

func KLogLevel() klog.Level {
	level := GetConf().Server.LogLevel
	switch level {
	case "trace":
		return klog.LevelTrace
	case "debug":
		return klog.LevelDebug
	case "info":
		return klog.LevelInfo
	case "notice":
		return klog.LevelNotice
	case "warn":
		return klog.LevelWarn
	case "error":
		return klog.LevelError
	case "fatal":
		return klog.LevelFatal
	default:
		return klog.LevelInfo
	}
}
