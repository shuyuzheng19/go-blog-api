package cache

import (
	"encoding/json"
	"gin-demo/common"
	"gin-demo/models"
	"gin-demo/vo"
	"github.com/go-redis/redis"
)

type CategoryCacheImpl struct {
	redis *redis.Client
}

func (c CategoryCacheImpl) RemoveKey() error {
	return c.redis.Del(common.CATEGORY_LIST_KEY).Err()
}

func (c CategoryCacheImpl) AddCategory(category models.Category) error {
	var buff, _ = json.Marshal(&category)

	return c.redis.LPush(common.CATEGORY_LIST_KEY, string(buff)).Err()
}

func (c CategoryCacheImpl) SaveCategoryList(list []vo.CategoryVo) error {
	var categoryStrs []string
	for _, category := range list {
		var buff, _ = json.Marshal(&category)
		categoryStrs = append(categoryStrs, string(buff))
	}
	err := c.redis.LPush(common.CATEGORY_LIST_KEY, categoryStrs).Err()
	if err != nil {
		return err
	}
	c.redis.Expire(common.CATEGORY_LIST_KEY, common.CATEGORY_LIST_EXPIRE)
	return nil
}

func (c CategoryCacheImpl) GetCategoryList() []string {
	return c.redis.LRange(common.CATEGORY_LIST_KEY, 0, -1).Val()
}

func NewCategoryCache(redis *redis.Client) CategoryCache {
	return CategoryCacheImpl{redis: redis}
}
