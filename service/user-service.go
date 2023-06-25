package service

import (
	"gin-demo/models"
	"gin-demo/request"
	"gin-demo/response"
	"image"
)

type UserService interface {
	//注册一个用户
	RegisteredUser(user models.User)
	//获取用户
	GetUser(username string) (user models.User)
	//发送邮件
	SendEmail(toEmail string, form string, subject string) (code string)
	//验证邮箱验证码
	ValidateEmailCode(toEmail string, code string) bool
	//生成图形验证码
	GenerateImageCode(ip string) image.Image
	//验证图形验证码是否正确
	ValidateImageCode(ip string, code string) bool
	//账号登录
	Login(userRequest request.UserLoginRequest) response.Token
	//联系我
	SendMessageToMyMail(contactRequest request.ContactRequest)
	//获取网站配置
	GetBlogConfigInfo() response.BlogConfigInfo
	//重新设置网站配置
	SetBlogConfigInfo(config response.BlogConfigInfo)
	//退出账号
	Logout(username string)
}
