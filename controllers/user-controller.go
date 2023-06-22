package controllers

import (
	"gin-demo/common"
	"gin-demo/config"
	"gin-demo/models"
	"gin-demo/myerr"
	"gin-demo/request"
	"gin-demo/service"
	"gin-demo/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"image/png"
	"net/http"
)

type UserController struct {
	service service.UserService
}

func (u UserController) RegisteredUser(ctx *gin.Context) {

	var userRequest request.UserRequest

	var err = ctx.ShouldBindJSON(&userRequest)

	if err != nil {
		ctx.JSON(http.StatusOK, common.BAD_REQUEST_ERROR)
		return
	}

	var ok = u.service.ValidateEmailCode(userRequest.Email, userRequest.Code)

	if ok {
		var checkError = userRequest.Check()

		myerr.IError(checkError, common.VALIDATOR_CODE)

		u.service.RegisteredUser(userRequest.ToUserDo())

		ctx.JSON(http.StatusOK, common.OK())
	} else {
		ctx.JSON(http.StatusOK, common.EMAIL_CODE_ERROR)
	}

}

func (u UserController) GenerateLoginImageCode(ctx *gin.Context) {
	var ip = ctx.ClientIP()
	var img = u.service.GenerateImageCode(ip)
	var write = ctx.Writer
	png.Encode(write, img)
}

func (u UserController) SendEmailCode(ctx *gin.Context) {

	var email = ctx.Query("email")

	if !utils.IsEmailValid(email) {
		ctx.JSON(http.StatusOK, common.EMAIL_WRONG_FORMAT)
		return
	}

	u.service.SendEmail(email, "", common.SEND_EMAIL_SUBJECT)

	ctx.JSON(common.OK_CODE, common.OK())

	config.LOGGER.Info("发送邮件成功", zap.String("ip", ctx.ClientIP()))
}

func (u UserController) GetCurrentUser(ctx *gin.Context) {
	var user, exists = ctx.Get("user")
	if exists {
		ctx.JSON(http.StatusOK, common.Success(user.(models.User).ToUserVo()))
	} else {
		ctx.JSON(http.StatusOK, common.USER_NOT_FOUNT)
	}
}

func (u UserController) Login(ctx *gin.Context) {
	var userLoginRequest request.UserLoginRequest
	var err = ctx.ShouldBindJSON(&userLoginRequest)
	if err != nil {
		myerr.PanicError(common.BAD_REQUEST_ERROR)
	}
	var validateError = userLoginRequest.Check()

	myerr.IError(validateError, common.VALIDATOR_CODE)

	var flag = u.service.ValidateImageCode(ctx.ClientIP(), userLoginRequest.Code)

	if !flag {
		myerr.PanicError(common.LOGIN_CODE_ERROR)
	}

	var tokenResponse = u.service.Login(userLoginRequest)

	tokenResponse.Ip = utils.GetIPAddress(ctx.Request)

	ctx.JSON(http.StatusOK, common.Success(tokenResponse))

	config.LOGGER.Info("登录成功",
		zap.String("ip", tokenResponse.Ip),
		zap.String("username", tokenResponse.Username),
		zap.String("role", tokenResponse.Role),
	)

}

func (u UserController) ContactMe(ctx *gin.Context) {

	var contactRequest request.ContactRequest

	var err = ctx.ShouldBindJSON(&contactRequest)

	if err != nil {
		ctx.JSON(http.StatusOK, common.BAD_REQUEST_ERROR)
		return
	}

	var checkError = contactRequest.Check()

	myerr.IError(checkError, common.VALIDATOR_CODE)

	u.service.SendMessageToMyMail(contactRequest)

	ctx.JSON(http.StatusOK, common.OK())

}

func NewUserController(service service.UserService) *UserController {
	return &UserController{service: service}
}
