package service

import (
	"gin-demo/request"
	"gin-demo/response"
	"gin-demo/vo"
)

type BlogService interface {
	FindBlogByCidPage(pageRequest request.BlogPageRequest) response.PageInfo
	GetHotBlog() []vo.SimpleBlogVo
	SetRecommend(ids []int)
	GetRecommend() (blogs []vo.SimpleBlogVo)
	RandomBlogs() (blogs []vo.SimpleBlogVo)
}
