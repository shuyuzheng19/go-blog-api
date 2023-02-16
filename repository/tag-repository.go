package repository

import "vs-blog-api/modal"

type TagRepository interface {
	FindAllTag() []modal.Tag
	FindTagIdBlog(id int, page int) (err error, blogs []modal.Blog, count int64)
	FindById(id int) (err error, tag modal.Tag)
	AddTag(tagInfo modal.Tag) (error, modal.Tag)
}
