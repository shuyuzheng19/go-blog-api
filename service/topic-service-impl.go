package service

import (
	"time"
	"vs-blog-api/common"
	"vs-blog-api/modal"
	"vs-blog-api/repository"
	"vs-blog-api/response"
)

type TopicServiceImpl struct {
}

func (t TopicServiceImpl) GetTopicByIdAllBlogs(topicId int) []modal.SimpleBlog {
	if topicId <= 0 {
		panic(response.NewGlobalException(response.ParamsError, "非法的专题ID"))
	}

	err, blogs := topicRepository.GetTopicByIdAllBlog(topicId)

	if err != nil {
		return make([]modal.SimpleBlog, 0)
	}

	var blogVos []modal.SimpleBlog

	for _, blog := range blogs {
		blogVos = append(blogVos, modal.SimpleBlog{
			Id:    blog.Id,
			Title: blog.Title,
		})
	}

	return blogVos
}

func (t TopicServiceImpl) GetAllTopic() []modal.SimpleTopicVo {

	err, topics := topicRepository.GetAllSimpleTopic()

	if err != nil {
		panic(response.NewGlobalException(response.ERROR, "查询专题失败"))
	}

	return topics
}

func (t TopicServiceImpl) AddTopic(userId int, name string, cover string, desc string) modal.SimpleTopicVo {

	if userId <= 0 {
		panic(response.NewGlobalException(response.ParamsError, "非法的用户ID"))
	}

	if len(name) < 3 || len(name) > 20 {
		panic(response.NewGlobalException(response.ParamsError, "专题名称长度不能小于3个并且不能大于20个"))
	}

	if len(desc) < 3 || len(desc) > 100 {
		panic(response.NewGlobalException(response.ParamsError, "专题描述长度不能小于3个并且不能大于100个"))
	}

	if cover == "" {
		panic(response.NewGlobalException(response.ParamsError, "专题封面不能为空"))
	}

	var topic = modal.Topic{
		Id:          0,
		Name:        name,
		Description: desc,
		CoverImage:  cover,
		UserId:      userId,
		User:        modal.User{},
		CreateTime:  time.Now(),
	}

	err, m := topicRepository.SaveTopic(topic)

	if err != nil {
		panic(response.NewGlobalException(response.ERROR, "添加失败"))
	}

	return modal.SimpleTopicVo{
		Id:   m.Id,
		Name: m.Name,
	}

}

func (t TopicServiceImpl) GetUserTopics(userId int) []modal.TopicVo {
	if userId <= 0 {
		panic(response.NewGlobalException(response.ParamsError, "非法参数"))
	}

	err, topics := topicRepository.FindTopicByUserId(userId)

	if err != nil {
		return []modal.TopicVo{}
	}

	var topicVos []modal.TopicVo

	for _, topic := range topics {
		topicVos = append(topicVos, *topic.ToVo())
	}

	return topicVos

}

func (t TopicServiceImpl) FindTopicById(id int) modal.TopicVo {
	if id <= 0 {
		panic(response.NewGlobalException(response.ParamsError, "非法参数"))
	}

	err, topic := topicRepository.FindById(id)

	if err != nil {
		panic(response.NewGlobalException(response.NOTFOUND, "找不到该专题"))
	}

	return *topic.ToVo()

}

func (t TopicServiceImpl) GetTopicBlog(page int, id int) response.PageInfoResponse {

	if id <= 0 {
		panic(response.NewGlobalException(response.ParamsError, "非法ID"))
	}

	err, blogs, count := topicRepository.GetTopicBlog(page, id)

	if err != nil {
		return response.PageInfoResponse{}
	}

	var blogVos = BlogArrayToBlogVoArray(blogs)

	return response.PageInfoResponse{
		Page:  page,
		Size:  common.PageSize,
		Total: count,
		Data:  blogVos,
	}

}

func (t TopicServiceImpl) GetTopicByPage(page int) []modal.TopicVo {

	err, topics := topicRepository.FindAllTopicByPage(page)

	if err != nil {
		return []modal.TopicVo{}
	}

	var topicVos []modal.TopicVo

	for _, topic := range topics {
		topicVos = append(topicVos, *topic.ToVo())
	}

	return topicVos

}

var topicRepository repository.TopicRepository

func NewTopicService() TopicService {
	topicRepository = repository.NewTopicRepository()
	return TopicServiceImpl{}
}
