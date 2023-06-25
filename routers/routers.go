package routers

import (
	"gin-demo/common"
	"gin-demo/config"
	"gin-demo/controllers"
	"gin-demo/middleware"
	"gin-demo/service"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
)

type Routers struct {
	router gin.RouterGroup
	cron   *cron.Cron
}

const API_PREFIX = "/api/v1"

func (r Routers) EnableCronJob() {
	r.cron.Start()
	select {}
}

func (r Routers) UserRouter() {
	var userService = service.NewUserService()

	var userController = controllers.NewUserController(userService)

	r.router.GET(API_PREFIX+"/config", userController.GetBlogConfigInfo)

	var group = r.router.Group(API_PREFIX + "/user")

	{
		group.POST("/registered", userController.RegisteredUser)
		group.GET("/send_mail", userController.SendEmailCode)
		group.GET("/captcha", userController.GenerateLoginImageCode)
		group.POST("/login", userController.Login)
		group.POST("/contact", userController.ContactMe)
		group.GET("/auth/logout", middleware.Authorized(common.USER_ROLE), userController.Logout)
		group.GET("/auth/get", middleware.Authorized(common.USER_ROLE), userController.GetCurrentUser)
	}
}

func (r Routers) TagRouters() {
	var tagService = service.NewTagService()

	var tagController = controllers.NewTagController(tagService)

	var group = r.router.Group(API_PREFIX + "/tags")
	{
		group.GET("/random", tagController.RandomController)
		group.GET("/list", tagController.GetTagList)
		group.GET("/:tid/blogs", tagController.GetBlogByTagId)
		group.GET("/get/:tid", tagController.GetTagById)
		group.POST("/admin/add", middleware.Authorized(common.ADMIN_ROLE), tagController.AddTag)
	}
}

func (r Routers) CategoryRouters() {
	var categoryService = service.NewCategoryService()

	var categoryController = controllers.NewCategoryController(categoryService)

	var group = r.router.Group(API_PREFIX + "/category")
	{
		group.GET("/list", categoryController.GetCategoryListForCache)
		group.GET("/list2", middleware.Authorized(common.ADMIN_ROLE), categoryController.GetCategoryListForDB)
		group.POST("/admin/add", middleware.Authorized(common.ADMIN_ROLE), categoryController.AddCategory)
	}
}

func (r Routers) CommentRouters() {
	var commentService = service.NewCommentService()

	var commentController = controllers.NewCommentController(commentService)

	var group = r.router.Group(API_PREFIX + "/comments")
	{
		group.GET("/:bid", commentController.GetBlogComment)
		group.GET("/auth/like/:cid", middleware.Authorized(common.USER_ROLE), commentController.LikeComment)
		group.GET("/auth/user", middleware.Authorized(common.USER_ROLE), commentController.GetCommentUser)
		group.POST("/auth/add_comment", middleware.Authorized(common.USER_ROLE), commentController.AddComment)
	}
}

func (r Routers) FileRouters() {
	var fileService = service.NewFileService()

	var fileController = controllers.NewFileController(fileService)

	var group = r.router.Group(API_PREFIX + "/file")
	{
		group.POST("/upload/avatar", fileController.UploadAvatar)
		group.GET("/public", fileController.GetPublicFile)
		group.GET("/admin/current", middleware.Authorized(common.ADMIN_ROLE), fileController.GetCurrentUserFile)
		group.POST("/upload/auth/image", middleware.Authorized(common.USER_ROLE), fileController.UploadImage)
		group.POST("/upload/admin/file", middleware.Authorized(common.ADMIN_ROLE), fileController.UploadFile)
	}
}

func (r Routers) TopicRouters() {

	var topicService = service.NewTopicService()

	var topicController = controllers.NewTopicController(topicService)

	var group = r.router.Group(API_PREFIX + "/topics")
	{
		group.GET("/list", topicController.GetTopicByPage)
		group.GET("/:tid/list", topicController.GetTopicBlogList)
		group.GET("/user/:uid/list", topicController.GetUserTopic)
		group.GET("/get/:tid", topicController.GetTopicById)
		group.GET("/:tid/blogs", topicController.GetTopicBlogByPage)
		group.GET("/admin/current/list", middleware.Authorized(common.ADMIN_ROLE), topicController.GetCurrentTopics)
		group.POST("/admin/add", middleware.Authorized(common.ADMIN_ROLE), topicController.AddTopic)
	}
}

func (r Routers) BlogRouter() {
	var blogService = service.NewBlogService()

	{
		r.cron.AddFunc("0 0 * * *", func() {
			blogService.InitBlogEyeCount()
		})

		r.cron.AddFunc("0 2 * * *", func() {
			blogService.InitBlogLikeCount()
		})
	}

	if config.CONFIG.InitSearch {
		blogService.InitSearch()
	}

	var blogController = controllers.NewBlogController(blogService)

	var admin = r.router.Group(API_PREFIX + "/admin")

	{
		admin.POST("/recommend", middleware.Authorized(common.ADMIN_ROLE), blogController.SetRecommendBlog)
		admin.GET("/get_edit", middleware.Authorized(common.ADMIN_ROLE), blogController.GetUserEditor)
		admin.POST("/save_edit", middleware.Authorized(common.ADMIN_ROLE), blogController.SaveUserEditor)
		admin.POST("/save", middleware.Authorized(common.ADMIN_ROLE), blogController.SaveBlog)
		admin.GET("/edit/:bid", middleware.Authorized(common.ADMIN_ROLE), blogController.GetEditBlog)
		admin.POST("/update_blog/:bid", middleware.Authorized(common.ADMIN_ROLE), blogController.UpdateBlog)

	}

	var group = r.router.Group(API_PREFIX + "/blog")

	{
		group.GET("/chat", blogController.Chat)
		group.GET("/list", blogController.FindBlogPage)
		group.GET("/hots", blogController.GetHotBlog)
		group.GET("/get/:id", blogController.GetBlogById)
		group.GET("/is_like/:id", blogController.IsLikeBlog)
		group.GET("/like/:id", blogController.LikeBlog)
		group.GET("/range", blogController.RangeBlog)
		group.GET("/search", blogController.SearchBlog)
		group.GET("/search2", blogController.CountSearchBlog)
		group.GET("/hot_keyword", blogController.GetHotKeyword)
		group.GET("/similar", blogController.SearchBlog)
		group.GET("/random", blogController.GetRandomBlog)
		group.GET("/recommend", blogController.GetRecommendBlog)
		group.GET("/user/:uid", blogController.GetBlogByUser)
		group.GET("/user/top/:uid", blogController.GetUserBlogTop)
	}
}

func (r Routers) SetupRouter() {
	r.UserRouter()
	r.BlogRouter()
	r.TagRouters()
	r.CategoryRouters()
	r.TopicRouters()
	r.CommentRouters()
	r.FileRouters()
}

func NewRouters(router gin.RouterGroup) Routers {
	var cron = cron.New()
	return Routers{router: router, cron: cron}
}
