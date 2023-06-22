package service

import (
	"gin-demo/cache"
	"gin-demo/common"
	"gin-demo/config"
	"gin-demo/myerr"
	"gin-demo/repository"
	"gin-demo/request"
	"gin-demo/response"
	"gin-demo/vo"
)

type BlogServiceImpl struct {
	cache      cache.BlogCache
	repository repository.BlogRepository
}

func (b BlogServiceImpl) RandomBlogs() (blogs []vo.SimpleBlogVo) {
	var result = b.cache.GetRandomBlog()
	if len(result) == 0 {
		b.cache.SaveRandomBlog(b.repository.FindAllSimpleBlog())
	}
	return b.cache.GetRandomBlog()
}

func (b BlogServiceImpl) SetRecommend(ids []int) {
	if len(ids) < 4 {
		myerr.PanicError(common.SAVE_RECOMMEND_ERROR)
	}
	var blogs = b.repository.FindByIdIn(ids)

	b.cache.SetRecommendBlog(blogs)
}

func (b BlogServiceImpl) GetRecommend() (blogs []vo.SimpleBlogVo) {
	return b.cache.GetRecommendBlog()
}

func (b BlogServiceImpl) GetHotBlog() []vo.SimpleBlogVo {
	var blogs = b.cache.GetTopBlog()

	if len(blogs) == 0 {
		blogs = b.repository.GetHotBlog()
		b.cache.SaveTopBlog(blogs)
	} else {
		blogs = b.cache.GetTopBlog()
	}

	return blogs
}

func (b BlogServiceImpl) FindBlogByCidPage(pageRequest request.BlogPageRequest) response.PageInfo {

	if pageRequest.Page == 1 {
		var pageInfo = b.cache.GetHomeFirstPageBlog()
		if pageInfo.Total > 0 {
			return pageInfo
		}
	}

	var blogs, count = b.repository.PaginatedBlogQueries(pageRequest)

	var blogVos []vo.BlogVo

	for _, blog := range blogs {
		blogVos = append(blogVos, blog.ToVo())
	}

	var pageInfo = response.PageInfo{
		Page:  pageRequest.Page,
		Size:  common.PAGE_SIZE,
		Total: count,
		Data:  blogVos,
	}

	if pageInfo.Page == 1 {
		b.cache.SaveHomeFirstPageBlog(pageInfo)
	}

	return pageInfo
}

func NewBlogService() BlogService {
	var repository = repository.NewBlogRepository(config.DB)
	var cache = cache.NewBlogCache(config.REDIS)
	return BlogServiceImpl{repository: repository, cache: cache}
}
