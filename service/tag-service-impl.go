package service

import (
	"encoding/json"
	"gin-demo/cache"
	"gin-demo/common"
	"gin-demo/config"
	"gin-demo/models"
	"gin-demo/myerr"
	"gin-demo/repository"
	"gin-demo/response"
	"gin-demo/vo"
	"go.uber.org/zap"
	"strconv"
	"time"
)

type TagServiceImpl struct {
	repository repository.TagRepository
	cache      cache.TagCache
}

func (t TagServiceImpl) AddTag(name string) vo.TagVo {
	var tag = models.Tag{
		Name:     name,
		CreateAt: time.Now(),
	}
	var result = t.repository.AddTag(tag)

	if result.Id == 0 {
		myerr.PanicError(common.ADD_TAG_ERROR)
	}

	t.cache.RemoveKey()

	config.LOGGER.Info("添加标签", zap.String("name", name))

	return result.ToVo()
}

func (t TagServiceImpl) GetAllTag() (tags []vo.TagVo) {
	return t.repository.FindAllSimpleTag()
}

func (t TagServiceImpl) RandomTag() (tags []vo.TagVo) {
	var result = t.cache.GetRandomTags()

	if len(result) == 0 {
		t.cache.SaveAllTags(t.GetAllTag())
	}

	var tagsStrs = t.cache.GetRandomTags()

	for _, tagStr := range tagsStrs {
		var tag vo.TagVo
		json.Unmarshal([]byte(tagStr), &tag)
		tags = append(tags, tag)
	}

	return tags
}

func (t TagServiceImpl) GetTagBlogByPage(page int, tid int) response.PageInfo {

	if page == 1 {
		var result = t.cache.GetFirstPageBlog(strconv.Itoa(tid))

		if result != "" {
			var pageInfo response.PageInfo

			json.Unmarshal([]byte(result), &pageInfo)

			return pageInfo
		}

	}

	var blogs, count = t.repository.FindBlogByTagId(tid, page)

	var blogVos []vo.BlogVo

	for _, blog := range blogs {
		blogVos = append(blogVos, blog.ToVo())
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

func (t TagServiceImpl) GetTag(tid int) (tag vo.TagVo) {

	var idStr = strconv.Itoa(tid)

	var result = t.cache.GetTagInfo(idStr)

	if result == "" {
		tag = t.repository.FindTagById(tid)
		if tag.Id > 0 {
			t.cache.SaveTagInfo(tag)
		} else {
			myerr.PanicError(common.TAG_NOT_FOUND)
		}
	} else {
		var cacheTagStr = t.cache.GetTagInfo(idStr)
		json.Unmarshal([]byte(cacheTagStr), &tag)
	}

	return tag
}

func NewTagService() TagService {
	var tagRepository = repository.NewTagRepository(config.DB)
	var tagCache = cache.NewTagCache(config.REDIS)
	return TagServiceImpl{repository: tagRepository, cache: tagCache}
}
