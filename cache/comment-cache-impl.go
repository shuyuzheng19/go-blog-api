package cache

import (
	"gin-demo/common"
	"github.com/go-redis/redis"
	"strconv"
)

type CommentCacheImpl struct {
	redis *redis.Client
}

func (c CommentCacheImpl) GetUserCommentLikes(uid int) []string {
	return c.redis.SMembers(common.COMMENT_USER_LIKE + strconv.Itoa(uid)).Val()
}

func (c CommentCacheImpl) AddUserLike(uid int, cid int64) int64 {
	return c.redis.SAdd(common.COMMENT_USER_LIKE+strconv.Itoa(uid), cid).Val()
}

func (c CommentCacheImpl) CancelUserLike(uid int, cid int64) int64 {
	return c.redis.SRem(common.COMMENT_USER_LIKE+strconv.Itoa(uid), cid).Val()
}

func (c CommentCacheImpl) UserIsLikeComment(uid int, cid int64) bool {
	return c.redis.SIsMember(common.COMMENT_USER_LIKE+strconv.Itoa(uid), cid).Val()
}

func NewCommentCache(redis *redis.Client) CommentCache {
	return CommentCacheImpl{redis: redis}
}
