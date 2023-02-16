package service

import (
	"encoding/json"
	"strings"
	"time"
	"vs-blog-api/common"
	"vs-blog-api/config"
	"vs-blog-api/modal"
	"vs-blog-api/repository"
	"vs-blog-api/response"
	"vs-blog-api/utils"
)

var categoryRepository repository.CategoryRepository

type CategoryServiceImpl struct {
}

func NewCategoryService() CategoryService {
	categoryRepository = repository.NewCategoryRepository()
	return CategoryServiceImpl{}
}

func (c CategoryServiceImpl) FindAllCategory() []modal.CategoryVo {

	const KEY = common.CategoryList

	list := config.Redis.LRange(KEY, 0, -1).Val()

	var result []modal.CategoryVo

	if len(list) == 0 {
		err, categories := categoryRepository.FindAllCategory()

		if err != nil {
			panic(response.NewGlobalException(response.ERROR, "后台接口出错,查询分类失败!"))
		}
		var categoryStrs []string

		for _, category := range categories {
			categoryVo := category.ToVo()

			result = append(result, *categoryVo)

			categoryStrs = append(categoryStrs, utils.ObjectToJson(categoryVo))
		}

		config.Redis.RPush(KEY, categoryStrs)

		config.Redis.Expire(KEY, common.CategoryListExpire)

	} else {

		for _, str := range list {
			var category modal.CategoryVo
			if err := json.Unmarshal([]byte(str), &category); err == nil {
				result = append(result, category)
			}
		}
	}

	return result

}

func (c CategoryServiceImpl) SaveCategory(name string) *modal.CategoryVo {

	if name == "" || strings.TrimSpace(name) == "" {
		panic(response.NewGlobalException(response.ERROR, "分类名字不能为空!"))
	}

	var category = modal.Category{
		Id:         0,
		Name:       name,
		CreateTime: time.Now(),
		DeletedAt:  nil,
	}

	err, m := categoryRepository.AddCategory(category)

	if err != nil {
		panic(response.NewGlobalException(response.ERROR, "添加失败"))
	}

	config.Redis.Del(common.CategoryList)

	return m.ToVo()

}
