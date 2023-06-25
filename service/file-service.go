package service

import (
	"gin-demo/response"
	"github.com/gin-gonic/gin"
)

type FileService interface {
	UploadFile(ctx *gin.Context, dir string, userId int) (urls []string)
	GetFileInfos(uid int, page int, keyword string, sort string) response.PageInfo
}
