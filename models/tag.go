package models

import (
	"gin-demo/common"
	"gin-demo/vo"
	"gorm.io/gorm"
	"time"
)

type Tag struct {
	Id        int             `gorm:"primary_key"`
	Name      string          `gorm:"column:name;not null;unique"`
	CreateAt  time.Time       `gorm:"column:create_at"`
	DeletedAt *gorm.DeletedAt `gorm:"index"`
}

func (Tag) TableName() string {
	return common.TAG_TABLE_NAME
}

func (tag Tag) ToVo() vo.TagVo {
	return vo.TagVo{
		Id:   tag.Id,
		Name: tag.Name,
	}
}
