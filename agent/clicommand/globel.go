package clicommand

import (
	"fmt"
	"github.com/oleiade/reflections"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func CreateLogger(cfg *viper.Viper) *zap.Logger {
	zapCfg := zap.NewProductionConfig()
	err := handleLogLevelFlag(&zapCfg, cfg)
	if err != nil {
		fmt.Printf("Error when setting log level: %v. Defaulting log level to NOTICE", err)
		os.Exit(1)
	}

	debugI, _ := reflections.GetField(cfg, "Debug")
	if debug, ok := debugI.(bool); ok && debug {
		zapCfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}
	zapCfg.Encoding = "json"
	zapCfg.OutputPaths = []string{"stdout", "/tmp/logs"}
	zapCfg.ErrorOutputPaths = []string{"stderr"}
	zapCfg.InitialFields = map[string]interface{}{"foo": "bar"}
	zapCfg.EncoderConfig = zapcore.EncoderConfig{
		MessageKey: "message",
		LevelKey:   "level",
	}
	logger := zap.Must(zapCfg.Build())
	defer logger.Sync()
	return logger
}

func handleLogLevelFlag(zapCfg *zap.Config, cfg *viper.Viper) error {
	logLevel := cfg.Get("server.logLevel")
	llStr, ok := logLevel.(string)
	if !ok {
		return fmt.Errorf("log level %v (%T) couldn't be cast to string", logLevel, logLevel)
	}

	level, err := zap.ParseAtomicLevel(llStr)
	if err != nil {
		return err
	}

	zapCfg.Level = level
	return nil
}
