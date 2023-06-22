package models

import (
	"gin-demo/common"
	"gorm.io/gorm"
	"time"
)

type Topic struct {
	Id          int             `gorm:"primary_key"`
	Name        string          `gorm:"column:name;not null;unique"`
	Description string          `gorm:"column:description;not null"`
	Cover       string          `gorm:"column:cover_image;not null"`
	CreateAt    time.Time       `gorm:"column:create_at"`
	DeletedAt   *gorm.DeletedAt `gorm:"index"`
	UserId      int
	User        User
}

func (Topic) TableName() string {
	return common.TOPIC_TABLE_NAME
}
