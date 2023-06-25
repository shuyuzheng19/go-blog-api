package cache

import (
	"gin-demo/models"
	"gin-demo/response"
	"gin-demo/vo"
)

type TagCache interface {
	GetRandomTags() []string
	SaveAllTags(tags []vo.TagVo)
	GetTagInfo(id string) string
	SaveTagInfo(tag vo.TagVo) error
	SaveFirstPageBlog(tid string, pageInfo response.PageInfo) error
	GetFirstPageBlog(tid string) string
	AddTag(tag models.Tag) error
	RemoveKey() error
}
