package service

import (
	"gin-demo/request"
	"gin-demo/response"
	"gin-demo/vo"
)

type TopicService interface {
	GetTopicByPage(page int) response.PageInfo
	GetTopicBlogByPage(tid int, page int) response.PageInfo
	GetTopicById(tid int) vo.SimpleTopicVo
	GetUserTopics(uid int) []vo.SimpleTopicVo
	TopicBlogList(tid int) []vo.SimpleBlogVo
	GetUserTopic(uid int) []vo.SimpleTopicVo
	AddTopic(uid int, request request.TopicRequest) vo.SimpleTopicVo
}
