package controllers

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"vs-blog-api/common"
	"vs-blog-api/modal"
	"vs-blog-api/response"
	"vs-blog-api/service"
	"vs-blog-api/utils"
)

type fileController struct {
}

var fileService service.FileService

func NewFileController() *fileController {
	fileService = service.NewFileServiceImpl()
	return &fileController{}
}

func (*fileController) GetPublicFile(ctx *gin.Context) {

	var page = ctx.DefaultQuery("page", "1")

	var sort = ctx.DefaultQuery("sort", "date")

	var keyword = ctx.Query("keyword")

	var flagStr = ctx.DefaultQuery("flag", "false")

	flag, err := strconv.ParseBool(flagStr)

	if err != nil {
		flag = false
	}

	files := fileService.GetCurrentFiles(-1, utils.ToInt(page), sort, flag, keyword)

	ctx.JSON(200, response.SUCCESS(files))

}

func (*fileController) GetCurrentUserFile(ctx *gin.Context) {

	var page = ctx.DefaultQuery("page", "1")

	var sort = ctx.DefaultQuery("sort", "date")

	var keyword = ctx.Query("keyword")

	var flagStr = ctx.DefaultQuery("flag", "false")

	user := modal.GetUser(ctx)

	flag, err := strconv.ParseBool(flagStr)

	if err != nil {
		flag = false
	}

	files := fileService.GetCurrentFiles(user.Id, utils.ToInt(page), sort, flag, keyword)

	ctx.JSON(200, response.SUCCESS(files))

}

func (*fileController) DeleteCurrentUserFiles(ctx *gin.Context) {

	var ids []string

	ctx.ShouldBindJSON(&ids)

	user := modal.GetUser(ctx)

	if user.Role.Name == common.RoleUser {
		ctx.JSON(200, response.FAILURE(response.AUTHORIZATION, "你的权限不足!"))
		return
	}

	if user.Id <= 0 {
		ctx.JSON(200, response.FAILURE(response.AUTHORIZATION, "账号不存在!"))
		return
	}

	fileService.DeleteUserFiles(user.Id, ids)

	ctx.JSON(200, response.OK_RESULT)

}

func (*fileController) DeleteFiles(ctx *gin.Context) {

	var ids []string

	ctx.ShouldBindJSON(&ids)

	user := modal.GetUser(ctx)

	if user.Role.Name != common.RoleSuperAdmin {
		ctx.JSON(200, response.FAILURE(response.AUTHORIZATION, "只有超级管理员才有权限删除!"))
		return
	}

	fileService.DeleteFiles(ids)

	ctx.JSON(200, response.OK_RESULT)
}
