package service

import "gin-demo/vo"

type CategoryService interface {
	GetAllCategoryListForCache() (list []vo.CategoryVo)
	GetAllCategoryListForDB() []vo.CategoryVo
	AddCategory(name string) vo.CategoryVo
}
