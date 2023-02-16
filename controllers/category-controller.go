package controllers

import (
	"github.com/gin-gonic/gin"
	"vs-blog-api/response"
	"vs-blog-api/service"
)

type categoryController struct {
}

func NewCategoryController() *categoryController {
	categoryService = service.NewCategoryService()
	return &categoryController{}
}

var categoryService service.CategoryService

func (*categoryController) GetAllCategory(ctx *gin.Context) {

	categorys := categoryService.FindAllCategory()

	ctx.JSON(200, response.SUCCESS(categorys))

}

func (*categoryController) SaveCategory(ctx *gin.Context) {

	var name = ctx.Query("name")

	category := categoryService.SaveCategory(name)

	ctx.JSON(200, response.SUCCESS(category))

}
