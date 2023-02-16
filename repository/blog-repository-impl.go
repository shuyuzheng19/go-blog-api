package repository

import (
	"gorm.io/gorm"
	"strings"
	"vs-blog-api/common"
	"vs-blog-api/config"
	"vs-blog-api/dto"
	"vs-blog-api/modal"
	"vs-blog-api/utils"
)

type BlogRepositoryImpl struct {
}

var SELECT = []string{"id", "title", "markdown", "description", "cover_image", "create_time", "category_id", "user_id"}

var SELECT_JOIN = strings.Join(SELECT, ",")

func (b BlogRepositoryImpl) SaveBlog(blog modal.Blog) (error, modal.Blog) {
	err := config.DB.Model(&modal.Blog{}).Create(&blog).Error
	return err, blog
}

func (b BlogRepositoryImpl) FindBlogIdIn(ids []string) (err error, blogs []modal.Blog) {

	var idsStr = strings.Join(ids, ",")

	var sql = "select " + SELECT_JOIN + " from blogs where id in ? order by position(id::text in ?)"

	err = getPreload().Raw(sql, ids, idsStr).Find(&blogs).Error

	return err, blogs

}

func (b BlogRepositoryImpl) FindBlogByUserId(id int, page int, sortField string) (err error, blogs []modal.Blog, count int64) {

	err = getPreload().Offset((page-1)*common.PageSize).Limit(common.PageSize).Order(sortField).Find(&blogs, "user_id = ?", id).Offset(-1).Limit(-1).Count(&count).Error

	return err, blogs, count
}

func (b BlogRepositoryImpl) FindRangeDate(startTime int64, endTime int64, page int, sortField string) (err error, blogs []modal.RangeBlog) {

	build := config.DB.Model(&modal.Blog{}).Offset((page - 1) * common.PageSize).Limit(common.ArchivePageSize).Order(sortField)

	err = build.Where("create_time between ? and ?", utils.FormatDate2(startTime), utils.FormatDate2(endTime)).Find(&blogs).Error

	return err, blogs
}

func (b BlogRepositoryImpl) FindAllIdAndTitle() (err error, blogs []modal.SimpleBlog) {

	config.DB.Model(&modal.Blog{}).Select("title", "id").Find(&blogs)

	return err, blogs
}

func (b BlogRepositoryImpl) FindById(id int) (err error, blog modal.Blog) {
	err = config.DB.Model(&modal.Blog{}).Preload("Tags").Preload("Category").Preload("User").First(&blog, "id = ?", id).Error
	return err, blog
}

func getPreload() *gorm.DB {
	return config.DB.Model(&modal.Blog{}).Select(SELECT).Preload("Category").Preload("User")
}

func NewBlogRepositoryImpl() BlogRepository {
	return BlogRepositoryImpl{}
}

func (b BlogRepositoryImpl) FindAll(sortRequest dto.BlogPageSortDto) (blogs []modal.Blog, err error, count int64) {

	var page = sortRequest.Page

	var size = sortRequest.Size

	var sort = dto.SortMap[sortRequest.Sort]

	build := getPreload().Offset((page - 1) * size).Limit(size).Order(sort).Where("category_id is not null")

	categoryId := sortRequest.SortId

	if categoryId > 0 {
		build.Where("category_id = ?", categoryId)
	}

	err = build.Find(&blogs).Offset(-1).Limit(-1).Count(&count).Error

	return blogs, err, count
}
