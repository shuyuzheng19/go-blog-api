package repository

import (
	"vs-blog-api/common"
	"vs-blog-api/config"
	"vs-blog-api/modal"
)

type TopicRepositoryImpl struct {
}

func (t TopicRepositoryImpl) GetTopicByIdAllBlog(id int) (err error, blogs []modal.Blog) {
	err = config.DB.Model(&modal.Blog{}).Order("create_time asc").Find(&blogs, "topic_id = ?", id).Error

	return err, blogs
}

func (t TopicRepositoryImpl) GetAllSimpleTopic() (err error, topics []modal.SimpleTopicVo) {
	err = config.DB.Model(&modal.Topic{}).Find(&topics).Error

	return err, topics
}

func (t TopicRepositoryImpl) SaveTopic(topic modal.Topic) (error, modal.Topic) {
	err := config.DB.Model(&modal.Topic{}).Create(&topic).Error

	return err, topic
}

func (t TopicRepositoryImpl) FindTopicByUserId(userId int) (err error, topics []modal.Topic) {

	err = config.DB.Model(&modal.Topic{}).Find(&topics, "user_id = ?", userId).Error

	return err, topics
}

func (t TopicRepositoryImpl) FindById(id int) (err error, topic modal.Topic) {

	err = config.DB.Model(&topic).First(&topic, "id = ?", id).Error

	return err, topic
}

func (t TopicRepositoryImpl) GetTopicBlog(page int, id int) (err error, blogs []modal.Blog, count int64) {

	preload := config.DB.Model(&modal.Blog{}).Select(SELECT).Preload("User").Preload("Topic").Where("topic_id = ?", id)

	err = preload.Offset((page - 1) * common.PageSize).Limit(common.PageSize).Order("create_time asc").Find(&blogs).Offset(-1).Limit(-1).Count(&count).Error

	return err, blogs, count
}

func NewTopicRepository() TopicRepository {
	return &TopicRepositoryImpl{}
}

func (t TopicRepositoryImpl) FindAllTopicByPage(page int) (err error, topics []modal.Topic) {

	err = config.DB.Model(&modal.Topic{}).Preload("User").Offset((page - 1) * common.TopicSize).Limit(common.TopicSize).Order("create_time desc").Find(&topics).Error

	return err, topics
}
