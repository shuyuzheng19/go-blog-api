package service

import (
	"vs-blog-api/dto"
	"vs-blog-api/modal"
	"vs-blog-api/response"
)

type BlogService interface {
	FindBlogs(dto dto.BlogPageSortDto) response.PageInfoResponse
	GetHotBlogs() []modal.SimpleBlog
	GetRandomBlogs() []modal.SimpleBlog
	FindByIdBlog(id int) modal.Blog
	RangeDateBlog(start int64, end int64, page int) []modal.RangeBlog
	SearchBlog(keyword string, page int) response.PageInfoResponse
	FindUserBlog(id int, page int) response.PageInfoResponse
	GetUserTopBlog(id int) []modal.SimpleBlog
	GetSimilarBlog(keyword string, blogId int) []modal.SimpleBlog
	SaveBlog(userId int, content string)
	GetSaveBlog(userId int) string
	AddBlog(userId int, blog modal.BlogRequest)
	GetUserLikeBlog(userId int, page int) response.PageInfoResponse
}
