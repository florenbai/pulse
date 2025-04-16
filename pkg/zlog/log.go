package zlog

import (
	"darkhawk/conf"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var log *zap.Logger

func init() {
	// log init
	_, err := os.Stat(conf.GetConf().Server.LogPath)
	if err != nil {
		if os.IsNotExist(err) {
			_ = os.Mkdir(conf.GetConf().Server.LogPath, os.ModePerm)
		} else {
			panic(err)
		}
	}
	debugPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.DebugLevel
	})
	// 日志级别
	infoPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.InfoLevel
	})
	// 错误级别
	errorPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.ErrorLevel
	})

	AccessLogLumberJackLogger := &lumberjack.Logger{
		Filename:   conf.GetConf().Server.AccessLog,
		MaxSize:    10,
		MaxBackups: 20,
		MaxAge:     30,
		Compress:   false,
	}

	errorLumberJackLogger := &lumberjack.Logger{
		Filename:   conf.GetConf().Server.ErrorLog,
		MaxSize:    10,
		MaxBackups: 20,
		MaxAge:     30,
		Compress:   false,
	}

	debugLumberJackLogger := &lumberjack.Logger{
		Filename:   conf.GetConf().Server.DebugLog,
		MaxSize:    10,
		MaxBackups: 20,
		MaxAge:     30,
		Compress:   false,
	}

	eConfig := zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.RFC3339TimeEncoder,     //自定义时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //编码间隔 s
		EncodeCaller:   zapcore.FullCallerEncoder,      //控制打印的文件位置是绝对路径,ShortCallerEncoder 是相对路径
	}

	acWs := zapcore.AddSync(AccessLogLumberJackLogger)
	errWs := zapcore.AddSync(errorLumberJackLogger)
	debugWs := zapcore.AddSync(debugLumberJackLogger)

	encoder := zapcore.NewJSONEncoder(eConfig)

	cores := []zapcore.Core{
		zapcore.NewCore(encoder, acWs, infoPriority),
		zapcore.NewCore(encoder, errWs, errorPriority),
		zapcore.NewCore(encoder, debugWs, debugPriority),
	}

	log = zap.New(zapcore.NewTee(cores...), zap.AddCaller(), zap.AddCallerSkip(1))
}

func AccessLog(msg string, fields ...zapcore.Field) {
	log.Info(msg, fields...)
}

func Error(msg string, fields ...zapcore.Field) {
	log.Error(msg, fields...)
}

func Debug(msg string, fields ...zapcore.Field) {
	log.Debug(msg, fields...)
}
