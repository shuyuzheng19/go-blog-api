package config

import (
	"gin-demo/myerr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/yaml.v3"
	"os"
)

var CONFIG config

const config_path = "application.yml"

var LOGGER *zap.Logger

type config struct {
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`
	Db struct {
		Init     bool   `yaml:"init"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
		TimeZone string `yaml:"timeZone"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
	} `yaml:"db"`
	Redis struct {
		Addr     string `yaml:"addr"`
		DB       int    `yaml:"db"`
		PoolSize int    `yaml:"poolSize"`
		Password string `yaml:"password"`
		Timeout  int    `yaml:"timeout"`
	} `yaml:"redis"`
	Email struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Addr     string `yaml:"addr"`
	} `yaml:"email"`
}

func LoadConfig() {
	var file, err = os.ReadFile(config_path)

	myerr.MessageError(err, "文件打开失败!")

	yaml.Unmarshal(file, &CONFIG)
}

func LoadLogger() *zap.Logger {

	config := zap.NewProductionConfig()

	config.Encoding = "console"

	config.OutputPaths = []string{"stdout"}

	config.ErrorOutputPaths = []string{"stderr"}

	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	var logger, err = config.Build()

	myerr.MessageError(err, "日志初始化失败")

	return logger
}
