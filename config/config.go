package config

import (
	"gin-demo/common"
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
	InitSearch bool `yaml:"initSearch"`
	Server     struct {
		Port int `yaml:"port"`
	} `yaml:"server"`
	Db struct {
		Auto     bool   `yaml:"auto"`
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
	MeiliSearchConfig struct {
		ApiHost string `yaml:"apiHost"`
		ApiKey  string `yaml:"apiKey"`
	} `yaml:"meilisearch"`
	Upload struct {
		MaxImageSize int64  `yaml:"maxImageSize"`
		MaxFileSize  int64  `yaml:"maxFileSize"`
		Prefix       string `yaml:"prefix"`
		Path         string `yaml:"path"`
		Uri          string `yaml:"uri"`
	} `yaml:"upload"`
}

type UploadConfig struct {
	MaxImageSize int64  `yaml:"maxImageSize"`
	MaxFileSize  int64  `yaml:"maxFileSize"`
	Prefix       string `yaml:"prefix"`
	Path         string `yaml:"path"`
	Uri          string `yaml:"uri"`
}

func GetUploadConfig() UploadConfig {

	var upload = CONFIG.Upload

	os.MkdirAll(upload.Path+"/"+common.AVATAR, os.ModePerm)

	os.MkdirAll(upload.Path+"/"+common.IMAGES, os.ModePerm)

	os.MkdirAll(upload.Path+"/"+common.FILES, os.ModePerm)

	return UploadConfig{
		MaxImageSize: upload.MaxImageSize * 1024 * 1024,
		MaxFileSize:  upload.MaxFileSize * 1024 * 1024,
		Prefix:       upload.Prefix,
		Path:         upload.Path,
		Uri:          upload.Uri,
	}
}

func LoadConfig() {
	var file, err = os.ReadFile(config_path)

	myerr.MessageError(err, "文件打开失败!")

	yaml.Unmarshal(file, &CONFIG)
}

func LoadLogger() *zap.Logger {

	config := zap.NewProductionConfig()

	config.Encoding = "console"

	config.OutputPaths = []string{"info.log"}

	config.ErrorOutputPaths = []string{"error.log"}

	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	var logger, err = config.Build()

	myerr.MessageError(err, "日志初始化失败")

	return logger

}
