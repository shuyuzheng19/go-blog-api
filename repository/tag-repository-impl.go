package repository

import (
	"vs-blog-api/common"
	"vs-blog-api/config"
	"vs-blog-api/modal"
)

type TagRepositoryImpl struct {
}

func (t TagRepositoryImpl) AddTag(tagInfo modal.Tag) (error, modal.Tag) {

	err := config.DB.Model(&modal.Tag{}).Create(&tagInfo).Error

	return err, tagInfo
}

func (t TagRepositoryImpl) FindById(id int) (err error, tag modal.Tag) {
	err = config.DB.Model(&modal.Tag{}).First(&tag, "id = ?", id).Error

	return err, tag
}

func (t TagRepositoryImpl) FindTagIdBlog(id int, page int) (err error, blogs []modal.Blog, count int64) {

	var sql = "select " + SELECT_JOIN + " from blogs a join blogs_tags b on a.id = b.blog_id where b.tag_id = ? order by create_time desc offset ? limit ?"

	var countSql = "select count(id) from blogs a join blogs_tags b on a.id = b.blog_id where b.tag_id = ?"

	err = config.DB.Model(&modal.Blog{}).Select(SELECT).Raw(sql, id, (page-1)*common.PageSize, common.PageSize).Preload("User").Preload("Category").Find(&blogs).Raw(countSql, id).First(&count).Error

	return err, blogs, count
}

func NewTagRepository() TagRepository {
	return TagRepositoryImpl{}
}

func (t TagRepositoryImpl) FindAllTag() []modal.Tag {

	var tags []modal.Tag

	config.DB.Model(&modal.Tag{}).Find(&tags)

	return tags
}
