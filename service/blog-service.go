package service

import (
	"gin-demo/models"
	"gin-demo/request"
	"gin-demo/response"
	"gin-demo/vo"
	"io"
)

type BlogService interface {
	FindBlogByCidPage(pageRequest request.BlogPageRequest) response.PageInfo
	GetHotBlog() []vo.SimpleBlogVo
	SetRecommend(ids []int)
	GetRecommend() (blogs []vo.SimpleBlogVo)
	RandomBlogs() (blogs []vo.SimpleBlogVo)
	SearchBlog(keyword string) interface{}
	CountSearch(keyword string, page int) response.PageInfo
	GetHotKeywords() []string
	GetRangeBlog(page int, startStamp int64, endStamp int64) response.PageInfo
	GetBlogByUser(uid int, page int) response.PageInfo
	GetUserBlogTop(uid int) []vo.SimpleBlogVo
	GetBlogById(id int) (blogVo vo.BlogContentVo)
	CurrentIpIsLikeBlog(ip string, id int) bool
	AddLikeBlog(ip string, id int) int64
	Chat(token string, message string) io.ReadCloser
	SaveUserEditorContent(uid int, content string)
	GetUserEditorContent(uid int) string
	SaveBlog(uid int, blogRequest request.BlogRequest)
	AddBlogToSearch(blog []vo.SimpleBlogVo)
	UpdateBlog(id int, user models.User, blogRequest request.BlogRequest)
	InitBlogEyeCount()
	InitBlogLikeCount()
	InitSearch()
}
