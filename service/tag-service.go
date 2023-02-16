package service

import (
	"vs-blog-api/modal"
	"vs-blog-api/response"
)

type TagService interface {
	GetRandomTag() []modal.TagVo
	GetTagBlog(id int, page int) response.PageInfoResponse
	FindByIdTag(id int) modal.TagVo
	GetAllTag() []modal.TagVo
	SaveTag(name string) modal.TagVo
}
