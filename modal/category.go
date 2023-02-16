package modal

import (
	"gorm.io/gorm"
	"time"
)

type Category struct {
	Id         int    `gorm:"primary_key"`
	Name       string `gorm:"unique;notnull;size:20"`
	CreateTime time.Time
	DeletedAt  *gorm.DeletedAt `sql:"index" json:"deletedAt"`
}

type CategoryVo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (category Category) ToVo() *CategoryVo {

	if category.Id == 0 {
		return nil
	}

	return &CategoryVo{
		Id:   category.Id,
		Name: category.Name,
	}
}
