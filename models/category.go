package models

import (
	"gin-demo/common"
	"gorm.io/gorm"
	"time"
)

type Category struct {
	Id        int             `gorm:"primary_key"`
	Name      string          `gorm:"column:name;not null;unique"`
	CreateAt  time.Time       `gorm:"column:create_at"`
	DeletedAt *gorm.DeletedAt `gorm:"index"`
}

func (Category) TableName() string {
	return common.CATEGORY_TABLE_NAME
}
