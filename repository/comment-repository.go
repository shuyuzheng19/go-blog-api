package repository

import "gin-demo/models"

type CommentRepository interface {
	SaveComment(comment models.Comment) models.Comment
	GetBlogComments(blogId int, page int) (comments []models.Comment, count int64)
	LikeCountPlusOne(id int64) error
	LikeCountSubOne(id int64) error
}
