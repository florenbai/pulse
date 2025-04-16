package conf

import "time"

const (
	// MinGoVersion 最小 Go 版本
	MinGoVersion = 1.20

	// AppVersion 项目版本
	AppVersion = "v1.20"

	// AppName 项目名称
	AppName = "darkhawk"

	// Dom 租户/域
	Dom = "darkhawk"

	// AppDomain 应用域名
	AppDomain = "http://darkhawk.maas.local"

	// AppPort 项目端口
	AppPort = 80

	// KeyPairs
	KeyPairs = "minos_key_2023"

	RequestIDHeaderValue = "X-Request-ID"
	SchemeHeaderValue    = "X-Darkhawk-Scheme"

	// AppAccessLogFile 项目访问日志存放文件
	AppAccessLogFile = "./log/" + AppName + "-access.log"

	// AppCronLogFile 项目后台任务日志存放文件
	AppCronLogFile = "./log/" + AppName + "-cron.log"

	// AppInstallMark 项目安装完成标识
	AppInstallMark = "INSTALL.lock"

	// HeaderLoginToken 登录验证 Token，Header 中传递的参数
	HeaderLoginToken = "Token"

	// HeaderSignToken 签名验证 Authorization，Header 中传递的参数
	HeaderSignToken = "Authorization"

	// HeaderSignTokenDate 签名验证 Date，Header 中传递的参数
	HeaderSignTokenDate = "Authorization-Date"

	// HeaderSignTokenTimeout 签名有效期为 2 分钟
	HeaderSignTokenTimeout = time.Minute * 2

	// RedisKeyPrefixLoginUser Redis Key 前缀 - 登录用户信息
	RedisKeyPrefixLoginUser = AppName + ":login-user:"

	// RedisKeyPrefixSignature Redis Key 前缀 - 签名验证信息
	RedisKeyPrefixSignature = AppName + ":signature:"

	// ZhCN 简体中文 - 中国
	ZhCN = "zh-cn"

	// EnUS 英文 - 美国
	EnUS = "en-us"

	// MaxRequestsPerSecond 每秒最大请求量
	MaxRequestsPerSecond = 10000

	// LoginSessionTTL 登录有效期为 24 小时
	LoginSessionTTL = time.Hour * 24
)
