package service

import (
	"gin-demo/response"
	"gin-demo/vo"
)

type TagService interface {
	GetAllTag() (tags []vo.TagVo)
	RandomTag() (tags []vo.TagVo)
	GetTagBlogByPage(page int, tid int) response.PageInfo
	GetTag(tid int) (tag vo.TagVo)
	AddTag(name string) vo.TagVo
}
