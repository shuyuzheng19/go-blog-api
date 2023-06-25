package service

import (
	"encoding/json"
	"gin-demo/cache"
	"gin-demo/common"
	"gin-demo/config"
	"gin-demo/myerr"
	"gin-demo/repository"
	"gin-demo/request"
	"gin-demo/response"
	"gin-demo/vo"
	"go.uber.org/zap"
	"strconv"
)

type TopicServiceImpl struct {
	repository repository.TopicRepository
	cache      cache.TopicCache
}

func (t TopicServiceImpl) AddTopic(uid int, request request.TopicRequest) vo.SimpleTopicVo {

	var err = request.Check()

	myerr.IError(err, common.FAIL_CODE)

	var topic = request.ToDo(uid)

	var result = t.repository.AddTopic(topic)

	if result.Id == 0 {
		myerr.PanicError(common.ADD_TOPIC_ERROR)
	}

	t.cache.RemoveTopicPageKey()

	config.LOGGER.Info("添加专题", zap.Int("user_id", result.UserId), zap.String("name", result.Name))

	return result.ToSimpleVo()
}

func (t TopicServiceImpl) GetUserTopic(uid int) []vo.SimpleTopicVo {
	return t.repository.GetUserTopic(uid)
}

func (t TopicServiceImpl) TopicBlogList(tid int) []vo.SimpleBlogVo {
	var blogs = t.repository.FindAllTopicList(tid)
	return blogs
}

func (t TopicServiceImpl) GetUserTopics(uid int) []vo.SimpleTopicVo {
	return t.repository.FindTopicByUser(uid)
}

func (t TopicServiceImpl) GetTopicById(tid int) vo.SimpleTopicVo {
	var result = t.cache.GetTopicFromMap(strconv.Itoa(tid))

	if result == "" {
		var topic = t.repository.FindTopicById(tid)

		if topic.Id > 0 {
			t.cache.SetTopicToMap(topic)
		} else {
			myerr.PanicError(common.TOPIC_NOT_FOUND)
		}

		return topic

	} else {
		var topic vo.SimpleTopicVo

		json.Unmarshal([]byte(result), &topic)

		return topic
	}

}

func (t TopicServiceImpl) GetTopicBlogByPage(tid int, page int) response.PageInfo {

	if page == 1 {
		var result = t.cache.GetFirstPageBlog(strconv.Itoa(tid))

		if result != "" {
			var pageInfo response.PageInfo

			json.Unmarshal([]byte(result), &pageInfo)

			return pageInfo
		}
	}

	var blogs, count = t.repository.FindTopicBlogByPage(tid, page)

	var blogVos []vo.BlogVo

	for _, blog := range blogs {
		blogVos = append(blogVos, blog.ToTopicVo())
	}

	var pageInfo = response.PageInfo{
		Page:  page,
		Size:  common.PAGE_SIZE,
		Total: count,
		Data:  blogVos,
	}

	if pageInfo.Total > 0 {
		t.cache.SaveFirstPageBlog(strconv.Itoa(tid), pageInfo)
	}

	return pageInfo

}

func (t TopicServiceImpl) GetTopicByPage(page int) response.PageInfo {

	if page == 1 {
		var topicStr = t.cache.GetFirstPageTopic()

		if topicStr != "" {
			var pageInfo response.PageInfo

			json.Unmarshal([]byte(topicStr), &pageInfo)

			return pageInfo
		}

	}

	var topics, count = t.repository.FindTopicByPage(page)

	var topicVos []vo.TopicVo

	for _, topic := range topics {
		topicVos = append(topicVos, topic.ToVo())
	}

	var pageInfo = response.PageInfo{
		Page:  page,
		Size:  common.TOPIC_PAGE_SIZE,
		Total: count,
		Data:  topicVos,
	}

	if page == 1 && count > 0 {
		t.cache.SaveFirstPageTopic(pageInfo)
	}

	return pageInfo
}

func NewTopicService() TopicService {
	var topicRepository = repository.NewTopicRepository(config.DB)
	var topicCache = cache.NewTopicCache(config.REDIS)
	return TopicServiceImpl{repository: topicRepository, cache: topicCache}
}
