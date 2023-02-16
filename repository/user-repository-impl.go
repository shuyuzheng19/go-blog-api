package repository

import (
	"encoding/json"
	"time"
	"vs-blog-api/common"
	"vs-blog-api/config"
	"vs-blog-api/modal"
	"vs-blog-api/response"
	"vs-blog-api/utils"
)

type UserRepositoryImpl struct {
}

func FindByUsername(username string) modal.User {

	var user modal.User

	md5Username := utils.ToMd5String([]byte(username))

	var KEY = common.User + ":" + md5Username

	resultUser := config.Redis.Get(KEY).Val()

	if resultUser == "" {
		err := config.DB.Model(&user).Preload("Role").First(&user, "username = ?", username).Error

		if err != nil {
			panic(response.NewGlobalException(response.NOTFOUND, "找不到该用户"))
		}

		config.Redis.Set(KEY, utils.ObjectToJson(user), time.Minute*30)

	} else {
		err := json.Unmarshal([]byte(resultUser), &user)

		if err != nil {
			return modal.User{}
		}

	}

	return user

}

func (u UserRepositoryImpl) FindByUsername(username string) (modal.User, error) {

	var user modal.User

	err := config.DB.Model(&user).Preload("Role").First(&user, "username = ?", username).Error

	return user, err
}

func (u UserRepositoryImpl) SaveUser(user modal.User) error {

	err := config.DB.Model(&user).Create(&user).Error

	return err
}

func (u UserRepositoryImpl) UpdateUser(user modal.User) error {
	panic("implement me")
}

func (u UserRepositoryImpl) DeleteUserById(id int) error {
	panic("implement me")
}

func NewUserRepository() UserRepository {
	return UserRepositoryImpl{}
}
