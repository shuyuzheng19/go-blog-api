package service

import (
	"gin-demo/common"
	"gin-demo/config"
	"gin-demo/models"
	"gin-demo/myerr"
	"gin-demo/repository"
	"gin-demo/response"
	"gin-demo/utils"
	"gin-demo/vo"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"path/filepath"
	"strconv"
	"time"
)

type FileServiceImpl struct {
	config     config.UploadConfig
	repository repository.FileRepository
}

func (f FileServiceImpl) GetFileInfos(uid int, page int, keyword string, sort string) response.PageInfo {

	if sort != "size" {
		sort = "date"
	}

	var result, count = f.repository.FindFileByPage(uid, page, keyword, sort)

	var files = make([]vo.FileVo, 0)

	for _, info := range result {
		files = append(files, info.ToVo())
	}

	var pageInfo = response.PageInfo{
		Page:  page,
		Size:  common.FILE_PAGE_COUNT,
		Total: count,
		Data:  files,
	}

	return pageInfo

}

func (f FileServiceImpl) UploadFile(ctx *gin.Context, dir string, userId int) (urls []string) {

	var files, err = ctx.MultipartForm()

	if err != nil {
		ctx.JSON(http.StatusOK, common.GET_UPLOAD_FILE_ERROR)
		return
	}

	for _, file := range files.File["files"] {

		var fileName = file.Filename

		var flag = dir == common.FILES

		var size = file.Size

		if !flag {
			if !utils.IsImageFile(fileName) {
				myerr.PanicError(common.NOT_IAMGE_FILE)
			}
			if size > f.config.MaxImageSize {
				myerr.PanicError(common.MAX_IMAGE_SIZE_ERROR)
			}
		} else {
			if size > f.config.MaxFileSize {
				myerr.PanicError(common.MAX_IMAGE_SIZE_ERROR)
			}
		}

		var fe, err = file.Open()

		if err != nil {
			myerr.PanicError(common.OPEN_FILE_ERROR)
		}

		var md5 = utils.GetFileMd5(fe)

		var md5Url = f.repository.FindMD5(md5)

		var exists = md5Url != ""

		var flag2 = dir == common.AVATAR

		if exists && flag2 {
			urls = append(urls, md5Url)
			continue
		}

		var suffix = filepath.Ext(fileName)

		var newFileName = uuid.NewString() + suffix

		var filePath = f.config.Path + "/" + dir + "/" + newFileName

		var url string

		if !exists {

			var uploadErr = ctx.SaveUploadedFile(file, filePath)

			if uploadErr != nil {
				myerr.PanicError(common.UPLOAD_FILE_ERROR)
			} else {
				url = f.config.Uri + dir + "/" + newFileName
			}

		} else {
			url = md5Url
		}

		if flag {
			var public, err = strconv.ParseBool(ctx.Query("isPublic"))

			if err != nil {
				public = false
			}

			var fileInfo = models.FileInfo{
				UserID:       userId,
				OldName:      fileName,
				NewName:      newFileName,
				Date:         time.Now(),
				Size:         size,
				Suffix:       suffix,
				URL:          url,
				AbsolutePath: filePath,
				IsPublic:     public,
				FileMd5: models.FileMd5{
					Md5: md5,
					Url: url,
				},
			}

			var err3 = f.repository.SaveFileInfo(fileInfo)

			if err3 != nil {
				myerr.PanicError(common.UPLOAD_FILE_ERROR)
			}
		} else {
			if !exists {
				f.repository.SaveMd5(models.FileMd5{
					Md5: md5,
					Url: url,
				})
			}
		}
		urls = append(urls, url)
	}

	return urls
}

func NewFileService() FileService {
	var uploadConfig = config.GetUploadConfig()
	var repository = repository.NewFileRepository(config.DB)
	return FileServiceImpl{config: uploadConfig, repository: repository}
}
