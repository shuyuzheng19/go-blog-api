package modal

import (
	"gorm.io/gorm"
	"time"
)

type Tag struct {
	Id         int    `gorm:"primary_key"`
	Name       string `gorm:"unique;notnull;size:20"`
	CreateTime time.Time
	DeletedAt  *gorm.DeletedAt `sql:"index" json:"deletedAt"`
}

type TagVo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (tag Tag) ToVo() TagVo {
	return TagVo{
		Id:   tag.Id,
		Name: tag.Name,
	}
}
