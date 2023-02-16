package service

import (
	"vs-blog-api/modal"
	"vs-blog-api/response"
)

type TopicService interface {
	GetTopicByPage(page int) []modal.TopicVo
	GetTopicBlog(page int, id int) response.PageInfoResponse
	FindTopicById(id int) modal.TopicVo
	GetUserTopics(userId int) []modal.TopicVo
	AddTopic(userId int, name string, cover string, desc string) modal.SimpleTopicVo
	GetAllTopic() []modal.SimpleTopicVo
	GetTopicByIdAllBlogs(topicId int) []modal.SimpleBlog
}
