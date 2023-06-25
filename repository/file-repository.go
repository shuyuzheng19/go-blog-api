package repository

import "gin-demo/models"

type FileRepository interface {
	FindMD5(md5 string) string
	SaveMd5(md5 models.FileMd5) error
	SaveFileInfo(fileInfo models.FileInfo) error
	FindFileByPage(uid int, page int, keyword string, sort string) (files []models.FileInfo, count int64)
}
