package cache

import (
	"gin-demo/models"
	"gin-demo/vo"
)

type CategoryCache interface {
	SaveCategoryList(list []vo.CategoryVo) error
	GetCategoryList() []string
	AddCategory(category models.Category) error
	RemoveKey() error
}
