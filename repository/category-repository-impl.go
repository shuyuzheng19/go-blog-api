package repository

import (
	"vs-blog-api/config"
	"vs-blog-api/modal"
)

type CategoryRepositoryImpl struct {
}

func (c CategoryRepositoryImpl) AddCategory(category modal.Category) (error, modal.Category) {

	err := config.DB.Model(&modal.Category{}).Create(&category).Error

	return err, category
}

func NewCategoryRepository() CategoryRepository {
	return CategoryRepositoryImpl{}
}

func (c CategoryRepositoryImpl) FindAllCategory() (err error, categories []modal.Category) {

	err = config.DB.Model(&categories).Order("create_time desc").Find(&categories).Error

	return err, categories
}
