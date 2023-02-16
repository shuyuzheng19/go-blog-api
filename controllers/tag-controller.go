package controllers

import (
	"github.com/gin-gonic/gin"
	"vs-blog-api/response"
	"vs-blog-api/service"
	"vs-blog-api/utils"
)

type tagController struct {
}

func NewTagController() *tagController {
	tagService = service.NewTagService()
	return &tagController{}
}

var tagService service.TagService

func (*tagController) RandomTag(ctx *gin.Context) {

	randomTags := tagService.GetRandomTag()

	ctx.JSON(200, response.SUCCESS(randomTags))

}

func (*tagController) GetAllTags(ctx *gin.Context) {

	allTags := tagService.GetAllTag()

	ctx.JSON(200, response.SUCCESS(allTags))

}

func (*tagController) SaveTag(ctx *gin.Context) {

	var name = ctx.Query("name")

	tag := tagService.SaveTag(name)

	ctx.JSON(200, response.SUCCESS(tag))

}

func (*tagController) GetTagBlog(ctx *gin.Context) {

	var id = ctx.Param("id")

	var page = ctx.DefaultQuery("page", "1")

	pageInfo := tagService.GetTagBlog(utils.ToInt(id), utils.ToInt(page))

	ctx.JSON(200, response.SUCCESS(pageInfo))

}

func (*tagController) GetTag(ctx *gin.Context) {

	var id = ctx.Param("id")

	tag := tagService.FindByIdTag(utils.ToInt(id))

	ctx.JSON(200, response.SUCCESS(tag))

}
