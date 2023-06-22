package routers

import (
	"gin-demo/common"
	"gin-demo/controllers"
	"gin-demo/middleware"
	"gin-demo/service"
	"github.com/gin-gonic/gin"
)

type Routers struct {
	router gin.RouterGroup
}

const API_PREFIX = "/api/v1"

func (r Routers) UserRouter() {
	var userService = service.NewUserService()

	var userController = controllers.NewUserController(userService)

	var group = r.router.Group(API_PREFIX + "/user")

	{
		group.POST("/registered", userController.RegisteredUser)
		group.GET("/send_mail", userController.SendEmailCode)
		group.GET("/captcha", userController.GenerateLoginImageCode)
		group.POST("/login", userController.Login)
		group.POST("/contact", userController.ContactMe)
		group.GET("/auth/get", middleware.Authorized(common.USER_ROLE), userController.GetCurrentUser)
	}
}

func (r Routers) BlogRouter() {
	var blogService = service.NewBlogService()

	var blogController = controllers.NewBlogController(blogService)

	var group = r.router.Group(API_PREFIX + "/blog")

	{
		group.GET("/list", blogController.FindBlogPage)
		group.GET("/hots", blogController.GetHotBlog)
		group.GET("/random", blogController.GetRandomBlog)
		group.GET("/recommend", blogController.GetRecommendBlog)
		group.POST("/recommend", middleware.Authorized(common.SUPER_ROLE), blogController.SetRecommendBlog)
	}
}

func (r Routers) SetupRouter() {
	r.UserRouter()
	r.BlogRouter()
}

func NewRouters(router gin.RouterGroup) Routers {
	return Routers{router: router}
}
