package repository

import (
	"gin-demo/common"
	"gin-demo/models"
	"gin-demo/request"
	"gin-demo/vo"
	"gorm.io/gorm"
)

type BlogRepositoryImpl struct {
	db *gorm.DB
}

func (b BlogRepositoryImpl) FindAllSimpleBlog() (blogs []vo.SimpleBlogVo) {
	b.db.Model(&models.Blog{}).Select("id,title").Find(&blogs)
	return blogs
}

func (b BlogRepositoryImpl) FindByIdIn(ids []int) (blogs []vo.SimpleBlogVo) {

	b.db.Model(&models.Blog{}).Select("id,title,cover_image").Where("id in ?", ids).Find(&blogs)

	return blogs
}

func (b BlogRepositoryImpl) GetHotBlog() (blogs []vo.SimpleBlogVo) {
	b.db.Model(&models.Blog{}).Select("id,title").Order(request.Sort(request.EYE).String()).Offset(0).Limit(common.HOT_BLOG_SIZE).Find(&blogs)
	return blogs
}

func (b BlogRepositoryImpl) PaginatedBlogQueries(pageRequest request.BlogPageRequest) (blogs []models.Blog, count int64) {
	var query = b.db.Model(&models.Blog{}).Select("id,title,description,cover_image,create_at,user_id,category_id")

	if pageRequest.Cid > 0 {
		query.Where("category_id = ?", pageRequest.Cid)
	} else {
		query.Where("category_id is not null")
	}

	if err := query.Count(&count).Error; err != nil {
		return nil, 0
	}

	query.Offset((pageRequest.Page - 1) * common.PAGE_SIZE).Limit(common.PAGE_SIZE).Preload("User").Preload("Category").Order(pageRequest.Sort.String()).Find(&blogs)

	return blogs, count
}

func NewBlogRepository(db *gorm.DB) BlogRepository {
	return BlogRepositoryImpl{db: db}
}
