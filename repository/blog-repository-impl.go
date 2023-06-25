package repository

import (
	"gin-demo/common"
	"gin-demo/models"
	"gin-demo/request"
	"gin-demo/vo"
	"gorm.io/gorm"
	"time"
)

type BlogRepositoryImpl struct {
	db *gorm.DB
}

func (b BlogRepositoryImpl) UpdateBlogEyeCount(bid int, count int64) error {
	var sql = "update blogs set eye_count = ? where id = ?"
	return b.db.Raw(sql, count, bid).Error
}

func (b BlogRepositoryImpl) UpdateBlogLikeCount(bid int, count int64) error {
	var sql = "update blogs set like_count = ? where id = ?"
	return b.db.Raw(sql, count, bid).Error
}

func (b BlogRepositoryImpl) UpdateBlog(blog models.Blog) error {

	var ids []int

	for _, tag := range blog.Tags {
		ids = append(ids, tag.Id)
	}

	b.db.Model(&blog).Association("Tags").Replace(blog.Tags)

	var err = b.db.Model(&models.Blog{}).Where("id = ?", blog.Id).Save(&blog).Error

	if err != nil {
		return err
	}

	return nil
}

func (b BlogRepositoryImpl) AddBlogToDb(blog models.Blog) models.Blog {
	b.db.Model(&models.Blog{}).Create(&blog)
	return blog
}

func (b BlogRepositoryImpl) SaveLikeBlog(like models.BlogLike) error {
	return b.db.Model(&models.BlogLike{}).Save(&like).Error
}

func (b BlogRepositoryImpl) CurrentIpIsLikeBlog(ip string, id int) (count int64) {
	b.db.Model(&models.BlogLike{}).Where("ip = ? and blog_id = ?", ip, id).Count(&count)
	return count
}

func (b BlogRepositoryImpl) FindBlogById(id int) (blog models.Blog) {
	b.db.Model(&models.Blog{}).Preload("User").Preload("Topic").Preload("Category").Preload("Tags").First(&blog, "id = ?", id)
	return blog
}

func (b BlogRepositoryImpl) FindBlogByUserTop(uid int) (blogs []vo.SimpleBlogVo) {
	var query = b.db.Model(&models.Blog{}).Select("id,title").Where("category_id is not null and user_id = ?", uid)
	query.Offset(0).Limit(common.USER_HOT_BLOG_COUNT).Order(request.Sort("EYE").String()).Find(&blogs)
	return blogs
}

func (b BlogRepositoryImpl) FindBlogByUserId(uid int, page int) (blogs []models.Blog, count int64) {
	var query = b.db.Model(&models.Blog{}).Select("id,title,description,cover_image,create_at,user_id,category_id").Where("category_id is not null and user_id = ?", uid)

	if err := query.Count(&count).Error; err != nil || count == 0 {
		return make([]models.Blog, 0), 0
	}

	query.Offset((page - 1) * common.PAGE_SIZE).Limit(common.PAGE_SIZE).Preload("User").Preload("Category").Order(request.Sort("CREATE").String()).Find(&blogs)

	return blogs, count
}

func (b BlogRepositoryImpl) FindRangeBlog(page int, start time.Time, end time.Time) (blogs []vo.ArchiveBlogVo, count int64) {

	var query = b.db.Model(&models.Blog{}).Select("id,title,description,create_at").Where("create_at BETWEEN ? AND ?", start, end)

	if err := query.Count(&count).Error; err != nil || count == 0 {
		return make([]vo.ArchiveBlogVo, 0), 0
	}

	query.Offset((page - 1) * common.PAGE_SIZE).Limit(common.ARCHIVE_PAGE_SIZE).Order(request.Sort(request.CREATE).String()).Find(&blogs)

	return blogs, count

}

func (b BlogRepositoryImpl) FindAllSimpleSearchBlog() (blogs []vo.SimpleBlogVo) {
	b.db.Model(&models.Blog{}).Select("id,title,description").Find(&blogs)
	return blogs
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

	if err := query.Count(&count).Error; err != nil || count == 0 {
		return make([]models.Blog, 0), 0
	}

	query.Offset((pageRequest.Page - 1) * common.PAGE_SIZE).Limit(common.PAGE_SIZE).Preload("User").Preload("Category").Order(pageRequest.Sort.String()).Find(&blogs)

	return blogs, count
}

func NewBlogRepository(db *gorm.DB) BlogRepository {
	return BlogRepositoryImpl{db: db}
}
