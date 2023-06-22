package cache

import "gin-demo/models"

type UserCache interface {
	//通过用户名查找缓存的用户
	FindByUsernameCache(username string) (user models.User)
	//将用户缓存到redis
	SaveUserToCache(user models.User) error
	//保存邮箱验证码到redis
	SaveEmailCodeToRedis(email string, code string) error
	///获取对应的邮箱验证码
	GetEmailCode(toEmail string) string
	//存取登录图像验证码
	SaveLoginCode(ip string, code string) (err error)
	//验证登录验证码
	GetLoginCode(ip string) string
	//保存Token
	SaveToken(username string, token string) error
	//获取Token
	GetToken(username string) string
}
