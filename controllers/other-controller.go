package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"vs-blog-api/common"
	"vs-blog-api/config"
	"vs-blog-api/manager"
	"vs-blog-api/modal"
	"vs-blog-api/response"
)

type otherController struct {
}

var musicManager manager.MusicPlayListManager

func NewOtherController() *otherController {
	musicManager = *manager.NewMusicPlayListManager()
	return &otherController{}
}

func (*otherController) InitMusicCloud(ctx *gin.Context) {

	var mid = ctx.Param("mid")

	if mid == "" {
		ctx.JSON(200, response.FAILURE(response.ParamsError, "网易云播放列表ID为空"))
		return
	}

	musicManager.SetMusicCloudPlayList(mid)

	ctx.JSON(200, response.OK_RESULT)

}

func (*otherController) GetTimeLines(ctx *gin.Context) {
	var timelines = make([]modal.TimeLine, 0)

	err := config.DB.Model(&modal.TimeLine{}).Find(&timelines).Error

	if err != nil {
		ctx.JSON(200, response.SUCCESS(timelines))
		return
	}

	ctx.JSON(200, response.SUCCESS(timelines))
}

func (*otherController) GetRecommend(ctx *gin.Context) {
	blogsStr := config.Redis.Get(common.RECOMMEND).Val()
	if blogsStr != "" {
		var blogs []modal.RecommendVo
		err := json.Unmarshal([]byte(blogsStr), &blogs)
		if err == nil {
			ctx.JSON(200, response.SUCCESS(blogs))
			return
		}
	}
	ctx.JSON(200, response.SUCCESS(make([]modal.RecommendVo, 0)))
}

func (*otherController) InitRecommend(ctx *gin.Context) {
	var ids []int

	ctx.ShouldBindJSON(&ids)

	if len(ids) > 4 {
		ids = ids[0:3]
	}

	var simpleBlogs []modal.RecommendVo

	err := config.DB.Model(&modal.Blog{}).Find(&simpleBlogs, "id in ?", ids).Error

	if err != nil {
		ctx.JSON(200, response.FAIL)
		return
	}

	marshal, err := json.Marshal(&simpleBlogs)

	if err != nil {
		ctx.JSON(200, response.FAIL)
		return
	}

	config.Redis.Set(common.RECOMMEND, marshal, -1)
}

func (*otherController) GetMusicPlayList(ctx *gin.Context) {

	musics := musicManager.GetMusicPlayList()

	ctx.JSON(200, response.SUCCESS(musics))
}
