package controllers

import (
	"gin-demo/common"
	"gin-demo/request"
	"gin-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TopicController struct {
	service service.TopicService
}

func (t TopicController) GetTopicByPage(ctx *gin.Context) {
	var page, _ = strconv.Atoi(ctx.DefaultQuery("page", "1"))

	var pageInfo = t.service.GetTopicByPage(page)

	ctx.JSON(http.StatusOK, common.Success(pageInfo))
}

func (t TopicController) AddTopic(ctx *gin.Context) {

	var topicRequest request.TopicRequest

	ctx.ShouldBindJSON(&topicRequest)

	var topic = t.service.AddTopic(GetUser(ctx).Id, topicRequest)

	ctx.JSON(http.StatusOK, common.Success(topic))

}

func (t TopicController) GetCurrentTopics(ctx *gin.Context) {
	var user = GetUser(ctx)

	ctx.JSON(http.StatusOK, common.Success(t.service.GetUserTopic(user.Id)))
}

func (t TopicController) GetTopicById(ctx *gin.Context) {
	var tid, err = strconv.Atoi(ctx.Param("tid"))

	if err != nil || tid <= 0 {
		ctx.JSON(http.StatusOK, common.TOPIC_ID_ERROR)
		return
	}

	var topic = t.service.GetTopicById(tid)

	ctx.JSON(http.StatusOK, common.Success(topic))
}

func (t TopicController) GetUserTopic(ctx *gin.Context) {

	var uid, err = strconv.Atoi(ctx.Param("uid"))

	if err != nil || uid <= 0 {
		ctx.JSON(http.StatusOK, common.USER_ID_ERROR)
		return
	}

	var topics = t.service.GetUserTopics(uid)

	ctx.JSON(http.StatusOK, common.Success(topics))

}

func (t TopicController) GetTopicBlogList(ctx *gin.Context) {

	var tid, err = strconv.Atoi(ctx.Param("tid"))

	if err != nil || tid <= 0 {
		ctx.JSON(http.StatusOK, common.TOPIC_ID_ERROR)
		return
	}

	var blogs = t.service.TopicBlogList(tid)

	ctx.JSON(http.StatusOK, common.Success(blogs))
}

func (t TopicController) GetTopicBlogByPage(ctx *gin.Context) {

	var tid, err = strconv.Atoi(ctx.Param("tid"))

	if err != nil || tid <= 0 {
		ctx.JSON(http.StatusOK, common.TOPIC_ID_ERROR)
		return
	}

	var page, _ = strconv.Atoi(ctx.DefaultQuery("page", "1"))

	var pageInfo = t.service.GetTopicBlogByPage(tid, page)

	ctx.JSON(http.StatusOK, common.Success(pageInfo))
}

func NewTopicController(service service.TopicService) TopicController {
	return TopicController{service: service}
}
