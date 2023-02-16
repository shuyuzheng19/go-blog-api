package controllers

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"vs-blog-api/common"
	"vs-blog-api/manager"
	"vs-blog-api/response"
)

type uploadController struct {
}

var uploadManager manager.UploadManager

func NewUploadController() uploadController {

	uploadManager = manager.NewUploadManager()

	return uploadController{}
}

func (*uploadController) UploadAvatar(ctx *gin.Context) {
	uploadManager.UploadAvatar(ctx, "/avatars", common.ImageTypes, common.MB*5, "图片大小不能大于5MB", "这不是一个图片文件,请核对")
}

func (*uploadController) ParseMarkDownFile(ctx *gin.Context) {
	file, err := ctx.FormFile("file")

	size := file.Size

	if size > common.MB*10 {
		ctx.JSON(http.StatusOK, response.FAILURE(response.ParamsError, "md文件大小不能超过10MB"))
		return
	}

	if err != nil {
		ctx.JSON(http.StatusOK, response.FAILURE(response.ERROR, "读取文件失败"))
		return
	}

	open, err2 := file.Open()

	if err2 != nil {
		ctx.JSON(http.StatusOK, response.FAILURE(response.ERROR, "读取文件失败"))
		return
	}

	result, err3 := io.ReadAll(open)

	if err3 != nil {
		ctx.JSON(http.StatusOK, response.FAILURE(response.ERROR, "解析文件失败"))
		return
	}

	ctx.JSON(200, response.SUCCESS(string(result)))

}

func (*uploadController) UploadOther(ctx *gin.Context) {
	uploadManager.Upload(ctx)
}
