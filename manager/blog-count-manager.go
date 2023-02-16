package manager

import (
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
	"time"
	"vs-blog-api/common"
	"vs-blog-api/config"
	"vs-blog-api/response"
	"vs-blog-api/utils"
)

type BlogCountManager struct {
}

func NewBlogCountManager() *BlogCountManager {
	return &BlogCountManager{}
}

func (*BlogCountManager) GetLikeCount(blogId int, defaultCount int64) int64 {
	count := config.Redis.HGet(common.LikeCountMap, strconv.Itoa(blogId)).Val()

	if count == "" {
		config.Redis.HSet(common.LikeCountMap, strconv.Itoa(blogId), defaultCount)
		return defaultCount
	}

	return utils.ToInt64(count)
}

func (*BlogCountManager) IsLike(userId int, blogId int) bool {

	if userId <= 0 || blogId <= 0 {
		panic(response.NewGlobalException(response.ParamsError, "含有非法参数"))
	}

	count := config.Redis.ZScore(common.UserLikeSet+":"+strconv.Itoa(userId), strconv.Itoa(blogId)).Val()

	fmt.Println(count)

	if count <= 0 {
		return false
	}

	return true
}

func (*BlogCountManager) Like(userId int, blogId int) int64 {
	countStr := config.Redis.HGet(common.LikeCountMap, strconv.Itoa(blogId)).Val()

	if countStr != "" {
		config.Redis.ZAdd(common.UserLikeSet+":"+strconv.Itoa(userId), redis.Z{
			Score:  float64(time.Now().Unix()),
			Member: blogId,
		})
		config.Redis.HIncrBy(common.LikeCountMap, strconv.Itoa(blogId), 1)
		return utils.ToInt64(countStr) + 1
	}

	return -1
}

func (*BlogCountManager) UnLike(userId int, blogId int) int64 {

	var KEY = common.UserLikeSet + ":" + strconv.Itoa(userId)

	countStr := config.Redis.HGet(common.LikeCountMap, strconv.Itoa(blogId)).Val()

	if countStr != "" {
		config.Redis.HIncrBy(common.LikeCountMap, strconv.Itoa(blogId), -1)
		config.Redis.ZRem(KEY, blogId)
		return utils.ToInt64(countStr) - 1
	}

	return -1
}

func (*BlogCountManager) GetBlogEyeCount(blogId int, defaultCount int64) int64 {

	count, err := config.Redis.HGet(common.EyeCountMap, strconv.Itoa(blogId)).Int64()

	if err != nil {
		return defaultCount
	}

	return count
}

// 博客浏览量+1
func (*BlogCountManager) BlogEyeCountAdd(blogId int, defaultCount int64) int64 {

	var strBlogId = strconv.Itoa(blogId)

	exists := config.Redis.HExists(common.EyeCountMap, strBlogId).Val()

	if !exists {
		config.Redis.HSet(common.EyeCountMap, strBlogId, defaultCount+1)
		return defaultCount + 1
	}

	newCount := config.Redis.HIncrBy(common.EyeCountMap, strBlogId, 1).Val()

	return newCount

}

// 用户今天是否已经查看了博客
func (*BlogCountManager) IsUserTodayEye(userId string, blogId int) bool {
	if userId == "" || blogId <= 0 {
		panic(response.NewGlobalException(response.ParamsError, "含有非法参数"))
	}

	var KEY = common.UserTodayEyeBlog + ":" + userId

	var strBlogId = strconv.Itoa(blogId)

	isEye := config.Redis.SIsMember(KEY, strBlogId).Val()

	if !isEye {
		config.Redis.SAdd(KEY, strBlogId)
	}

	return isEye
}
