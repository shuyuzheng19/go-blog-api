package repository

import (
	"gin-demo/common"
	"gin-demo/models"
	"gin-demo/vo"
	"gorm.io/gorm"
)

type TagRepositoryImpl struct {
	db *gorm.DB
}

func (t TagRepositoryImpl) AddTag(tag models.Tag) models.Tag {
	t.db.Model(&models.Tag{}).Create(&tag)
	return tag
}

func (t TagRepositoryImpl) FindTagById(id int) (tag vo.TagVo) {

	t.db.Model(&models.Tag{}).Select("id,name").First(&tag, "id = ?", id)

	return tag
}

func (t TagRepositoryImpl) FindBlogByTagId(tid int, page int) (blogs []models.Blog, count int64) {
	var query = t.db.Model(&models.Blog{}).Select("id,title,description,create_at,category_id,user_id").Joins("JOIN blogs_tags tag on tag.blog_id = id").Where("tag.tag_id = ?", tid)

	if err := query.Count(&count).Error; err != nil || count == 0 {
		return make([]models.Blog, 0), 0
	}

	query.Offset((page - 1) * common.PAGE_SIZE).Limit(common.PAGE_SIZE).Preload("Category").Preload("User").Find(&blogs)

	return blogs, count
}

func (t TagRepositoryImpl) FindAllSimpleTag() (tags []vo.TagVo) {
	t.db.Model(&models.Tag{}).Select("id,name").Find(&tags)
	return tags
}

func NewTagRepository(db *gorm.DB) TagRepository {
	return TagRepositoryImpl{db: db}
}
