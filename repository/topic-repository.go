package repository

import (
	"gin-demo/models"
	"gin-demo/vo"
)

type TopicRepository interface {
	FindTopicByPage(page int) (topics []models.Topic, count int64)
	FindTopicBlogByPage(tid int, page int) (blogs []models.Blog, count int64)
	FindTopicById(tid int) (topic vo.SimpleTopicVo)
	FindTopicByUser(uid int) (topics []vo.SimpleTopicVo)
	FindAllTopicList(tid int) (blogs []vo.SimpleBlogVo)
	GetUserTopic(uid int) (topics []vo.SimpleTopicVo)
	AddTopic(topic models.Topic) models.Topic
}
