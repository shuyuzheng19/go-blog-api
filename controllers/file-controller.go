package controllers

import (
	"gin-demo/common"
	"gin-demo/models"
	"gin-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type FileController struct {
	service service.FileService
}

func (f FileController) UploadAvatar(ctx *gin.Context) {

	var urls = f.service.UploadFile(ctx, common.AVATAR, -1)

	ctx.JSON(http.StatusOK, common.Success(urls))
}

func (f FileController) GetPublicFile(ctx *gin.Context) {

	var page, _ = strconv.Atoi(ctx.DefaultQuery("page", "1"))

	var sort = ctx.Query("sort")

	var keyword = ctx.Query("keyword")

	var pageInfo = f.service.GetFileInfos(0, page, keyword, sort)

	ctx.JSON(http.StatusOK, common.Success(pageInfo))

}

func (f FileController) GetCurrentUserFile(ctx *gin.Context) {

	var user, exists = ctx.Get("user")

	if !exists {
		ctx.JSON(http.StatusOK, common.USER_NOT_FOUNT)
		return
	}

	var page, _ = strconv.Atoi(ctx.DefaultQuery("page", "1"))

	var sort = ctx.Query("sort")

	var keyword = ctx.Query("keyword")

	var pageInfo = f.service.GetFileInfos(user.(models.User).Id, page, keyword, sort)

	ctx.JSON(http.StatusOK, common.Success(pageInfo))

}

func (f FileController) UploadImage(ctx *gin.Context) {

	var user, exists = ctx.Get("user")

	if !exists {
		ctx.JSON(http.StatusOK, common.USER_NOT_FOUNT)
		return
	}

	var urls = f.service.UploadFile(ctx, common.IMAGES, user.(models.User).Id)

	ctx.JSON(http.StatusOK, common.Success(urls))
}

func (f FileController) UploadFile(ctx *gin.Context) {

	var user, exists = ctx.Get("user")

	if !exists {
		ctx.JSON(http.StatusOK, common.USER_NOT_FOUNT)
		return
	}

	var urls = f.service.UploadFile(ctx, common.FILES, user.(models.User).Id)

	ctx.JSON(http.StatusOK, common.Success(urls))
}

func NewFileController(service service.FileService) FileController {
	return FileController{service: service}
}
