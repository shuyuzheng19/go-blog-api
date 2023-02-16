package repository

import (
	"vs-blog-api/dto"
	"vs-blog-api/modal"
)

type BlogRepository interface {
	FindAll(pageRequest dto.BlogPageSortDto) (blogs []modal.Blog, err error, count int64)
	FindById(id int) (err error, blog modal.Blog)
	FindAllIdAndTitle() (err error, blogs []modal.SimpleBlog)
	FindRangeDate(startTime int64, endTime int64, page int, sortField string) (err error, blogs []modal.RangeBlog)
	FindBlogByUserId(id int, page int, sortField string) (err error, blogs []modal.Blog, count int64)
	FindBlogIdIn(ids []string) (err error, blogs []modal.Blog)
	SaveBlog(blog modal.Blog) (error, modal.Blog)
}
