package service

import (
	"fmt"
	"gin-demo/cache"
	"gin-demo/common"
	"gin-demo/config"
	"gin-demo/models"
	"gin-demo/myerr"
	"gin-demo/repository"
	"gin-demo/request"
	"gin-demo/response"
	"gin-demo/utils"
	"github.com/afocus/captcha"
	"github.com/jordan-wright/email"
	"go.uber.org/zap"
	"image"
	"net/smtp"
	"time"
)

type UserServiceImpl struct {
	repository repository.UserRepository
	cache      cache.UserCache
}

func (u UserServiceImpl) SendMessageToMyMail(contactRequest request.ContactRequest) {

	var e = email.NewEmail()

	var emailConfig = config.CONFIG.Email

	e.From = fmt.Sprintf("%s <%s>", "", emailConfig.Username)

	e.To = []string{common.MY_EMAIL}

	e.Subject = contactRequest.Subject

	e.HTML = []byte(fmt.Sprintf(
		"<h3>%s</h3><p>对方名字:%s</p><p>对方邮箱:%s</p>留言内容:<p>%s</p>",
		contactRequest.Subject, contactRequest.Name, contactRequest.Email, contactRequest.Content,
	))

	err := e.Send(emailConfig.Addr, smtp.PlainAuth("", emailConfig.Username, emailConfig.Password, emailConfig.Host))

	myerr.ResultError(err, common.SEND_EMAIL_ERROR)

}

var userService *UserServiceImpl

func (u UserServiceImpl) Login(userRequest request.UserLoginRequest) response.Token {
	var user = u.GetUser(userRequest.Username)

	if user.Id == 0 {
		myerr.PanicError(common.USER_NOT_FOUNT)
	}

	var flag = utils.VerifyPassword(userRequest.Password, user.Password)

	if flag {
		var token = utils.CreateAccessToken(user.Username, user.Nickname)
		u.cache.SaveToken(user.Username, token)
		return response.Token{
			Username:    user.Username,
			AccessToken: token,
			Type:        common.TOKEN_TYPE,
			Role:        user.Role.Name,
			CreateAt:    utils.FormatDate(time.Now()),
			ExpireAt:    utils.FormatDate(time.Now().Add(common.TOKEN_EXPIRE)),
		}
	} else {
		myerr.PanicError(common.PASSWORD_ERROR)
	}

	return response.Token{}
}

func (u UserServiceImpl) ValidateImageCode(ip string, code string) bool {
	if ip == "" || code == "" {
		return false
	}

	var redisLoginCode = u.cache.GetLoginCode(ip)

	return redisLoginCode == code
}

func (u UserServiceImpl) GenerateImageCode(ip string) image.Image {
	c := captcha.New()

	c.SetFont(common.IMAGE_FONT_PATH)

	c.SetDisturbance(captcha.NORMAL)

	c.SetSize(100, 40)

	var image, code = c.Create(4, captcha.ALL)

	fmt.Println(code)

	u.cache.SaveLoginCode(ip, code)

	return image
}

func (u UserServiceImpl) ValidateEmailCode(toEmail string, code string) bool {
	if toEmail == "" || code == "" {
		return false
	}
	var redisEmailCode = u.cache.GetEmailCode(toEmail)

	return redisEmailCode == code
}

func (u UserServiceImpl) SendEmail(toEmail string, form string, subject string) (code string) {

	var e = email.NewEmail()

	var emailConfig = config.CONFIG.Email

	e.From = fmt.Sprintf("%s <%s>", form, emailConfig.Username)

	e.To = []string{toEmail}

	e.Subject = subject

	var randomCode = utils.CreateRandomNumber()

	e.Text = []byte("你的验证码:(" + randomCode + ") 两分钟内有效")

	err := e.Send(emailConfig.Addr, smtp.PlainAuth("", emailConfig.Username, emailConfig.Password, emailConfig.Host))

	myerr.ResultError(err, common.SEND_EMAIL_ERROR)

	u.cache.SaveEmailCodeToRedis(toEmail, randomCode)

	config.LOGGER.Info("发送验证码服务", zap.String("email", toEmail), zap.String("code", randomCode))

	return randomCode

}

func (u UserServiceImpl) GetUser(username string) (user models.User) {
	user = u.cache.FindByUsernameCache(username)
	if user.Id == 0 {
		user = u.repository.FindByUsername(username)
		if user.Id > 0 {
			u.cache.SaveUserToCache(user)
		}
	}
	return user
}

func (u UserServiceImpl) RegisteredUser(user models.User) {
	var exists = u.repository.ExistsByUsername(user.Username)

	if exists > 0 {
		myerr.PanicError(common.DUPLICATE_USERNAME_ERROR)
	}

	var err = u.repository.Save(user)

	myerr.ResultError(err, common.REGISTRATION_FAILED)
}

func NewUserService() UserService {
	if userService == nil {
		var userRepository = repository.NewUserRepository(config.DB)
		var cache = cache.NewUserCache(config.REDIS)
		userService = &UserServiceImpl{repository: userRepository, cache: cache}
	}
	return userService
}
