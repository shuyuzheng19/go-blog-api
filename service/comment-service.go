package service

import (
	"gin-demo/models"
	"gin-demo/response"
	"gin-demo/vo"
)

type CommentService interface {
	AddComment(comment models.Comment) vo.CommentVo
	GetBlogComment(blogId int, page int) response.PageInfo
	LikeComment(uid int, cid int64) bool
	UnLikeComment(uid int, cid int64) bool
	IsLikeComment(uid int, cid int64) bool
	GetUserLikeComments(uid int) []string
}
