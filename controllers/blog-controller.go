package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"vs-blog-api/common"
	"vs-blog-api/dto"
	"vs-blog-api/manager"
	"vs-blog-api/modal"
	"vs-blog-api/response"
	"vs-blog-api/service"
	"vs-blog-api/utils"
)

type blogController struct {
}

var blogService service.BlogService

var blogCountManager *manager.BlogCountManager

func NewBlogController() blogController {
	blogService = service.NewBlogService()
	blogCountManager = manager.NewBlogCountManager()
	return blogController{}
}

func (*blogController) RangeBlog(ctx *gin.Context) {
	var startTime = ctx.Query("start")

	start, err := strconv.ParseInt(startTime, 10, 64)

	if err != nil {
		ctx.JSON(200, response.FAILURE(response.ParamsError, "开始日期格式错误"))
		return
	}

	var endTime = ctx.Query("end")

	end, err := strconv.ParseInt(endTime, 10, 64)

	if err != nil {
		ctx.JSON(200, response.FAILURE(response.ParamsError, "结束日期格式错误"))
		return
	}

	var page = ctx.DefaultQuery("page", "1")

	pageInfo := blogService.RangeDateBlog(start, end, utils.ToInt(page))

	ctx.JSON(200, response.SUCCESS(pageInfo))
}

func (*blogController) FindBlogs(ctx *gin.Context) {

	page := ctx.DefaultQuery("page", "1")

	sort := ctx.DefaultQuery("sort", "CREATE")

	cid := ctx.DefaultQuery("cid", "-1")

	sortRequest := dto.NewBlogPageSortDto(utils.ToInt(page), sort, utils.ToInt(cid))

	pageResult := blogService.FindBlogs(sortRequest)

	ctx.JSON(200, response.SUCCESS(pageResult))
}

func (*blogController) GetHotBlogs(ctx *gin.Context) {

	blogs := blogService.GetHotBlogs()

	ctx.JSON(200, response.SUCCESS(blogs))
}

func (*blogController) SearchBlog(ctx *gin.Context) {
	var keyword = ctx.Query("keyword")

	if keyword == "" {
		ctx.JSON(200, response.FAILURE(response.ParamsError, "缺少关键字参数"))
	}

	var page = ctx.DefaultQuery("page", "1")

	pageInfo := blogService.SearchBlog(keyword, utils.ToInt(page))

	ctx.JSON(200, response.SUCCESS(pageInfo))
}

func (*blogController) FindUserBlog(ctx *gin.Context) {
	var id = ctx.Param("id")

	var page = ctx.DefaultQuery("page", "1")

	pageInfo := blogService.FindUserBlog(utils.ToInt(id), utils.ToInt(page))

	ctx.JSON(200, response.SUCCESS(pageInfo))
}

func (*blogController) FindUserTop10(ctx *gin.Context) {

	var id = ctx.Param("id")

	blogs := blogService.GetUserTopBlog(utils.ToInt(id))

	ctx.JSON(200, response.SUCCESS(blogs))
}

func (*blogController) GetSimilarBlog(ctx *gin.Context) {

	var id = ctx.Query("blogId")

	var keyword = ctx.Query("keyword")

	blogs := blogService.GetSimilarBlog(keyword, utils.ToInt(id))

	ctx.JSON(200, response.SUCCESS(blogs))
}

func (*blogController) SaveUserBlog(ctx *gin.Context) {

	var content string

	ctx.ShouldBindJSON(&content)

	userId := modal.GetUser(ctx).Id

	blogService.SaveBlog(userId, content)

	ctx.JSON(200, response.OK_RESULT)
}

func (*blogController) GetUserSaveBlog(ctx *gin.Context) {

	userId := modal.GetUser(ctx).Id

	content := blogService.GetSaveBlog(userId)

	ctx.JSON(200, response.SUCCESS(content))
}

func (*blogController) GetBlogById(ctx *gin.Context) {

	var id = ctx.Param("id")

	blog := blogService.FindByIdBlog(utils.ToInt(id))

	var userId = ctx.GetHeader(common.TokenHeader)

	if userId != "" {
		isEye := blogCountManager.IsUserTodayEye(userId, blog.Id)
		if !isEye {
			blog.EyeCount = blogCountManager.BlogEyeCountAdd(blog.Id, blog.EyeCount)
		} else {
			blog.EyeCount = blogCountManager.GetBlogEyeCount(blog.Id, blog.EyeCount)
		}
	}

	blog.LikeCount = blogCountManager.GetLikeCount(blog.Id, blog.LikeCount)

	var blogVo = blog.ToContentBlogVo()

	ctx.JSON(200, response.SUCCESS(blogVo))
}

func (*blogController) RandomBlogs(ctx *gin.Context) {

	blogs := blogService.GetRandomBlogs()

	ctx.JSON(200, response.SUCCESS(blogs))
}

func (*blogController) GetUserLikeBlog(ctx *gin.Context) {
	var userId = ctx.Param("id")

	var page = ctx.DefaultQuery("page", "1")

	pageInfo := blogService.GetUserLikeBlog(utils.ToInt(userId), utils.ToInt(page))

	ctx.JSON(200, response.SUCCESS(pageInfo))
}

func (*blogController) AddBlog(ctx *gin.Context) {

	var blogRequest modal.BlogRequest

	ctx.ShouldBindJSON(&blogRequest)

	var userId = modal.GetUser(ctx).Id

	fmt.Println(blogRequest)

	blogService.AddBlog(userId, blogRequest)

	ctx.JSON(200, response.OK_RESULT)
}

func (*blogController) CurrentUserIsLike(ctx *gin.Context) {

	var blogId = ctx.Query("blogId")

	var userId = modal.GetUser(ctx).Id

	ctx.JSON(200, response.SUCCESS(blogCountManager.IsLike(userId, utils.ToInt(blogId))))

}

func (*blogController) CurrentUserLikeOrUnLikeBlog(ctx *gin.Context) {

	var blogId = ctx.Query("blogId")

	var userId = modal.GetUser(ctx).Id

	if !blogCountManager.IsLike(userId, utils.ToInt(blogId)) {
		count := blogCountManager.Like(userId, utils.ToInt(blogId))

		if count == -1 {
			ctx.JSON(200, response.FAILURE(response.FAIL, "点赞失败"))
			return
		}
		ctx.JSON(200, response.SUCCESS(
			map[string]interface{}{
				"like":  true,
				"count": count,
			}))
	} else {
		count := blogCountManager.UnLike(userId, utils.ToInt(blogId))

		if count == -1 {
			ctx.JSON(200, response.FAILURE(response.FAIL, "取消点赞失败"))
			return
		}

		ctx.JSON(200, response.SUCCESS(
			map[string]interface{}{
				"like":  false,
				"count": count,
			}))
	}
}
