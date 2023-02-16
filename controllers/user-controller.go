package controllers

import (
	"github.com/gin-gonic/gin"
	"image/png"
	"vs-blog-api/common"
	"vs-blog-api/config"
	"vs-blog-api/manager"
	"vs-blog-api/modal"
	"vs-blog-api/response"
	"vs-blog-api/service"
	"vs-blog-api/utils"
)

type userController struct {
}

var emailManager manager.SendEmailManager

var emailConfig config.EmailConfig

var userService service.UserService

var validateManager manager.ValidateCodeManager

func NewUserController() userController {

	emailConfig = config.LoadEmailConfig()

	emailManager = manager.NewSendEmailManager(emailConfig)

	userService = service.NewUserService()

	validateManager = manager.NewValidateManager()

	return userController{}
}

func (*userController) UploadAvatar(ctx *gin.Context) {
	uploadManager.UploadAvatar(ctx, "/avatars", common.ImageTypes, common.MB*5, "图片大小不能大于5MB", "这不是一个图片文件,请核对")
}

func (*userController) Logout(ctx *gin.Context) {
	var username = modal.GetUser(ctx).Username
	userService.Logout(username)
}

func (*userController) SendEmail(ctx *gin.Context) {
	var email = ctx.Query("email")

	if utils.IsEmpty(email) {
		ctx.JSON(200, response.FAILURE(response.ParamsError, "邮箱不能为空"))
		return
	}

	if !utils.IsEmail(email) {
		ctx.JSON(200, response.FAILURE(response.MatherError, "这不是一个正确的邮箱格式"))
		return
	}

	var randomCode = utils.CreateRandomNumber(6)

	emailManager.SendEmail(email, "", "ZSY-BLOG", randomCode)

	ctx.JSON(200, response.OK_RESULT)
}

func (*userController) Login(ctx *gin.Context) {

	var user modal.UserLoginRequest

	ctx.ShouldBindJSON(&user)

	if utils.AnyEmpty(user.Username, user.Password, user.Code) {
		ctx.JSON(200, response.FAILURE(response.ParamsError, "参数有空"))
		return
	}

	var ip = ctx.ClientIP()

	validateManager.VerifyLoginCode(user.Code, ip)

	tokenResponse := userService.Login(user)

	ctx.JSON(200, response.SUCCESS(tokenResponse))

	config.Redis.Del(common.LoginValidateCode + ":" + ip)

}

func (*userController) GetUser(ctx *gin.Context) {

	user := modal.GetUser(ctx)

	ctx.JSON(200, response.SUCCESS(user.ToVo()))

}

func (*userController) CreateCaptcha(ctx *gin.Context) {

	var ip = ctx.ClientIP()

	image := validateManager.GenerateValidateCode(ip)

	writer := ctx.Writer

	writer.Header().Set("content-type", "image/png")

	png.Encode(writer, image)
}

func (*userController) RegisteredUser(ctx *gin.Context) {

	var user = modal.UserRegisteredRequest{}

	ctx.ShouldBindJSON(&user)

	user.Validator()

	emailManager.ValidatorEmailCode(user.Email, user.Code)

	userService.RegisteredUser(user.ToUserDo())

	ctx.JSON(200, response.OK_RESULT)

}
