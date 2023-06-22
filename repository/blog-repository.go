package repository

import (
	"gin-demo/models"
	"gin-demo/request"
	"gin-demo/vo"
)

type BlogRepository interface {
	PaginatedBlogQueries(request request.BlogPageRequest) (blogs []models.Blog, count int64)
	GetHotBlog() (blogs []vo.SimpleBlogVo)
	FindByIdIn(ids []int) (blogs []vo.SimpleBlogVo)
	FindAllSimpleBlog() (blogs []vo.SimpleBlogVo)
}
