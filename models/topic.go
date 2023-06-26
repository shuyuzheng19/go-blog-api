package models

import (
	"gin-demo/common"
	"gin-demo/vo"
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

func (topic Topic) ToSimpleVo() vo.SimpleTopicVo {
	return vo.SimpleTopicVo{
		Id:   topic.Id,
		Name: topic.Name,
	}
}

func (topic Topic) ToVo() vo.TopicVo {
	return vo.TopicVo{
		Id:          topic.Id,
		Name:        topic.Name,
		Description: topic.Description,
		CoverImage:  topic.Cover,
		User: vo.SimpleUserVo{
			Id:       topic.User.Id,
			Nickname: topic.User.Nickname,
		},
		//CreateAt: response.FormatTimeAgo(topic.CreateAt),
		TimeStamp: topic.CreateAt.UnixMilli(),
	}
}
