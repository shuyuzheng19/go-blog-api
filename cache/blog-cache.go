package cache

import (
	"gin-demo/response"
	"gin-demo/vo"
)

type BlogCache interface {
	//获取热门博客
	GetTopBlog() (blogs []vo.SimpleBlogVo)
	//存取热门博客
	SaveTopBlog(blogs []vo.SimpleBlogVo) error
	//获取推荐文章
	GetRecommendBlog() (blogs []vo.SimpleBlogVo)
	//存取推荐文章
	SetRecommendBlog(blogs []vo.SimpleBlogVo) error
	//存取所有博客到缓存 使用set随机
	SaveRandomBlog(blogs []vo.SimpleBlogVo) error
	//获取随随机博客
	GetRandomBlog() (blogs []vo.SimpleBlogVo)
	//存取首页第一页 很有必要,加载会快很多
	SaveHomeFirstPageBlog(pageInfo response.PageInfo) error
	//获取首页的第一页博客内容
	GetHomeFirstPageBlog() (pageInfo response.PageInfo)
}
