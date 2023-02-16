package modal

import (
	"time"
	"vs-blog-api/utils"
)

type Topic struct {
	Id          int    `gorm:"primary_key"`
	Name        string `gorm:"not null;size:20"`
	Description string `gorm:"not null;size:100"`
	CoverImage  string `gorm:"not null;"`
	UserId      int
	User        User `gorm:"foreignKey:UserId;references:Id"`
	CreateTime  time.Time
}

type SimpleTopicVo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type TopicVo struct {
	Id     int          `json:"id"`
	Name   string       `json:"name"`
	Desc   string       `json:"desc"`
	Cover  string       `json:"cover"`
	User   SimpleUserVo `json:"user,omitempty"`
	Create string       `json:"create"`
}

func (t Topic) ToVo() *TopicVo {
	if t.Id == 0 {
		return nil
	}
	return &TopicVo{
		Id:     t.Id,
		Name:   t.Name,
		Desc:   t.Description,
		Cover:  t.CoverImage,
		User:   t.User.ToSimpleVo(),
		Create: utils.DateToString(t.CreateTime),
	}
}
