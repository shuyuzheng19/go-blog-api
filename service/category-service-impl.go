package service

import (
	"encoding/json"
	"gin-demo/cache"
	"gin-demo/common"
	"gin-demo/config"
	"gin-demo/models"
	"gin-demo/myerr"
	"gin-demo/repository"
	"gin-demo/vo"
	"go.uber.org/zap"
	"time"
)

type CategoryServiceImpl struct {
	repository repository.CategoryRepository
	cache      cache.CategoryCache
}

func (c CategoryServiceImpl) AddCategory(name string) vo.CategoryVo {
	var category = models.Category{
		Name:     name,
		CreateAt: time.Now(),
	}
	var result = c.repository.AddCategory(category)

	if result.Id == 0 {
		myerr.PanicError(common.ADD_CATEGORY_ERROR)
	}

	c.cache.RemoveKey()

	config.LOGGER.Info("添加分类", zap.String("name", name))

	return *result.ToVo()
}

func (c CategoryServiceImpl) GetAllCategoryListForCache() (list []vo.CategoryVo) {
	var categoryStrs = c.cache.GetCategoryList()

	if len(categoryStrs) == 0 {
		list = c.repository.FindAllCategoryList()
		c.cache.SaveCategoryList(list)
	} else {
		for _, str := range categoryStrs {
			var category vo.CategoryVo
			var err = json.Unmarshal([]byte(str), &category)
			if err != nil {
				continue
			}
			list = append(list, category)
		}
	}

	return list
}

func (c CategoryServiceImpl) GetAllCategoryListForDB() []vo.CategoryVo {
	return c.repository.FindAllCategoryList()
}

func NewCategoryService() CategoryService {
	var categoryRepository = repository.NewCategoryRepository(config.DB)
	var categoryCache = cache.NewCategoryCache(config.REDIS)
	return CategoryServiceImpl{repository: categoryRepository, cache: categoryCache}
}
