package controllers

import (
	"gin-demo/common"
	"gin-demo/request"
	"gin-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type BlogController struct {
	service service.BlogService
}

func (b BlogController) FindBlogPage(ctx *gin.Context) {

	var page, _ = strconv.Atoi(ctx.DefaultQuery("page", "1"))

	var cid, _ = strconv.Atoi(ctx.DefaultQuery("cid", "-1"))

	var pageRequest = request.BlogPageRequest{
		Page: page,
		Cid:  cid,
		Sort: request.Sort(ctx.Query("sort")),
	}

	var pageInfo = b.service.FindBlogByCidPage(pageRequest)

	ctx.JSON(http.StatusOK, common.Success(pageInfo))
}

func (b BlogController) GetHotBlog(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, common.Success(b.service.GetHotBlog()))
}

func (b BlogController) SetRecommendBlog(ctx *gin.Context) {
	var ids []int

	ctx.ShouldBindJSON(&ids)

	b.service.SetRecommend(ids)

	ctx.JSON(http.StatusOK, common.OK())
}

func (b BlogController) GetRecommendBlog(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, b.service.GetRecommend())
}

func (b BlogController) GetRandomBlog(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, common.Success(b.service.RandomBlogs()))
}

func NewBlogController(service service.BlogService) BlogController {
	return BlogController{service: service}
}
