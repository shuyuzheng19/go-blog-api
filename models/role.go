package models

import "gin-demo/common"

type Role struct {
	Id          int    `gorm:"primary_key"`
	Name        string `gorm:"unique;not null;size:20"`
	Description string `gorm:"not null"`
}

func (Role) TableName() string {
	return common.ROLE_TABLE_NAME
}
