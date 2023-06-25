package repository

import (
	"gin-demo/common"
	"gin-demo/models"
	"gorm.io/gorm"
)

type CommentRepositoryImpl struct {
	db *gorm.DB
}

func (c CommentRepositoryImpl) LikeCountPlusOne(id int64) error {
	return c.db.Model(&models.Comment{}).Where("id = ?", id).UpdateColumn("likes", gorm.Expr("likes + ?", 1)).Error
}

func (c CommentRepositoryImpl) LikeCountSubOne(id int64) error {
	return c.db.Model(&models.Comment{}).Where("id = ?", id).UpdateColumn("likes", gorm.Expr("likes - ?", 1)).Error
}

func (c CommentRepositoryImpl) GetBlogComments(blogId int, page int) (comments []models.Comment, count int64) {

	var query = c.db.Model(&models.Comment{}).Where("parent_id is null and blog_id = ?", blogId)

	if err := query.Count(&count).Error; err != nil || count == 0 {
		return make([]models.Comment, 0), 0
	}

	query.Offset((page - 1) * common.COMMENT_PAGE_COUNT).Limit(common.COMMENT_PAGE_COUNT).Preload("User").Preload("Reply").Preload("Reply.User").Order("likes desc, create_time desc").Find(&comments)

	return comments, count
}

func (c CommentRepositoryImpl) SaveComment(comment models.Comment) models.Comment {

	if comment.ParentId != nil {

		var id = *comment.ParentId

		var count int64

		c.db.Model(&models.Comment{}).Where("id = ? ", id).Count(&count)

		if count > 0 {
			c.db.Model(&models.Comment{}).Create(&comment)
			return comment
		}

	} else {
		var err = c.db.Model(&models.Comment{}).Create(&comment).Error
		if err != nil {
			return models.Comment{}
		}
		return comment
	}

	return comment
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return CommentRepositoryImpl{db: db}
}
