package mysql

import (
	"darkhawk/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/opentelemetry/tracing"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func Init() {
	var err error
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,  // 慢 SQL 阈值
			LogLevel:      logger.Error, // TaskLog level
			Colorful:      false,        // 禁用彩色打印
		},
	)
	var gconf = gorm.Config{
		Logger:                                   newLogger,
		SkipDefaultTransaction:                   true,
		PrepareStmt:                              true,
		QueryFields:                              true,
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
	}
	DB, err = gorm.Open(mysql.Open(conf.GetConf().MySQL.DSN), &gconf)
	if err != nil {
		panic("failed to connect database")
	}
	if err := DB.Use(tracing.NewPlugin()); err != nil {
		panic(err)
	}
}
