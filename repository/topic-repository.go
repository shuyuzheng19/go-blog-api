package repository

import "vs-blog-api/modal"

type TopicRepository interface {
	FindAllTopicByPage(page int) (err error, topics []modal.Topic)
	GetTopicBlog(page int, id int) (err error, blogs []modal.Blog, count int64)
	FindById(id int) (err error, topic modal.Topic)
	FindTopicByUserId(userId int) (err error, topics []modal.Topic)
	SaveTopic(topic modal.Topic) (error, modal.Topic)
	GetAllSimpleTopic() (err error, topics []modal.SimpleTopicVo)
	GetTopicByIdAllBlog(id int) (err error, blogs []modal.Blog)
}
