package cache

import (
	"encoding/json"
	"gin-demo/common"
	"gin-demo/response"
	"gin-demo/vo"
	"github.com/go-redis/redis"
	"strconv"
)

type TopicCacheImpl struct {
	redis *redis.Client
}

func (t TopicCacheImpl) RemoveTopicPageKey() error {
	return t.redis.Del(common.FIRST_PAGE_TOPIC_KEY).Err()
}

func (t TopicCacheImpl) SaveFirstPageTopic(pageInfo response.PageInfo) error {
	var buff, err = json.Marshal(&pageInfo)

	if err != nil {
		return err
	}

	return t.redis.Set(common.FIRST_PAGE_TOPIC_KEY, string(buff), common.FIRST_PAGE_TOPIC_EXIPRE).Err()
}

func (t TopicCacheImpl) GetFirstPageTopic() string {
	return t.redis.Get(common.FIRST_PAGE_TOPIC_KEY).Val()
}

func (t TopicCacheImpl) SaveFirstPageBlog(tid string, pageInfo response.PageInfo) error {
	var buff, err = json.Marshal(&pageInfo)
	if err != nil {
		return err
	}
	return t.redis.Set(common.FIRST_TOPIC_BLOG_PAGE_KEY+tid, string(buff), common.FIRST_TOPIC_BLOG_PAGE_EXPIRE).Err()
}

func (t TopicCacheImpl) GetFirstPageBlog(tid string) string {
	return t.redis.Get(common.FIRST_TOPIC_BLOG_PAGE_KEY + tid).Val()
}

func (t TopicCacheImpl) SetTopicToMap(topic vo.SimpleTopicVo) error {
	var buff, err = json.Marshal(&topic)

	if err != nil {
		return err
	}

	var err2 = t.redis.HSet(common.TOPIC_MAP_KEY, strconv.Itoa(topic.Id), string(buff)).Err()

	if err2 != nil {
		return err2
	}

	t.redis.Expire(common.TOPIC_MAP_KEY, common.TOPIC_MAP_EXPIRE)

	return nil
}

func (t TopicCacheImpl) GetTopicFromMap(id string) string {
	return t.redis.HGet(common.TOPIC_MAP_KEY, id).Val()
}

func NewTopicCache(redis *redis.Client) TopicCache {
	return TopicCacheImpl{redis: redis}
}
