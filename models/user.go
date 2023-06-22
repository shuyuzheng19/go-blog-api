package models

import (
	"gin-demo/common"
	"gin-demo/vo"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        int             `gorm:"primary_key"`
	Username  string          `gorm:"size:16;not null;unique"`
	Nickname  string          `gorm:"column:nick_name;not null"`
	Password  string          `gorm:"not null"`
	Email     string          `gorm:"not null;unique"`
	Icon      string          `json:"icon"`
	DeletedAt *gorm.DeletedAt `gorm:"index"`
	CreateAt  time.Time       `gorm:"column:create_at"`
	UpdateAt  time.Time       `gorm:"column:update_at"`
	RoleId    int
	Role      Role
}

func (User) TableName() string {
	return common.USER_TABLE_NAME
}

func (user User) ToUserVo() vo.UserVo {
	return vo.UserVo{
		Id:       user.Id,
		Username: user.Username,
		Nickname: user.Nickname,
		Role: vo.RoleVo{
			Id:   user.Role.Id,
			Name: user.Role.Name,
		},
	}
}
