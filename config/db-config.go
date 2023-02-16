package config

import (
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"vs-blog-api/modal"
)

var DB *gorm.DB

// 连接到数据库
func getDB() *gorm.DB {

	var DBConfig = LoadDBConfig()

	//var url = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=%s",DBConfig.Username,DBConfig.Password,DBConfig.Host,DBConfig.Port,DBConfig.Name,DBConfig.TimeZone)

	var url = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=disable TimeZone=Asia/Shanghai", DBConfig.Host, DBConfig.Username, DBConfig.Password, DBConfig.Name, DBConfig.Port)

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		panic(errors.New("链接数据库失败"))
		return nil
	}

	return db
}

func InitDb() {
	DB = getDB()
	DB.AutoMigrate(&modal.User{}, &modal.Blog{}, &modal.Topic{}, &modal.UserLike{}, &modal.TimeLine{})
}
