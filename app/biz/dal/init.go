package dal

import (
	"darkhawk/app/biz/dal/mysql"
	"darkhawk/app/biz/dal/redis"
	"darkhawk/app/biz/dal/redisStore"
	"darkhawk/conf"
)

func Init() {
	mysql.Init()
	// Redis连接初始化
	redis.Init(conf.GetConf())

	// init redis store
	redisStore.Init()
}
