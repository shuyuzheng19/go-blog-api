package repository

import "vs-blog-api/modal"

type CategoryRepository interface {
	FindAllCategory() (err error, categories []modal.Category)
	AddCategory(category modal.Category) (error, modal.Category)
}
