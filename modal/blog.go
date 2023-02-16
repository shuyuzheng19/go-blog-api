package modal

import (
	"fmt"
	"gorm.io/gorm"
	"time"
	"vs-blog-api/response"
	"vs-blog-api/utils"
)

type Blog struct {
	Id          int    `gorm:"primary_key"`
	Description string `gorm:"not null;size:100"`
	Title       string `gorm:"not null;size:50"`
	Content     string `gorm:"type:text"`
	SourceUrl   string
	CoverImage  string          `gorm:"not null;"`
	EyeCount    int64           `gorm:"not null;default:0"`
	LikeCount   int64           `gorm:"not null;default:0"`
	Tags        []Tag           `gorm:"many2many:blogs_tags"`
	CategoryId  int             `gorm:"default:null;"`
	Category    Category        `gorm:"foreignKey:CategoryId;references:Id"`
	TopicId     int             `gorm:"type:uint;default:null;"`
	Topic       Topic           `gorm:"foreignKey:TopicId;references:Id"`
	UserId      int             `gorm:"type:uint"`
	User        User            `gorm:"foreignKey:UserId;references:Id"`
	DeletedAt   *gorm.DeletedAt `sql:"index"`
	Markdown    bool            `gorm:'default:true'`
	CreateTime  time.Time       `gorm:"notnull"`
	UpdateTime  time.Time       `gorm:"notnull"`
}

type BlogRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	SourceUrl   string `json:"source_url"`
	CoverImage  string `json:"coverImage"`
	Tags        []int  `json:"tags"`
	Topic       int    `json:"topic"`
	Markdown    string `json:"markdown"`
	Category    int    `json:"category"`
}

func (blog BlogRequest) Validator() {
	if utils.AnyEmpty(blog.Content, blog.Title, blog.Description) {
		panic(response.NewGlobalException(response.ParamsError, "参数有空"))
	}

	if len(blog.Title) > 50 || len(blog.Description) > 100 {
		panic(response.NewGlobalException(response.ParamsError, "博客标题字符大于50获取博客描述字符大于100"))
	}

	if blog.Topic == 0 {
		if blog.Category <= 0 {
			panic(response.NewGlobalException(response.ParamsError, "请选择分类"))
		}
		if len(blog.Tags) == 0 {
			panic(response.NewGlobalException(response.ParamsError, "请选择标签"))
		}
	}
}

type BlogVo struct {
	Id         int          `json:"id"`
	Desc       string       `json:"desc"`
	Title      string       `json:"title"`
	CoverImage string       `json:"coverImage"`
	Category   *CategoryVo  `json:"category,omitempty"`
	User       SimpleUserVo `json:"user,omitempty"`
	DateStr    string       `json:"dateStr"`
}

type BlogContentVo struct {
	Blog       BlogVo  `json:"blog"`
	CreateTime string  `json:"createTime"`
	UpdateTime string  `json:"updateTime"`
	EyeCount   int64   `json:"eyeCount"`
	LikeCount  int64   `json:"likeCount"`
	Content    string  `json:"content"`
	Tags       []TagVo `json:"tags,omitempty"`
	SourceUrl  string  `json:"sourceUrl"`
	Topic      int     `json:"topic"`
	Markdown   bool    `json:"markdown"`
}

type SimpleBlog struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type LocalTime time.Time

func (t *LocalTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format(utils.FORMAT_DATE_TIME))), nil
}

type RangeBlog struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	CreateTime  LocalTime `json:"createTime"`
	Description string    `json:"description"`
}

type EsBlog struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (blog Blog) ToContentBlogVo() BlogContentVo {

	var tags []TagVo

	for _, tag := range blog.Tags {
		tags = append(tags, tag.ToVo())
	}

	return BlogContentVo{
		Blog:       blog.ToBlogVo(),
		CreateTime: utils.FormatDate(blog.CreateTime, utils.FORMAT_DATE_SIMPLE_TIME),
		UpdateTime: utils.FormatDate(blog.CreateTime, utils.FORMAT_DATE_TIME),
		EyeCount:   blog.EyeCount,
		Topic:      blog.TopicId,
		LikeCount:  blog.LikeCount,
		Content:    blog.Content,
		Markdown:   blog.Markdown,
		Tags:       tags,
		SourceUrl:  blog.SourceUrl,
	}
}

type RecommendVo struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	CoverImage string `json:"cover"`
}

func (blog Blog) ToBlogVo() BlogVo {

	if blog.Id == 0 {
		return BlogVo{}
	}

	return BlogVo{
		Id:         blog.Id,
		Desc:       blog.Description,
		Title:      blog.Title,
		CoverImage: blog.CoverImage,
		Category:   blog.Category.ToVo(),
		User:       blog.User.ToSimpleVo(),
		DateStr:    utils.DateToString(blog.CreateTime),
	}
}
