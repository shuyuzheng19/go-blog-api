package repository

import (
	"gin-demo/models"
	"gin-demo/vo"
)

type TagRepository interface {
	FindAllSimpleTag() (tags []vo.TagVo)
	FindTagById(id int) (tag vo.TagVo)
	FindBlogByTagId(tid int, page int) (blogs []models.Blog, count int64)
	AddTag(tag models.Tag) models.Tag
}
