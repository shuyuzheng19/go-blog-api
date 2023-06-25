package config

import (
	"fmt"
	"gin-demo/models"
	"gin-demo/myerr"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func getDB() *gorm.DB {

	var db_config = CONFIG.Db

	//var url = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=%s", db_config.Username, db_config.Password, db_config.Host, db_config.Port, db_config.Name, db_config.TimeZone)

	var url = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=disable TimeZone=%s", db_config.Host, db_config.Username, db_config.Password, db_config.Name, db_config.Port, db_config.TimeZone)

	var db, err = gorm.Open(postgres.Open(url), &gorm.Config{})

	myerr.MessageError(err, "连接数据库失败")

	if db_config.Auto {
		db.AutoMigrate(&models.User{}, &models.Blog{}, &models.BlogLike{}, &models.Comment{}, &models.FileInfo{}, &models.FileMd5{})
	}

	return db
}

func LoadDbConfig() {
	DB = getDB()
}
