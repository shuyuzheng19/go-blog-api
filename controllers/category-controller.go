package controllers

import (
	"gin-demo/common"
	"gin-demo/myerr"
	"gin-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CategoryController struct {
	service service.CategoryService
}

func (c CategoryController) GetCategoryListForDB(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, common.Success(c.service.GetAllCategoryListForCache()))
}

func (c CategoryController) AddCategory(ctx *gin.Context) {

	var name = ctx.Query("name")

	if name == "" {
		myerr.PanicError(common.CATEGORY_NAME_EMPTY_ERROR)
	}

	var categoryVo = c.service.AddCategory(name)

	ctx.JSON(http.StatusOK, common.Success(categoryVo))

}

func (c CategoryController) GetCategoryListForCache(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, common.Success(c.service.GetAllCategoryListForDB()))
}

func NewCategoryController(service service.CategoryService) CategoryController {
	return CategoryController{service: service}
}
