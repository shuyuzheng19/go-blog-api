package models

import (
	"gin-demo/common"
	"gin-demo/response"
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
	CreateAt    time.Time `gorm:"column:create_at"`
	UpdateAt    time.Time `gorm:"column:update_at"`
	Tags        []Tag     `gorm:"many2many:blogs_tags"`
	CategoryId  *int
	Category    Category
	UserId      int
	User        User
	TopicId     *int
	Topic       Topic
	DeletedAt   *gorm.DeletedAt `gorm:"index"`
}

func (Blog) TableName() string {
	return common.BLOG_TABLE_NAME
}

func (blog Blog) ToContentVo() vo.BlogContentVo {

	var topic *vo.SimpleTopicVo

	var tags []vo.TagVo

	var category *vo.CategoryVo

	if blog.Topic.Id > 0 {
		tags = []vo.TagVo{}
		category = nil
		topic = &vo.SimpleTopicVo{
			Id:   blog.Topic.Id,
			Name: blog.Topic.Name,
		}
	} else {
		topic = nil
		for _, tag := range blog.Tags {
			tags = append(tags, tag.ToVo())
		}
		category = blog.Category.ToVo()
	}

	return vo.BlogContentVo{
		Id:          blog.Id,
		Description: blog.Description,
		Title:       blog.Title,
		CoverImage:  blog.CoverImage,
		SourceUrl:   blog.SourceUrl,
		Content:     blog.Content,
		Markdown:    blog.Markdown,
		EyeCount:    blog.EyeCount,
		LikeCount:   blog.LikeCount,
		Category:    category,
		AiMessage:   "",
		Tags:        tags,
		Topic:       topic,
		User: vo.SimpleUserVo{
			Id:       blog.User.Id,
			Nickname: blog.User.Nickname,
		},
		CreateTime: response.FormatTime(blog.CreateAt),
		UpdateTime: response.FormatTime(blog.UpdateAt),
	}
}

func (blog Blog) ToVo() vo.BlogVo {
	return vo.BlogVo{
		Id:          blog.Id,
		Title:       blog.Title,
		Description: blog.Description,
		CoverImage:  blog.CoverImage,
		TimeStamp:   blog.CreateAt.UnixMilli(),
		//DateStr:     response.FormatTimeAgo(blog.CreateAt),
		User: vo.SimpleUserVo{
			Id:       blog.User.Id,
			Nickname: blog.User.Nickname,
		},
		Category: &vo.CategoryVo{
			Id:   blog.Category.Id,
			Name: blog.Category.Name,
		},
	}
}

func (blog Blog) ToTopicVo() vo.BlogVo {
	return vo.BlogVo{
		Id:          blog.Id,
		Title:       blog.Title,
		Description: blog.Description,
		CoverImage:  blog.CoverImage,
		//DateStr:     response.FormatTimeAgo(blog.CreateAt),
		TimeStamp: blog.CreateAt.UnixMilli(),
		User: vo.SimpleUserVo{
			Id:       blog.User.Id,
			Nickname: blog.User.Nickname,
		},
		Category: nil,
	}
}
