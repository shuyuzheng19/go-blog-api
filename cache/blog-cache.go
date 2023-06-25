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
	GetHomeFirstPageBlog() string
	//每搜索一次增加一次关键字的搜索量
	SearchCountPlusOne(keyword string) error
	//获取前10条热门搜索
	GetTop10SearchKeyword() []string
	//保存博客详情信息到redis
	SaveBlogToMap(blog vo.BlogContentVo) error
	//获取博客的详情信息
	GetBlogFromMap(id string) string
	//浏览量+1
	IncreaseInView(defaultCount int64, id string) (count int64)
	//点赞量+1
	IncreaseInLike(id string) int64
	//获取点赞量
	GetLikeCount(id string) int64
	//保存用户编写的内容
	SaveUserEditorContent(uid string, content string) error
	//获取用户所编写的博客内容
	GetUserEditorContent(uid string) string
	//添加随即博客
	AddRandomBlog(blog vo.SimpleBlogVo) error
	//删除键值
	RemoveKeys(keys []string) int64
	//删除BLOG-INFO-MAP的博客
	RemoveBlogFromMap(id string) error
	//获取博客的点赞数量
	GetBlogLikeCount() map[string]string
	//获取博客的浏览数量
	GetBlogEyeCount() map[string]string
}
