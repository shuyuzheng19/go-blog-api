package cache

import (
	"encoding/json"
	"gin-demo/common"
	"gin-demo/models"
	"gin-demo/response"
	"gin-demo/vo"
	"github.com/go-redis/redis"
	"strconv"
)

type TagCacheImpl struct {
	redis *redis.Client
}

func (t TagCacheImpl) RemoveKey() error {
	return t.redis.Del(common.RANDOM_TAG_KEY).Err()
}

func (t TagCacheImpl) AddTag(tag models.Tag) error {
	var buff, _ = json.Marshal(&tag)
	return t.redis.SAdd(common.RANDOM_TAG_KEY, string(buff)).Err()
}

func (t TagCacheImpl) SaveFirstPageBlog(tid string, pageInfo response.PageInfo) error {
	var buff, err = json.Marshal(&pageInfo)
	if err != nil {
		return err
	}
	return t.redis.Set(common.FIRST_TAG_BLOG_PAGE_KEY+tid, string(buff), common.FIRST_TAG_BLOG_EXPIRE).Err()
}

func (t TagCacheImpl) GetFirstPageBlog(tid string) string {
	return t.redis.Get(common.FIRST_TAG_BLOG_PAGE_KEY + tid).Val()
}

func (t TagCacheImpl) SaveTagInfo(tag vo.TagVo) error {
	var buff, err = json.Marshal(&tag)
	if err != nil {
		return err
	}
	return t.redis.HSet(common.TAG_MAP_KEY, strconv.Itoa(tag.Id), string(buff)).Err()
}

func (t TagCacheImpl) GetTagInfo(id string) string {
	return t.redis.HGet(common.TAG_MAP_KEY, id).Val()
}

func (t TagCacheImpl) SaveAllTags(tags []vo.TagVo) {

	var jsons = make([]string, 0)

	for _, tag := range tags {
		var buff, _ = json.Marshal(&tag)
		jsons = append(jsons, string(buff))
	}

	t.redis.SAdd(common.RANDOM_TAG_KEY, jsons).Err()

	t.redis.Expire(common.RANDOM_TAG_KEY, common.RANDOM_TAG_EXPIRE)
}

func (t TagCacheImpl) GetRandomTags() []string {
	return t.redis.SRandMemberN(common.RANDOM_TAG_KEY, common.RANDOM_TAG_COUNT).Val()
}

func NewTagCache(redis *redis.Client) TagCache {
	return TagCacheImpl{redis: redis}
}
