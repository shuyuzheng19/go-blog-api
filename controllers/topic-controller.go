package controllers

import (
	"github.com/gin-gonic/gin"
	"vs-blog-api/modal"
	"vs-blog-api/response"
	"vs-blog-api/service"
	"vs-blog-api/utils"
)

type topicController struct {
}

func NewTopicController() *topicController {
	topicService = service.NewTopicService()
	return &topicController{}
}

var topicService service.TopicService

func (*topicController) GetTopics(ctx *gin.Context) {

	var page = ctx.DefaultQuery("page", "1")

	topics := topicService.GetTopicByPage(utils.ToInt(page))

	ctx.JSON(200, response.SUCCESS(topics))
}

func (*topicController) GetAllTopics(ctx *gin.Context) {

	topics := topicService.GetAllTopic()

	ctx.JSON(200, response.SUCCESS(topics))

}

func (*topicController) GetTopicIdBlogs(ctx *gin.Context) {

	var id = ctx.Param("id")

	pageInfo := topicService.GetTopicByIdAllBlogs(utils.ToInt(id))

	ctx.JSON(200, response.SUCCESS(pageInfo))

}

func (*topicController) AddTopic(ctx *gin.Context) {

	var maps map[string]string

	ctx.ShouldBindJSON(&maps)

	var userId = modal.GetUser(ctx).Id

	topic := topicService.AddTopic(userId, maps["name"], maps["cover"], maps["desc"])

	ctx.JSON(200, response.SUCCESS(topic))
}

func (*topicController) GetTopicBlog(ctx *gin.Context) {

	var page = ctx.DefaultQuery("page", "1")

	var id = ctx.Param("id")

	if id == "" {
		ctx.JSON(200, response.FAILURE(response.ParamsError, "缺少专题ID"))
	}

	pageInfo := topicService.GetTopicBlog(utils.ToInt(page), utils.ToInt(id))

	ctx.JSON(200, response.SUCCESS(pageInfo))
}

func (*topicController) GetById(ctx *gin.Context) {

	var id = ctx.Param("id")

	if id == "" {
		ctx.JSON(200, response.FAILURE(response.ParamsError, "缺少专题ID"))
	}

	topic := topicService.FindTopicById(utils.ToInt(id))

	ctx.JSON(200, response.SUCCESS(topic))
}

func (*topicController) GetUserTopic(ctx *gin.Context) {

	var id = ctx.Param("userId")

	topics := topicService.GetUserTopics(utils.ToInt(id))

	ctx.JSON(200, response.SUCCESS(topics))
}
