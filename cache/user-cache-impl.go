package cache

import (
	"encoding/json"
	"gin-demo/common"
	"gin-demo/models"
	"github.com/go-redis/redis"
)

type UserCacheImpl struct {
	redis *redis.Client
}

func (b UserCacheImpl) GetToken(username string) string {
	var token = b.redis.Get(common.TOKEN_KEY + username).Val()
	return token
}

func (b UserCacheImpl) SaveToken(username string, token string) error {
	return b.redis.Set(common.TOKEN_KEY+username, token, common.TOKEN_EXPIRE).Err()
}

func (b UserCacheImpl) GetLoginCode(ip string) string {
	return b.redis.Get(common.LOGIN_IMAGE_CODE_KEY + ip).Val()
}

func (b UserCacheImpl) SaveLoginCode(ip string, code string) (err error) {
	return b.redis.Set(common.LOGIN_IMAGE_CODE_KEY+ip, code, common.LOGIN_IMAGE_CODE_EXPIRE).Err()
}

func (b UserCacheImpl) GetEmailCode(email string) string {
	return b.redis.Get(common.EMAIL_CODE_KEY + email).Val()
}

func (b UserCacheImpl) SaveEmailCodeToRedis(email string, code string) error {
	return b.redis.Set(common.EMAIL_CODE_KEY+email, code, common.EMAIL_CODE_EXPIRE).Err()
}

func (b UserCacheImpl) SaveUserToCache(user models.User) error {
	var userJson, err = json.Marshal(&user)

	if err != nil {
		return err
	}

	return b.redis.Set(common.USER_INFO_KEY+user.Username, userJson, common.USER_INFO_EXPIRE).Err()
}

func (b UserCacheImpl) FindByUsernameCache(username string) (user models.User) {
	var result = b.redis.Get(common.USER_INFO_KEY + username).Val()

	if result == "" {
		return models.User{}
	}

	var err = json.Unmarshal([]byte(result), &user)

	if err != nil {
		return models.User{}
	}

	return user
}

func NewUserCache(redis *redis.Client) UserCache {
	return UserCacheImpl{redis: redis}
}
