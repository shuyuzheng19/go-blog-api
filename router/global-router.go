package router

import (
	"github.com/gin-gonic/gin"
	"vs-blog-api/controllers"
	controller "vs-blog-api/controllers/parse"
)

type globalRouter struct {
	group gin.RouterGroup
}

func NewUserRouter(group gin.RouterGroup) globalRouter {
	return globalRouter{group: group}
}

func (group globalRouter) addUserRouter() {

	var userController = controllers.NewUserController()

	routerGroup := group.group.Group("/api/v1/user")
	{
		//退出登录
		routerGroup.GET("/logout", userController.Logout)
		//注册账号
		routerGroup.POST("/registered", userController.RegisteredUser)
		//发送验证码
		routerGroup.GET("/send_email", userController.SendEmail)
		//获取验证码
		routerGroup.GET("/captcha", userController.CreateCaptcha)
		//登录
		routerGroup.POST("/login", userController.Login)
		//获取当前用户
		routerGroup.GET("/get", userController.GetUser)
	}
}

func (group globalRouter) uploadRouter() {

	var uploadController = controllers.NewUploadController()

	routerGroup := group.group.Group("/api/v1/upload")

	{
		//上传文件
		routerGroup.POST("/file", uploadController.UploadOther)

		//上传头像
		routerGroup.POST("/avatar", uploadController.UploadAvatar)

		//解析md文件
		routerGroup.POST("/md", uploadController.ParseMarkDownFile)

	}
}

func (group globalRouter) blogRouter() {

	var blogController = controllers.NewBlogController()

	routerGroup := group.group.Group("/api/v1/blog")
	{

		//获取用户点赞的博客
		routerGroup.GET("/user/like/:id", blogController.GetUserLikeBlog)

		//添加博客
		routerGroup.POST("/add", blogController.AddBlog)

		//用户是否点赞
		routerGroup.GET("/is_like", blogController.CurrentUserIsLike)

		//用户点赞或取消点赞
		routerGroup.POST("/like", blogController.CurrentUserLikeOrUnLikeBlog)

		//用户保存文章
		routerGroup.POST("/user_save", blogController.SaveUserBlog)

		//获取用户保存文章
		routerGroup.GET("/get/user_save", blogController.GetUserSaveBlog)

		//获取相关文章
		routerGroup.GET("/similar", blogController.GetSimilarBlog)

		//获取某篇博客
		routerGroup.GET("/get/:id", blogController.GetBlogById)

		//获取博客列表
		routerGroup.GET("/list", blogController.FindBlogs)

		//获取热门博客
		routerGroup.GET("/hots", blogController.GetHotBlogs)

		//随机获取博客
		routerGroup.GET("/random", blogController.RandomBlogs)

		//获取某个时间段的博客
		routerGroup.GET("/range", blogController.RangeBlog)

		//搜索博客
		routerGroup.GET("/search", blogController.SearchBlog)

		//获取某个用户的博客
		routerGroup.GET("/user/:id", blogController.FindUserBlog)

		//获取某个用户的博客的TOP10
		routerGroup.GET("/user/:id/top", blogController.FindUserTop10)
	}
}

func (group globalRouter) categoryRouter() {

	var categoryController = controllers.NewCategoryController()

	routerGroup := group.group.Group("/api/v1/category")
	{
		//添加分类
		routerGroup.POST("/add", categoryController.SaveCategory)

		//获取所有分类
		routerGroup.GET("/list", categoryController.GetAllCategory)
	}
}

func (group globalRouter) tagRouter() {

	var tagController = controllers.NewTagController()

	routerGroup := group.group.Group("/api/v1/tag")
	{
		//添加标签
		routerGroup.POST("/add", tagController.SaveTag)

		//获取所有标签
		routerGroup.GET("/list", tagController.GetAllTags)

		//获取随机标签
		routerGroup.GET("/random", tagController.RandomTag)

		//获取某个标签的博客
		routerGroup.GET("/:id", tagController.GetTagBlog)

		//获取某个标签
		routerGroup.GET("/get/:id", tagController.GetTag)
	}
}

func (group globalRouter) topicRouter() {

	var topicController = controllers.NewTopicController()

	routerGroup := group.group.Group("/api/v1/topic")
	{

		//获取专题下所有的博客
		routerGroup.GET("/blog/list/:id", topicController.GetTopicIdBlogs)

		//获取所有专题
		routerGroup.GET("/all", topicController.GetAllTopics)

		//添加专题
		routerGroup.POST("/add", topicController.AddTopic)

		//获取专题列表
		routerGroup.GET("/list", topicController.GetTopics)

		//获取专题下的博客
		routerGroup.GET("/:id", topicController.GetTopicBlog)

		//获取专题信息
		routerGroup.GET("/get/:id", topicController.GetById)

		//获取某个用户的专题
		routerGroup.GET("/user/:userId", topicController.GetUserTopic)
	}
}

func (group globalRouter) otherRouter() {

	var otherController = controllers.NewOtherController()

	routerGroup := group.group.Group("/api/v1/other")
	{

		//获取推荐文章
		routerGroup.GET("/recommend", otherController.GetRecommend)

		//更新推荐文章
		routerGroup.POST("/recommend", otherController.InitRecommend)

		//获取网站动态
		routerGroup.GET("/timeline", otherController.GetTimeLines)

		//获取音乐播放列表
		routerGroup.GET("/music/get", otherController.GetMusicPlayList)

		//更新网易云播放列表
		routerGroup.GET("/music/cloud/:mid", otherController.InitMusicCloud)

	}
}

func (group globalRouter) fileRouter() {

	var fileController = controllers.NewFileController()

	routerGroup := group.group.Group("/api/v1/file")
	{
		//获取当前用户的文件列表
		routerGroup.GET("/current_user/list", fileController.GetCurrentUserFile)

		//获取公告文件列表
		routerGroup.GET("/public/list", fileController.GetPublicFile)

		//删除当前用户 指定ID的文件
		routerGroup.POST("/current_user/delete", fileController.DeleteCurrentUserFiles)

		//获取公告文件列表
		routerGroup.POST("/delete", fileController.DeleteFiles)

	}
}

func (group globalRouter) parseRouter() {
	tikTokController := controller.NewTikTokController()

	titokRouter := group.group.Group("/api/v1/parse/douyin")

	{
		titokRouter.GET("/get", tikTokController.GetTikTokVideoInfo)

		titokRouter.GET("/user", tikTokController.GetUserVideos)

		titokRouter.GET("/cookie", tikTokController.ReplaceTilTokCookie)
	}

	twitterController := controller.NewTwitterController()

	twitterRouter := group.group.Group("/api/v1/parse/twitter")

	{

		twitterRouter.GET("/info", twitterController.GetTwitterInfo)

		twitterRouter.GET("/token", twitterController.UpdateToken)

		twitterRouter.GET("/guest", twitterController.UpdateGuestToken)
	}

	youtubeController := controller.NewYouTuBeController()

	youtubeRouter := group.group.Group("/api/v1/parse/youtube")

	{
		youtubeRouter.GET("/video/info", youtubeController.GetYouTuBeVideoInfo)

		youtubeRouter.GET("/video/list", youtubeController.GetYouTuBeVideoInfoPlayList)
	}

	bilibiliControlelr := controller.NewBiliBiliController()

	group.group.GET("/api/v1/parse/bili/get", bilibiliControlelr.GetVideoInfo)
}

func (group globalRouter) InitRouter() {
	group.addUserRouter()
	group.uploadRouter()
	group.blogRouter()
	group.categoryRouter()
	group.tagRouter()
	group.topicRouter()
	group.parseRouter()
	group.otherRouter()
	group.fileRouter()
}
