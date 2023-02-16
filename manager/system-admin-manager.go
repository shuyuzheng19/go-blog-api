package manager

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/olivere/elastic/v7"
	"strconv"
	"strings"
	"vs-blog-api/common"
	"vs-blog-api/config"
	"vs-blog-api/modal"
	"vs-blog-api/response"
	"vs-blog-api/utils"
)

type SystemManager struct {
}

func NewSystemManager() *SystemManager {
	return &SystemManager{}
}

// 将数据库博客数据存储到ES
func (*SystemManager) DbToElastic() {

	//exists, _ := config.ES.IndexExists(common.BlogIndex).Do(context.Background())
	//
	//if !exists {
	//	config.ES.Index().Index(common.BlogIndex).BodyJson(&modal.EsBlog{}).Do(context.Background())
	//}

	var blogs []modal.Blog

	err := config.DB.Model(&modal.Blog{}).Find(&blogs).Error

	if err != nil {
		panic(response.NewGlobalException(response.ERROR, "查询博客失败"))
	}

	bulkRequest := config.ES.Bulk()
	for _, blog := range blogs {

		doc := elastic.NewBulkCreateRequest().Index(common.BlogIndex).Id(strconv.Itoa(blog.Id)).Doc(&modal.EsBlog{
			Id:          blog.Id,
			Title:       blog.Title,
			Description: blog.Description,
		})

		bulkRequest.Add(doc)
	}

	bulkRequest.Do(context.Background())
}

// 将redis用户点赞信息更新到到DB
func (*SystemManager) RedisLikeToDb() {

	//获取所有用户的KEY
	keys := config.Redis.Keys(common.UserLikeSet + ":*").Val()

	if len(keys) > 0 {

		for _, key := range keys {
			var userId = strings.Split(key, ":")[1]

			members := config.Redis.ZRevRangeWithScores(key, 0, -1).Val()

			zStr, err := json.Marshal(&members)

			if err == nil {
				model := config.DB.Model(&modal.UserLike{})
				var count int64
				model.Where("user_id = ?", userId).Count(&count)
				if count == 0 {
					model.Create(&modal.UserLike{
						UserId: utils.ToInt(userId),
						Liked:  string(zStr),
					})
				} else {
					model.Where("user_id = ?", userId).Updates(&modal.UserLike{
						Liked: string(zStr),
					})
				}
			}

		}
	}

}

// 将DB用户点赞信息初始化到redis
func (*SystemManager) DbLikeToRedis() {
	var likes []modal.UserLike

	err := config.DB.Model(&modal.UserLike{}).Find(&likes).Error

	if err == nil && len(likes) > 0 {
		for _, like := range likes {
			var redisZ []redis.Z
			var KEY = common.UserLikeSet + ":" + strconv.Itoa(like.UserId)
			err := json.Unmarshal([]byte(like.Liked), &redisZ)
			if err == nil {
				config.Redis.ZAdd(KEY, redisZ...)
			}
		}
	}

}

// 更新点赞数量和浏览数量
func (*SystemManager) RedisEyeAndLikeCountToDb() {
	eyes := config.Redis.HGetAll(common.EyeCountMap).Val()
	if len(eyes) > 0 {
		for key, value := range eyes {
			config.DB.Model(&modal.Blog{}).Where("id = ?", key).UpdateColumn("eye_count", value)
		}
	}

	likes := config.Redis.HGetAll(common.LikeCountMap).Val()

	if len(likes) > 0 {
		for key, value := range likes {
			config.DB.Model(&modal.Blog{}).Where("id = ?", key).UpdateColumn("like_count", value)
		}
	}

	config.Redis.Del(common.LikeCountMap, common.EyeCountMap, common.HotBlog, common.UserTodayEyeBlog+":*")
}
