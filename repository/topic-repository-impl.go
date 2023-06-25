package repository

import (
	"gin-demo/common"
	"gin-demo/models"
	"gin-demo/request"
	"gin-demo/vo"
	"gorm.io/gorm"
)

type TopicRepositoryImpl struct {
	db *gorm.DB
}

func (t TopicRepositoryImpl) AddTopic(topic models.Topic) models.Topic {
	t.db.Model(&models.Topic{}).Create(&topic)
	return topic
}

func (t TopicRepositoryImpl) GetUserTopic(uid int) (topics []vo.SimpleTopicVo) {
	t.db.Model(&models.Topic{}).Select("id,name,user_id").Where("user_id = ?", uid).Find(&topics)
	return topics
}

func (t TopicRepositoryImpl) FindAllTopicList(tid int) (topics []vo.SimpleBlogVo) {
	t.db.Model(&models.Blog{}).Select("id,title,topic_id").Where("topic_id = ?", tid).Find(&topics)
	return topics
}

func (t TopicRepositoryImpl) FindTopicByUser(uid int) (topics []vo.SimpleTopicVo) {
	t.db.Model(&models.Topic{}).Select("id,name,description,cover_image,user_id").Where("user_id = ?", uid).Find(&topics)
	return topics
}

func (t TopicRepositoryImpl) FindTopicById(tid int) (topic vo.SimpleTopicVo) {
	t.db.Model(&models.Topic{}).Select("id,name").First(&topic, "id = ?", tid)
	return topic
}

func (t TopicRepositoryImpl) FindTopicBlogByPage(tid int, page int) (blogs []models.Blog, count int64) {
	var query = t.db.Model(&models.Blog{}).Select("id,title,description,cover_image,create_at,user_id,topic_id").Where("topic_id = ?", tid)

	if err := query.Count(&count).Error; err != nil || count == 0 {
		return make([]models.Blog, 0), 0
	}

	query.Offset((page - 1) * common.PAGE_SIZE).Limit(common.PAGE_SIZE).Preload("User").Order(request.Sort("BACK").String()).Find(&blogs)

	return blogs, count
}

func (t TopicRepositoryImpl) FindTopicByPage(page int) (topics []models.Topic, count int64) {
	var query = t.db.Model(&models.Topic{}).Offset((page - 1) * common.TOPIC_PAGE_SIZE).Limit(common.TOPIC_PAGE_SIZE)

	if err := query.Count(&count).Error; err != nil || count == 0 {
		return make([]models.Topic, 0), 0
	}

	query.Preload("User").Order(request.Sort("CREATE").String()).Find(&topics)

	return topics, count
}

func NewTopicRepository(db *gorm.DB) TopicRepository {
	return TopicRepositoryImpl{db: db}
}
