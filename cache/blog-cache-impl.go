package cache

import (
	"encoding/json"
	"gin-demo/common"
	"gin-demo/response"
	"gin-demo/vo"
	"github.com/go-redis/redis"
)

type BlogCacheImpl struct {
	redis *redis.Client
}

func (b BlogCacheImpl) SaveHomeFirstPageBlog(pageInfo response.PageInfo) error {
	var buff, _ = json.Marshal(&pageInfo)
	return b.redis.Set(common.FIREST_PAGE_BLOG_KEY, buff, common.FIREST_BLOG_PAGE_EXIPRE).Err()
}

func (b BlogCacheImpl) GetHomeFirstPageBlog() (pageInfo response.PageInfo) {
	var result = b.redis.Get(common.FIREST_PAGE_BLOG_KEY).Val()

	json.Unmarshal([]byte(result), &pageInfo)

	return pageInfo
}

func (b BlogCacheImpl) SaveRandomBlog(blogs []vo.SimpleBlogVo) error {

	var jsons []string

	for _, blog := range blogs {
		var buff, _ = json.Marshal(&blog)
		jsons = append(jsons, string(buff))
	}

	var err = b.redis.SAdd(common.RANDOM_BLOG_KEY, jsons).Err()

	if err == nil {
		b.redis.Expire(common.RANDOM_BLOG_KEY, common.RANDOM_BLOG_EXPIRE)
	}

	return err
}

func (b BlogCacheImpl) GetRandomBlog() (blogs []vo.SimpleBlogVo) {
	var result = b.redis.SRandMemberN(common.RANDOM_BLOG_KEY, common.RANDOM_BLOG_SIZE).Val()
	for _, r := range result {
		var blog vo.SimpleBlogVo
		var err = json.Unmarshal([]byte(r), &blog)
		if err == nil {
			blogs = append(blogs, blog)
		}
	}
	return blogs
}

func (b BlogCacheImpl) SetRecommendBlog(blogs []vo.SimpleBlogVo) error {
	var result, _ = json.Marshal(&blogs)
	return b.redis.Set(common.RECOMMEND_BLOG_KEY, result, common.RECOMMEND_BLOG_EXPIRE).Err()
}

func (b BlogCacheImpl) GetRecommendBlog() (blogs []vo.SimpleBlogVo) {

	var result = b.redis.Get(common.RECOMMEND_BLOG_KEY).Val()

	json.Unmarshal([]byte(result), &blogs)

	return blogs
}

func (b BlogCacheImpl) SaveTopBlog(blogs []vo.SimpleBlogVo) error {
	var result, _ = json.Marshal(&blogs)
	return b.redis.Set(common.HOT_BLOG_KEY, result, common.HOT_BLOG_EXIPRE).Err()
}

func (b BlogCacheImpl) GetTopBlog() (blogs []vo.SimpleBlogVo) {
	var result = b.redis.Get(common.HOT_BLOG_KEY).Val()

	json.Unmarshal([]byte(result), &blogs)

	return blogs
}

func NewBlogCache(redis *redis.Client) BlogCache {
	return BlogCacheImpl{redis: redis}
}
