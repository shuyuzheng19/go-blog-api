package models

import (
	"fmt"
	"gin-demo/common"
	"gin-demo/vo"
	"gorm.io/gorm"
	"time"
)

type Blog struct {
	Id          int       `gorm:"primary_key"`
	Description string    `gorm:"column:description;not null"`
	Title       string    `gorm:"column:title;not null"`
	CoverImage  string    `gorm:"column:cover_image;not null"`
	SourceUrl   string    `gorm:"column:source_url"`
	Content     string    `gorm:"column:content;type:text"`
	EyeCount    int64     `gorm:"column:eye_count;default:0"`
	LikeCount   int64     `gorm:"column:like_count;default:0"`
	Markdown    bool      `gorm:"column:markdown;default:true"`
	CreateAt    time.Time `gorm:"create_at"`
	UpdateAt    time.Time `gorm:"update_at"`
	Tags        []Tag     `gorm:"many2many:blogs_tags"`
	CategoryId  int
	Category    Category
	UserId      int
	User        User
	TopicId     int
	Topic       Topic
	DeletedAt   *gorm.DeletedAt `gorm:"index"`
}

func (Blog) TableName() string {
	return common.BLOG_TABLE_NAME
}

func (blog Blog) ToVo() vo.BlogVo {
	return vo.BlogVo{
		Id:          blog.Id,
		Title:       blog.Title,
		Description: blog.Description,
		CoverImage:  blog.CoverImage,
		DateStr:     formatTimeAgo(blog.CreateAt.Unix()),
		User: vo.SimpleUserVo{
			Id:       blog.User.Id,
			Nickname: blog.User.Nickname,
		},
		Category: vo.CategoryVo{
			Id:   blog.Category.Id,
			Name: blog.Category.Name,
		},
	}
}
func formatTimeAgo(unix int64) string {

	second := time.Now().Unix() - unix

	var dateStr string

	if second <= 60 {
		dateStr = "刚刚"
	} else if second > 60 && second <= 60*60 {
		dateStr = fmt.Sprintf("%d分钟前", second/60)
	} else if second > 60*60 && second <= 60*60*24 {
		dateStr = fmt.Sprintf("%d小时前", second/60/60)
	} else if second > 60*60*24 && second <= 60*60*24*30 {
		dateStr = fmt.Sprintf("%d天前", second/60/60/24)
	} else if second > 60*60*24*30 && second <= 60*60*24*30*12 {
		dateStr = fmt.Sprintf("%d月前", second/60/60/24/30)
	} else {
		dateStr = fmt.Sprintf("%d年前", second/60/60/24/(30*12))
	}

	return dateStr
}
