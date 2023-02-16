package manager

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"
	"vs-blog-api/common"
	"vs-blog-api/config"
	"vs-blog-api/modal"
	"vs-blog-api/response"
	"vs-blog-api/utils"
)

type UploadManager struct {
}

func NewUploadManager() UploadManager {
	return UploadManager{}
}

func (*UploadManager) UploadAvatar(ctx *gin.Context, pathName string, fileTypes []string, maxSize int64, sizeError string, typeError string) {

	file, err := ctx.FormFile("file")

	size := file.Size

	if size > maxSize {
		ctx.JSON(http.StatusOK, response.FAILURE(response.ParamsError, sizeError))
		return
	}

	if err != nil {
		ctx.JSON(http.StatusOK, response.FAILURE(response.ERROR, "读取文件失败"))
		return
	}

	filename := file.Filename

	if fileTypes != nil {
		if !utils.IsAssignFile(filename, fileTypes) {
			ctx.JSON(http.StatusOK, response.FAILURE(response.ParamsError, typeError))
			return
		}
	}

	var UUID = uuid.NewString()

	suffix := utils.GetFileNameExt(filename)

	var newFileName = UUID + "-" + strconv.Itoa(int(time.Now().UnixMilli())) + "." + suffix

	var filePath = common.UploadPath + pathName

	if _, err := os.Stat(filePath); err != nil {
		os.MkdirAll(filePath, os.ModePerm)
	}

	err = ctx.SaveUploadedFile(file, path.Join(filePath, newFileName))

	if err != nil {
		ctx.JSON(http.StatusOK, response.FAILURE(response.ERROR, "上传失败"))
		return
	}

	var url = "http://" + config.HostName + "/static" + pathName + "/" + newFileName

	ctx.JSON(http.StatusOK, response.SUCCESS(url))

}

func (*UploadManager) Upload(ctx *gin.Context) {

	user := modal.GetUser(ctx)

	var pathName = "/" + user.Username + "/files"

	file, err := ctx.FormFile("file")

	if err != nil {
		ctx.JSON(http.StatusOK, response.FAILURE(response.ERROR, "读取文件失败"))
		return
	}

	size := file.Size

	if size > int64(common.MaxUploadSize) {
		ctx.JSON(http.StatusOK, response.FAILURE(response.ParamsError, "文件最大大小不能超过 "+strconv.Itoa(int(common.MaxUploadSize/1024/1024/1024))+" GB"))
		return
	}

	hash := md5.New()

	open, _ := file.Open()

	io.Copy(hash, open)

	md5 := hex.EncodeToString(hash.Sum(nil))

	open.Close()

	query := ctx.PostForm("public")

	public, err2 := strconv.ParseBool(query)

	if err2 != nil {
		public = false
	}

	do, err3 := config.ES.Get().Index(common.FileIndex).Id(md5).Do(context.Background())

	if err3 == nil {

		var fileInfo modal.EsFile

		json.Unmarshal(do.Source, &fileInfo)

		if fileInfo.UserId != user.Id {

			fileInfo.UserId = user.Id

			fileInfo.Public = public

			fileInfo.First = false

			config.ES.Index().Index(common.FileIndex).BodyJson(fileInfo).Do(context.Background())

			ctx.JSON(http.StatusOK, response.SUCCESS(fileInfo.Url))

			return

		} else {
			if fileInfo.Public != public {

				fileInfo.Id = md5

				fileInfo.Public = public

				fileInfo.Date = time.Now()

				config.ES.Index().Index(common.FileIndex).Id(md5).BodyJson(fileInfo).Do(context.Background())

				ctx.JSON(http.StatusOK, response.SUCCESS(fileInfo.Url))

				return
			}
		}

	}

	filename := file.Filename

	suffix := utils.GetFileNameExt(filename)

	var newFileName = md5 + "." + suffix

	var filePath = common.UploadPath + pathName

	if _, err := os.Stat(filePath); err != nil {
		os.MkdirAll(filePath, os.ModePerm)
	}

	err = ctx.SaveUploadedFile(file, path.Join(filePath, newFileName))

	if err != nil {
		ctx.JSON(http.StatusOK, response.FAILURE(response.ERROR, "上传失败"))
		return
	}

	var url = "http://" + config.HostName + "/static" + pathName + "/" + newFileName

	fileInfo := modal.EsFile{
		Id:           md5,
		UserId:       user.Id,
		OldName:      filename,
		NewName:      newFileName,
		Date:         time.Now(),
		Size:         file.Size,
		Suffix:       suffix,
		Url:          url,
		AbsolutePath: filePath + "/" + newFileName,
		Path:         "/static" + pathName + "/" + newFileName,
		Public:       public,
		First:        true,
	}

	config.ES.Index().Index(common.FileIndex).Id(md5).BodyJson(fileInfo).Do(context.Background())

	ctx.JSON(http.StatusOK, response.SUCCESS(url))
}
