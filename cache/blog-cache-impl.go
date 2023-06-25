package cache

import (
	"encoding/json"
	"gin-demo/common"
	"gin-demo/response"
	"gin-demo/vo"
	"github.com/go-redis/redis"
	"strconv"
)

type BlogCacheImpl struct {
	redis *redis.Client
}

func (b BlogCacheImpl) GetBlogLikeCount() map[string]string {
	return b.redis.HGetAll(common.BLOG_LIKE_COUNT_MAP_KEY).Val()
}

func (b BlogCacheImpl) GetBlogEyeCount() map[string]string {
	return b.redis.HGetAll(common.BLOG_EYE_COUNT_MAP_KEY).Val()
}

func (b BlogCacheImpl) RemoveBlogFromMap(id string) error {
	return b.redis.HDel(common.BLOG_INFO_MAP_KEY, id).Err()
}

func (b BlogCacheImpl) AddRandomBlog(blog vo.SimpleBlogVo) error {

	var buff, _ = json.Marshal(&blog)

	var err = b.redis.SAdd(common.RANDOM_BLOG_KEY, string(buff)).Err()

	return err
}

func (b BlogCacheImpl) RemoveKeys(keys []string) int64 {
	return b.redis.Del(keys...).Val()
}

func (b BlogCacheImpl) SaveUserEditorContent(uid string, content string) error {
	return b.redis.HSet(common.USER_EDITOR_SAVE_MAP, uid, content).Err()
}

func (b BlogCacheImpl) GetUserEditorContent(uid string) string {
	return b.redis.HGet(common.USER_EDITOR_SAVE_MAP, uid).Val()
}

func (b BlogCacheImpl) GetLikeCount(id string) int64 {

	var count, err = b.redis.HGet(common.BLOG_LIKE_COUNT_MAP_KEY, id).Int64()

	if err != nil {
		return 0
	}

	return count
}

func (b BlogCacheImpl) IncreaseInLike(id string) int64 {
	return b.redis.HIncrBy(common.BLOG_LIKE_COUNT_MAP_KEY, id, 1).Val()
}

func (b BlogCacheImpl) IncreaseInView(defaultCount int64, id string) (count int64) {
	if flag := b.redis.HExists(common.BLOG_EYE_COUNT_MAP_KEY, id).Val(); flag {
		count = b.redis.HIncrBy(common.BLOG_EYE_COUNT_MAP_KEY, id, 1).Val()
	} else {
		count = defaultCount + 1
		b.redis.HSet(common.BLOG_EYE_COUNT_MAP_KEY, id, count)
	}
	return count
}

func (b BlogCacheImpl) SaveBlogToMap(blog vo.BlogContentVo) error {
	var buff, err = json.Marshal(&blog)
	if err != nil {
		return err
	}
	return b.redis.HSet(common.BLOG_INFO_MAP_KEY, strconv.Itoa(blog.Id), string(buff)).Err()
}

func (b BlogCacheImpl) GetBlogFromMap(id string) string {
	return b.redis.HGet(common.BLOG_INFO_MAP_KEY, id).Val()
}

func (b BlogCacheImpl) GetTop10SearchKeyword() []string {
	return b.redis.ZRevRangeByScore(common.SEARCH_KEYWORD_KEY, redis.ZRangeBy{
		Min:    "0",
		Max:    "+inf",
		Offset: 0,
		Count:  10,
	}).Val()
}

func (b BlogCacheImpl) SearchCountPlusOne(keyword string) error {
	return b.redis.ZIncrBy(common.SEARCH_KEYWORD_KEY, 1, keyword).Err()
}

func (b BlogCacheImpl) SaveHomeFirstPageBlog(pageInfo response.PageInfo) error {
	var buff, _ = json.Marshal(&pageInfo)
	return b.redis.Set(common.FIRST_PAGE_BLOG_PAGE_KEY, buff, common.FIRST_BLOG_PAGE_EXIPRE).Err()
}

func (b BlogCacheImpl) GetHomeFirstPageBlog() string {
	var result = b.redis.Get(common.FIRST_PAGE_BLOG_PAGE_KEY).Val()

	return result
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
