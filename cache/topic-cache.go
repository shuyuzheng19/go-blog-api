package cache

import (
	"gin-demo/response"
	"gin-demo/vo"
)

type TopicCache interface {
	SaveFirstPageTopic(pageInfo response.PageInfo) error
	GetFirstPageTopic() string
	SetTopicToMap(topic vo.SimpleTopicVo) error
	GetTopicFromMap(id string) string
	SaveFirstPageBlog(tid string, pageInfo response.PageInfo) error
	GetFirstPageBlog(tid string) string
	RemoveTopicPageKey() error
}
