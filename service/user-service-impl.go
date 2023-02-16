package service

import (
	"golang.org/x/crypto/bcrypt"
	"time"
	"vs-blog-api/common"
	"vs-blog-api/config"
	"vs-blog-api/modal"
	"vs-blog-api/repository"
	"vs-blog-api/response"
	"vs-blog-api/utils"
)

type UserServiceImpl struct {
}

func (u UserServiceImpl) Logout(username string) {
	config.Redis.Del(common.UserToken + ":" + utils.ToMd5String([]byte(username)))
}

func NewUserService() UserService {
	return UserServiceImpl{}
}

var userRepository = repository.NewUserRepository()

func (u UserServiceImpl) Login(userRequest modal.UserLoginRequest) response.Token {

	user, err := userRepository.FindByUsername(userRequest.Username)

	if err != nil {
		panic(response.NewGlobalException(response.NOTFOUND, "不存在的账号"))
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userRequest.Password))

	if err != nil {
		panic(response.NewGlobalException(response.ParamsError, "密码错误"))
	}

	token := utils.CreateAccessToken(user.Username, user.Nickname)

	var username = utils.ToMd5String([]byte(user.Username))

	config.Redis.Set(common.UserToken+":"+username, token, common.TokenExpire)

	now := time.Now()

	return response.Token{
		Type:   common.TokenPrefix,
		Token:  token,
		Create: utils.FormatDate(now, utils.FORMAT_DATE_TIME),
		Expire: utils.FormatDate(now.Add(common.TokenExpire), utils.FORMAT_DATE_TIME),
	}
}

func (u UserServiceImpl) RegisteredUser(user modal.User) {

	userId := repository.FindByUsername(user.Username).Id

	if userId > 0 {
		panic(response.NewGlobalException(response.ERROR, "该用户已存在,请换个账号吧!"))
	}

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		panic(response.NewGlobalException(response.ERROR, "密码加密失败"))
	}

	user.Password = string(password)

	err2 := userRepository.SaveUser(user)

	if err2 != nil {
		panic(response.NewGlobalException(response.ERROR, "服务器异常,注册失败"))
	}
}
