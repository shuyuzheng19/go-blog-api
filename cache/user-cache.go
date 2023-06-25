package cache

import (
	"gin-demo/models"
	"gin-demo/response"
)

type UserCache interface {
	//通过用户名查找缓存的用户
	FindByUsernameCache(username string) string
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
	//获取网站配置信息
	GetBlogConfig() string
	//重新设置网站配置
	SetBlogConfig(config response.BlogConfigInfo) error
	//删除TOKEN
	RemoveAccessToken(username string) error
}
