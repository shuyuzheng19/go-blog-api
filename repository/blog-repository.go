package repository

import (
	"gin-demo/models"
	"gin-demo/request"
	"gin-demo/vo"
	"time"
)

type BlogRepository interface {
	PaginatedBlogQueries(request request.BlogPageRequest) (blogs []models.Blog, count int64)
	GetHotBlog() (blogs []vo.SimpleBlogVo)
	FindByIdIn(ids []int) (blogs []vo.SimpleBlogVo)
	FindAllSimpleBlog() (blogs []vo.SimpleBlogVo)
	FindAllSimpleSearchBlog() (blogs []vo.SimpleBlogVo)
	FindRangeBlog(page int, start time.Time, end time.Time) (blogs []vo.ArchiveBlogVo, count int64)
	FindBlogByUserId(uid int, page int) (blogs []models.Blog, count int64)
	FindBlogByUserTop(uid int) (blogs []vo.SimpleBlogVo)
	FindBlogById(id int) (blog models.Blog)
	SaveLikeBlog(like models.BlogLike) error
	CurrentIpIsLikeBlog(ip string, id int) (count int64)
	AddBlogToDb(blog models.Blog) models.Blog
	UpdateBlog(blog models.Blog) error
	UpdateBlogEyeCount(bid int, count int64) error
	UpdateBlogLikeCount(bid int, count int64) error
}
