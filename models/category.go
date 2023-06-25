package models

import (
	"gin-demo/common"
	"gin-demo/vo"
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

func (category Category) ToVo() *vo.CategoryVo {
	return &vo.CategoryVo{
		Id:   category.Id,
		Name: category.Name,
	}
}
