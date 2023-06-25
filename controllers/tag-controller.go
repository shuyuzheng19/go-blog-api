package controllers

import (
	"gin-demo/common"
	"gin-demo/myerr"
	"gin-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TagController struct {
	service service.TagService
}

func (tag TagController) RandomController(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, common.Success(tag.service.RandomTag()))
}

func (tag TagController) GetTagList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, common.Success(tag.service.GetAllTag()))
}

func (tag TagController) AddTag(ctx *gin.Context) {

	var name = ctx.Query("name")

	if name == "" {
		myerr.PanicError(common.TAG_NAME_EMPTY_ERROR)
	}

	var tagVo = tag.service.AddTag(name)

	ctx.JSON(http.StatusOK, common.Success(tagVo))

}

func (tag TagController) GetTagById(ctx *gin.Context) {
	var tid, err = strconv.Atoi(ctx.Param("tid"))

	if err != nil || tid <= 0 {
		ctx.JSON(http.StatusOK, common.TAG_ID_ERROR)
		return
	}

	ctx.JSON(http.StatusOK, common.Success(tag.service.GetTag(tid)))
}

func (tag TagController) GetBlogByTagId(ctx *gin.Context) {
	var tidStr = ctx.Param("tid")

	var tid, err = strconv.Atoi(tidStr)

	if err != nil || tid <= 0 {
		ctx.JSON(http.StatusOK, common.TAG_ID_ERROR)
		return
	}

	var page, _ = strconv.Atoi(ctx.DefaultQuery("page", "1"))

	var pageInfo = tag.service.GetTagBlogByPage(page, tid)

	ctx.JSON(http.StatusOK, common.Success(pageInfo))

}

func NewTagController(service service.TagService) TagController {
	return TagController{service: service}
}
