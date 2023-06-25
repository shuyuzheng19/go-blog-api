package models

type BlogLike struct {
	Id     int    `gorm:"primary_key"`
	Ip     string `gorm:"not null;unique"`
	BlogId int    `gorm:"column:blog_id"`
}
