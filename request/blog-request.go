package request

import (
	"errors"
	"gin-demo/common"
	"gin-demo/models"
	"gin-demo/myerr"
	"gin-demo/utils"
	"time"
	"unicode/utf8"
)

type BlogRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	Markdown    bool   `json:"markdown"`
	SourceUrl   string `json:"sourceUrl"`
	CoverImage  string `json:"coverImage"`
	Tags        []int  `json:"tags"`
	Topic       int    `json:"topic"`
	Category    int    `json:"category"`
}

func (request BlogRequest) Check() error {
	var titleLen = utf8.RuneCountInString(request.Title)
	var descLen = utf8.RuneCountInString(request.Description)
	if titleLen < 1 || titleLen > 50 {
		return errors.New("博客标题不能小于1个字符并且不能大于50个字符")
	} else if descLen < 1 || descLen > 200 {
		return errors.New("博客简介不能小于1个字符并且不能大于200个字符")
	} else if request.Content == "" {
		return errors.New("博客内容不能为空")
	} else if !utils.IsImageURL(request.CoverImage) {
		return errors.New("这不是一个图片链接")
	}
	return nil
}

func (request BlogRequest) ToDo(uid int) models.Blog {
	var blog models.Blog
	if request.Category > 0 {
		blog.CategoryId = &request.Category
		blog.TopicId = nil
		if len(request.Tags) > 0 {
			for _, tag := range request.Tags {
				blog.Tags = append(blog.Tags, models.Tag{
					Id: tag,
				})
			}
		} else {
			myerr.PanicError(common.RELEASE_TAG_EMPTY)
		}
	} else {
		if request.Topic > 0 {
			blog.TopicId = &request.Topic
			blog.CategoryId = nil
			blog.Tags = nil
		} else {
			myerr.PanicError(common.RELEASE_CATEGORY_OR_TOPIC_EMPTY)
		}
	}
	blog.Description = request.Description
	blog.Title = request.Title
	blog.CoverImage = request.CoverImage
	blog.SourceUrl = request.SourceUrl
	blog.Content = request.Content
	blog.Markdown = request.Markdown
	blog.EyeCount = 0
	blog.LikeCount = 0
	blog.UserId = uid
	var now = time.Now()
	blog.CreateAt = now
	blog.UpdateAt = now
	return blog
}
