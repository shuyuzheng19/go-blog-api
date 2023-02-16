package service

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
	"vs-blog-api/common"
	"vs-blog-api/config"
	"vs-blog-api/modal"
	"vs-blog-api/repository"
	"vs-blog-api/response"
	"vs-blog-api/utils"
)

type TagServiceImpl struct {
}

func (t TagServiceImpl) SaveTag(name string) modal.TagVo {

	if name == "" || strings.TrimSpace(name) == "" {
		panic(response.NewGlobalException(response.ERROR, "标签名字不能为空!"))
	}

	var tag = modal.Tag{
		Id:         0,
		Name:       name,
		CreateTime: time.Now(),
		DeletedAt:  nil,
	}

	err, m := tagRepository.AddTag(tag)

	if err != nil {
		panic(response.NewGlobalException(response.ERROR, "添加失败"))
	}

	config.Redis.Del(common.RandomTag)

	return m.ToVo()
}

func (t TagServiceImpl) GetAllTag() []modal.TagVo {

	exists := config.Redis.Exists(common.RandomTag).Val()

	if exists <= 0 {
		t.GetRandomTag()
	}

	tagsStr := config.Redis.SMembers(common.RandomTag).Val()

	var tags = make([]modal.TagVo, 0)

	for _, tagStr := range tagsStr {

		var tagVo modal.TagVo

		if err := json.Unmarshal([]byte(tagStr), &tagVo); err == nil {
			tags = append(tags, tagVo)
		}

	}

	return tags
}

func (t TagServiceImpl) FindByIdTag(id int) modal.TagVo {
	if id <= 0 {
		panic(response.NewGlobalException(response.ParamsError, "非法ID"))
	}

	err, tag := tagRepository.FindById(id)

	if err != nil {
		panic(response.NewGlobalException(response.NOTFOUND, "找不到该标签"))
	}

	return tag.ToVo()
}

func (t TagServiceImpl) GetTagBlog(id int, page int) response.PageInfoResponse {

	if id <= 0 || page <= 0 {
		panic(response.NewGlobalException(response.ParamsError, "含有非法参数"))
	}

	err, blogs, count := tagRepository.FindTagIdBlog(id, page)

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

var tagRepository repository.TagRepository

func NewTagService() TagService {
	tagRepository = repository.NewTagRepository()
	return TagServiceImpl{}
}

func (t TagServiceImpl) GetRandomTag() []modal.TagVo {
	var KEY = common.RandomTag

	var tags []modal.TagVo

	var exists = config.Redis.Exists(KEY).Val()

	if exists == 0 {
		tagResult := tagRepository.FindAllTag()
		var tagStr []string
		for _, tag := range tagResult {
			var tagVo = tag.ToVo()
			fmt.Println(tagVo)
			tags = append(tags, tagVo)
			tagStr = append(tagStr, utils.ObjectToJson(tagVo))
		}
		config.Redis.SAdd(KEY, tagStr)
	} else {
		tagsStr := config.Redis.SRandMemberN(KEY, common.RandomTagCount).Val()

		for _, tagStr := range tagsStr {

			var tagVo modal.TagVo

			if err := json.Unmarshal([]byte(tagStr), &tagVo); err == nil {
				tags = append(tags, tagVo)
			}
		}
	}

	return tags
}
