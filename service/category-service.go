package service

import "vs-blog-api/modal"

type CategoryService interface {
	FindAllCategory() []modal.CategoryVo
	SaveCategory(name string) *modal.CategoryVo
}
