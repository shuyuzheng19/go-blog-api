package service

import (
	"gin-demo/cache"
	"gin-demo/common"
	"gin-demo/config"
	"gin-demo/models"
	"gin-demo/myerr"
	"gin-demo/repository"
	"gin-demo/response"
	"gin-demo/vo"
)

type CommentServiceImpl struct {
	repository repository.CommentRepository
	cache      cache.CommentCache
}

func (c CommentServiceImpl) GetUserLikeComments(uid int) []string {
	return c.cache.GetUserCommentLikes(uid)
}

func (c CommentServiceImpl) IsLikeComment(uid int, cid int64) bool {
	return c.cache.UserIsLikeComment(uid, cid)
}

func (c CommentServiceImpl) LikeComment(uid int, cid int64) bool {
	var result = c.cache.AddUserLike(uid, cid)

	if result == 0 {
		return false
	}

	var err = c.repository.LikeCountPlusOne(cid)

	if err != nil {
		return false
	}

	return true
}

func (c CommentServiceImpl) UnLikeComment(uid int, cid int64) bool {

	var result = c.cache.CancelUserLike(uid, cid)

	if result == 0 {
		return false
	}

	var err = c.repository.LikeCountSubOne(cid)

	if err != nil {
		return false
	}

	return true
}

func (c CommentServiceImpl) GetBlogComment(blogId int, page int) response.PageInfo {
	var comments, count = c.repository.GetBlogComments(blogId, page)

	var commentVos = make([]vo.CommentVo, 0)

	for _, comment := range comments {
		commentVos = append(commentVos, comment.ToVo())
	}

	var pageInfo = response.PageInfo{
		Page:  page,
		Size:  common.COMMENT_PAGE_COUNT,
		Total: count,
		Data:  commentVos,
	}

	return pageInfo
}

func (c CommentServiceImpl) AddComment(comment models.Comment) vo.CommentVo {
	comment.UserId = 1
	var commentResult = c.repository.SaveComment(comment)
	if commentResult.Id == 0 {
		myerr.PanicError(common.COMMENT_ERROR)
		return vo.CommentVo{}
	}
	return commentResult.ToVo()
}

func NewCommentService() CommentService {
	var repository = repository.NewCommentRepository(config.DB)
	var cache = cache.NewCommentCache(config.REDIS)
	return CommentServiceImpl{repository: repository, cache: cache}
}
