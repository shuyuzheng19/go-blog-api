package controllers

import (
	"gin-demo/common"
	"gin-demo/myerr"
	"gin-demo/request"
	"gin-demo/service"
	"gin-demo/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CommentController struct {
	service service.CommentService
}

func (c CommentController) GetBlogComment(ctx *gin.Context) {

	var blogId, err = strconv.Atoi(ctx.Param("bid"))

	if err != nil || blogId <= 0 {
		ctx.JSON(http.StatusOK, common.BLOG_ID_ERROR)
		return
	}

	var page, _ = strconv.Atoi(ctx.DefaultQuery("page", "1"))

	var pageInfo = c.service.GetBlogComment(blogId, page)

	ctx.JSON(http.StatusOK, common.Success(pageInfo))

}

func (c CommentController) LikeComment(ctx *gin.Context) {

	var cid, err = strconv.Atoi(ctx.Param("cid"))

	if err != nil {
		ctx.JSON(http.StatusOK, common.COMMENT_ID_ERROR)
		return
	}

	var user = GetUser(ctx)

	var uid = user.Id

	var id = int64(cid)

	var isLike = c.service.IsLikeComment(uid, id)

	if isLike {
		var flag = c.service.UnLikeComment(uid, id)
		ctx.JSON(http.StatusOK, common.Success(flag))
	} else {
		var flag = c.service.LikeComment(uid, id)
		ctx.JSON(http.StatusOK, common.Success(flag))
	}
}

func (c CommentController) GetCommentUser(ctx *gin.Context) {

	var user = GetUser(ctx)

	var vo = user.ToCommentUserVo()

	vo.LikeIds = c.service.GetUserLikeComments(user.Id)

	ctx.JSON(http.StatusOK, common.Success(vo))

}

func (c CommentController) AddComment(ctx *gin.Context) {

	var user = GetUser(ctx)

	var commentRequest request.CommentRequest

	ctx.ShouldBindJSON(&commentRequest)

	var err = commentRequest.Check()

	myerr.IError(err, common.FAIL_CODE)

	var commentDo = commentRequest.ToDo()

	var ip = "42.228.112.250"

	commentDo.Ip = ip

	commentDo.UserAgent = ctx.Request.UserAgent()

	commentDo.UserId = user.Id

	commentDo.Address = utils.GetIpCity(ip)

	var comment = c.service.AddComment(commentDo)

	comment.User = user.ToCommentUserVo()

	comment.UID = user.Id

	ctx.JSON(http.StatusOK, common.Success(comment))

}

func NewCommentController(service service.CommentService) CommentController {
	return CommentController{service: service}
}
