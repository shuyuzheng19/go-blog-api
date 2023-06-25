package vo

import (
	"fmt"
	"gin-demo/response"
	"time"
)

type BlogVo struct {
	Id          int          `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"desc"`
	CoverImage  string       `json:"coverImage"`
	DateStr     string       `json:"dateStr"`
	User        SimpleUserVo `json:"user"`
	Category    *CategoryVo  `json:"category,omitempty"`
}

type ArchiveBlogVo struct {
	Id          int          `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"desc"`
	CreateAt    myTimeFormat `json:"create"`
}

type BlogContentVo struct {
	Id          int            `json:"id"`
	Description string         `json:"description"`
	Title       string         `json:"title"`
	CoverImage  string         `json:"coverImage"`
	SourceUrl   string         `json:"sourceUrl"`
	Content     string         `json:"content"`
	EyeCount    int64          `json:"eyeCount"`
	Markdown    bool           `json:"markdown"`
	LikeCount   int64          `json:"likeCount"`
	Category    *CategoryVo    `json:"category"`
	AiMessage   string         `json:"aiMessage,omitempty"`
	Tags        []TagVo        `json:"tags"`
	Topic       *SimpleTopicVo `json:"topic,omitempty"`
	User        SimpleUserVo   `json:"user"`
	CreateTime  string         `json:"createTime"`
	UpdateTime  string         `json:"updateTime"`
}

type myTimeFormat time.Time

func (t *myTimeFormat) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", response.FormatTime(tTime))), nil
}

type SimpleBlogVo struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"desc,omitempty"`
	CoverImage  string `json:"coverImage,omitempty"`
	DateStr     string `json:"create,omitempty"`
}
