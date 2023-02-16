package config

import (
	"gopkg.in/ini.v1"
	"strconv"
	"vs-blog-api/common"
)

var CONFIG *ini.File

var ServerPort string

var HostName string

func InitConfig() {

	var config, err = ini.Load(common.SERVER_CONFIG_PATH)

	if err != nil {
		panic(err)
	}

	CONFIG = config

	section := CONFIG.Section("")

	ServerPort = section.Key("port").String()

	HostName = section.Key("hostname").String()
}

func LoadEmailConfig() EmailConfig {

	var emailConfig = CONFIG.Section("email")

	return EmailConfig{
		Username: emailConfig.Key("username").String(),
		Password: emailConfig.Key("password").String(),
		Host:     emailConfig.Key("host").String(),
		Addr:     emailConfig.Key("addr").String(),
	}
	return EmailConfig{}
}

type RedisConfig struct {
	Addr     string
	Password string
	Db       int
}

func LoadRedisConfig() RedisConfig {

	var redisConfig = CONFIG.Section("redis")

	db, err := redisConfig.Key("db").Int()

	port, err2 := redisConfig.Key("port").Int()

	if err != nil {
		db = 0
	}

	if err2 != nil {
		port = 6379
	}

	return RedisConfig{
		Addr:     HostName + ":" + strconv.Itoa(port),
		Password: redisConfig.Key("password").String(),
		Db:       db,
	}
}

func LoadDBConfig() myDBConfig {

	var dbConfig = CONFIG.Section("db")

	var port, err = dbConfig.Key("port").Int()

	if err != nil {
		port = 3306
	}

	return myDBConfig{
		Username:     dbConfig.Key("username").String(),
		Password:     dbConfig.Key("password").String(),
		Host:         HostName,
		Port:         port,
		Name:         dbConfig.Key("name").String(),
		CharEncoding: dbConfig.Key("charEncoding").String(),
		TimeZone:     dbConfig.Key("timeZone").String(),
	}
}

type myDBConfig struct {
	Username     string
	Password     string
	Host         string
	Port         int
	Name         string
	CharEncoding string
	TimeZone     string
}
