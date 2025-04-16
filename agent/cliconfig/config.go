package cliconfig

import (
	"github.com/bytedance/go-tagexpr/v2/validator"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path/filepath"
	"sync"
)

var (
	conf *Config
	once sync.Once
)
var ConfigPath string

type Config struct {
	Server     Server     `json:"server" yaml:"server"`
	Nacos      Nacos      `json:"nacos" yaml:"nacos"`
	Prometheus Prometheus `json:"prometheus" yaml:"prometheus"`
}

type Server struct {
	LogPath  string `json:"logPath" yaml:"logPath"`
	LogLevel string `json:"logLevel" yaml:"logLevel"`
	RpcHost  string `json:"rpcHost" yaml:"rpcHost"`
}

type Nacos struct {
	Address     string `json:"address" yaml:"address"`
	Port        uint64 `json:"port" yaml:"port"`
	NamespaceId string `json:"namespaceId" yaml:"namespaceId"`
	Username    string `json:"username" yaml:"username"`
	Password    string `json:"password" yaml:"password"`
}

type Prometheus struct {
	SourceSign string `json:"sourceSign" yaml:"sourceSign"`
	Endpoint   string `json:"endpoint" yaml:"endpoint"`
}

func GetConf() *Config {
	once.Do(initConf)
	return conf
}

func initConf() {
	var confFileRelPath string
	if ConfigPath != "" {
		confFileRelPath = ConfigPath
	} else {
		prefix := ""
		confFileRelPath = filepath.Join(prefix, "config.yaml")
	}
	content, err := os.ReadFile(confFileRelPath)
	if err != nil {
		log.Fatal(err)
	}

	conf = new(Config)
	err = yaml.Unmarshal(content, conf)
	if err != nil {
		log.Fatalf("parse yaml error - %v", err)
	}
	if err := validator.Validate(conf); err != nil {
		log.Fatalf("validate config error - %v", err)
	}
}
