package repository

import (
	"gin-demo/models"
	"gin-demo/vo"
)

type CategoryRepository interface {
	FindAllCategoryList() (list []vo.CategoryVo)
	AddCategory(category models.Category) models.Category
}
