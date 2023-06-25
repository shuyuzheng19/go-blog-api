package repository

import (
	"gin-demo/models"
	"gin-demo/vo"
	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	db *gorm.DB
}

func (c CategoryRepositoryImpl) AddCategory(category models.Category) models.Category {
	c.db.Model(&models.Category{}).Create(&category)
	return category
}

func (c CategoryRepositoryImpl) FindAllCategoryList() (list []vo.CategoryVo) {
	c.db.Model(&models.Category{}).Select("id,name").Order("create_at desc").Find(&list)
	return list
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return CategoryRepositoryImpl{db: db}
}
