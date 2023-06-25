package controllers

import (
	"bufio"
	"gin-demo/common"
	"gin-demo/config"
	"gin-demo/myerr"
	"gin-demo/request"
	"gin-demo/service"
	"gin-demo/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type BlogController struct {
	service service.BlogService
}

func (b BlogController) GetBlogByUser(ctx *gin.Context) {

	var uid, err = strconv.Atoi(ctx.Param("uid"))

	if err != nil {
		ctx.JSON(http.StatusOK, common.USER_ID_ERROR)
		return
	}

	var page, _ = strconv.Atoi(ctx.DefaultQuery("page", "1"))

	var pageInfo = b.service.GetBlogByUser(uid, page)

	ctx.JSON(http.StatusOK, common.Success(pageInfo))

}

func (b BlogController) IsLikeBlog(ctx *gin.Context) {
	var id, err = strconv.Atoi(ctx.Param("id"))

	if err != nil || id <= 0 {
		ctx.JSON(http.StatusOK, common.BLOG_ID_ERROR)
		return
	}

	var ip = utils.GetIPAddress(ctx.Request)

	var flag = b.service.CurrentIpIsLikeBlog(ip, id)

	ctx.JSON(http.StatusOK, common.Success(flag))
}

func (b BlogController) Chat(ctx *gin.Context) {

	var token = ctx.GetHeader(common.TOKEN_HEADER)

	if token == "" || !strings.HasPrefix(token, common.TOKEN_TYPE) {
		ctx.JSON(403, common.CHAT_TOKEN_ERROR)
		return
	}

	var message = ctx.Query("message")

	if message == "" {
		ctx.JSON(http.StatusOK, common.Fail("请输入要提问的内容"))
		return
	}

	flusher, ok := ctx.Writer.(http.Flusher)

	if !ok {
		ctx.String(http.StatusInternalServerError, "Streaming unsupported")
		return
	}

	var reader = b.service.Chat(token, message)

	if reader == nil {
		ctx.JSON(403, common.CHAT_AUTHENTICATE_ERROR)
		return
	}

	defer reader.Close()

	ctx.Stream(func(w io.Writer) bool {
		var scanner = bufio.NewScanner(reader)

		for scanner.Scan() {

			var text = scanner.Text()

			if strings.HasPrefix(text, "data: ") {
				ctx.SSEvent("", strings.Replace(text, "data: ", "", 1))
			}

			flusher.Flush()
		}
		return false
	})
}

func (b BlogController) LikeBlog(ctx *gin.Context) {
	var id, err = strconv.Atoi(ctx.Param("id"))

	if err != nil || id <= 0 {
		ctx.JSON(http.StatusOK, common.BLOG_ID_ERROR)
		return
	}

	var ip = utils.GetIPAddress(ctx.Request)

	var count = b.service.AddLikeBlog(ip, id)

	ctx.JSON(http.StatusOK, common.Success(count))

	config.LOGGER.Info("点赞博客", zap.String("ip", ip), zap.Int("blog_id", id))

}

func (b BlogController) GetBlogById(ctx *gin.Context) {

	var id, err = strconv.Atoi(ctx.Param("id"))

	if err != nil || id <= 0 {
		ctx.JSON(http.StatusOK, common.BLOG_ID_ERROR)
		return
	}

	var blog = b.service.GetBlogById(id)

	ctx.JSON(http.StatusOK, common.Success(blog))

}

func (b BlogController) GetUserBlogTop(ctx *gin.Context) {

	var uid, err = strconv.Atoi(ctx.Param("uid"))

	if err != nil || uid <= 0 {
		ctx.JSON(http.StatusOK, common.USER_ID_ERROR)
		return
	}

	var blogs = b.service.GetUserBlogTop(uid)

	ctx.JSON(http.StatusOK, common.Success(blogs))

}

func (b BlogController) SaveUserEditor(ctx *gin.Context) {
	var maps map[string]string

	ctx.ShouldBindJSON(&maps)

	var content = maps["content"]

	if content == "" {
		ctx.JSON(http.StatusOK, common.Fail("保存的内容不能为空"))
		return
	}

	b.service.SaveUserEditorContent(GetUser(ctx).Id, content)

	ctx.JSON(http.StatusOK, common.OK())

}

func (b BlogController) GetUserEditor(ctx *gin.Context) {

	var content = b.service.GetUserEditorContent(GetUser(ctx).Id)

	if content == "" {
		ctx.JSON(http.StatusOK, common.Fail("没有要获取的内容"))
		return
	}

	ctx.JSON(http.StatusOK, common.Success(content))
}

func (b BlogController) SaveBlog(ctx *gin.Context) {

	var requestBlog request.BlogRequest

	var err = ctx.ShouldBindJSON(&requestBlog)

	if err != nil {
		myerr.PanicError(common.BAD_REQUEST_ERROR)
	}

	b.service.SaveBlog(GetUser(ctx).Id, requestBlog)

	ctx.JSON(http.StatusOK, common.OK())

}

func (b BlogController) GetEditBlog(ctx *gin.Context) {

	var id, ierr = strconv.Atoi(ctx.Param("bid"))

	if ierr != nil {
		myerr.PanicError(common.BLOG_ID_ERROR)
	}

	var blog = b.service.GetBlogById(id)

	var user = GetUser(ctx)

	if user.Role.Name != common.SUPER_ROLE && user.Id != blog.User.Id {
		myerr.PanicError(common.UPDATE_BLOG_AUTH_ERROR)
	}

	var tags = make([]int, 0)

	if len(blog.Tags) > 0 {
		for _, tag := range blog.Tags {
			tags = append(tags, tag.Id)
		}
	}

	var tid = 0

	var cid = 0

	if blog.Topic != nil {
		tid = blog.Topic.Id
	}

	if blog.Category != nil {
		cid = blog.Category.Id
	}

	var requestBlog = request.BlogRequest{
		Title:       blog.Title,
		Description: blog.Description,
		Content:     blog.Content,
		Markdown:    blog.Markdown,
		SourceUrl:   blog.SourceUrl,
		CoverImage:  blog.CoverImage,
		Tags:        tags,
		Topic:       tid,
		Category:    cid,
	}

	ctx.JSON(http.StatusOK, common.Success(requestBlog))

}

func (b BlogController) UpdateBlog(ctx *gin.Context) {

	var id, ierr = strconv.Atoi(ctx.Param("bid"))

	if ierr != nil {
		myerr.PanicError(common.BLOG_ID_ERROR)
	}

	var requestBlog request.BlogRequest

	var err = ctx.ShouldBindJSON(&requestBlog)

	if err != nil {
		myerr.PanicError(common.BAD_REQUEST_ERROR)
	}

	b.service.UpdateBlog(id, GetUser(ctx), requestBlog)

	ctx.JSON(http.StatusOK, common.OK())

}

func (b BlogController) FindBlogPage(ctx *gin.Context) {

	var page, _ = strconv.Atoi(ctx.DefaultQuery("page", "1"))

	var cid, _ = strconv.Atoi(ctx.DefaultQuery("cid", "-1"))

	var pageRequest = request.BlogPageRequest{
		Page: page,
		Cid:  cid,
		Sort: request.Sort(ctx.Query("sort")),
	}

	var pageInfo = b.service.FindBlogByCidPage(pageRequest)

	ctx.JSON(http.StatusOK, common.Success(pageInfo))

}

func (b BlogController) GetHotBlog(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, common.Success(b.service.GetHotBlog()))
}

func (b BlogController) SetRecommendBlog(ctx *gin.Context) {
	var ids []int

	ctx.ShouldBindJSON(&ids)

	b.service.SetRecommend(ids)

	ctx.JSON(http.StatusOK, common.OK())
}

func (b BlogController) GetRecommendBlog(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, common.Success(b.service.GetRecommend()))
}

func (b BlogController) SearchBlog(ctx *gin.Context) {
	var keyword = ctx.Query("keyword")

	if keyword == "" {
		ctx.JSON(http.StatusOK, common.NO_KEYWORD)
		return
	}

	var blogs = b.service.SearchBlog(keyword)

	ctx.JSON(http.StatusOK, common.Success(blogs))
}

func (b BlogController) CountSearchBlog(ctx *gin.Context) {
	var keyword = ctx.Query("keyword")
	if keyword == "" {
		ctx.JSON(http.StatusOK, common.NO_KEYWORD)
		return
	}
	var pageStr = ctx.DefaultQuery("page", "1")

	var page, _ = strconv.Atoi(pageStr)

	var pageInfo = b.service.CountSearch(keyword, page)

	ctx.JSON(http.StatusOK, common.Success(pageInfo))
}

func (b BlogController) SimilarBlog(ctx *gin.Context) {
	var keyword = ctx.Query("keyword")
	if keyword == "" {
		ctx.JSON(http.StatusOK, common.NO_KEYWORD)
		return
	}

	var blogs = b.service.SearchBlog(keyword)

	ctx.JSON(http.StatusOK, common.Success(blogs))
}

func (b BlogController) RangeBlog(ctx *gin.Context) {

	var startStr = ctx.DefaultQuery("start", "0")

	var endStr = ctx.DefaultQuery("end", "0")

	var start, _ = strconv.Atoi(startStr)

	var end, _ = strconv.Atoi(endStr)

	if start == 0 || end == 0 {
		ctx.JSON(http.StatusOK, common.RANGE_TIME_EMPTY)
		return
	}

	var pageStr = ctx.DefaultQuery("page", "1")

	var page, _ = strconv.Atoi(pageStr)

	var pageInfo = b.service.GetRangeBlog(page, int64(start), int64(end))

	ctx.JSON(http.StatusOK, common.Success(pageInfo))
}

func (b BlogController) GetHotKeyword(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, common.Success(b.service.GetHotKeywords()))
}

func (b BlogController) GetRandomBlog(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, common.Success(b.service.RandomBlogs()))
}

func NewBlogController(service service.BlogService) BlogController {
	return BlogController{service: service}
}
