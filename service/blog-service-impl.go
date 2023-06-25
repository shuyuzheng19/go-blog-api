package service

import (
	"encoding/json"
	"gin-demo/cache"
	"gin-demo/common"
	"gin-demo/config"
	"gin-demo/models"
	"gin-demo/myerr"
	"gin-demo/repository"
	"gin-demo/request"
	"gin-demo/response"
	"gin-demo/search"
	"gin-demo/utils"
	"gin-demo/vo"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type BlogServiceImpl struct {
	cache      cache.BlogCache
	repository repository.BlogRepository
	search     *search.MeiliSearch
}

func (b BlogServiceImpl) InitSearch() {
	config.LOGGER.Info("开始初始化搜索.....")
	b.search.DeleteAllDocument(common.INDEX)
	var blogs = b.repository.FindAllSimpleSearchBlog()
	b.AddBlogToSearch(blogs)
	config.LOGGER.Info("搜索初始化完毕")
}

func (b BlogServiceImpl) InitBlogEyeCount() {

	config.LOGGER.Info("开始初始化浏览量......")

	var maps = b.cache.GetBlogEyeCount()

	for key, value := range maps {
		var bid, _ = strconv.Atoi(key)
		var count, _ = strconv.Atoi(value)
		b.repository.UpdateBlogEyeCount(bid, int64(count))
	}

	config.LOGGER.Info("浏览量初始化完毕")

	b.cache.RemoveKeys([]string{common.BLOG_EYE_COUNT_MAP_KEY})

}

func (b BlogServiceImpl) InitBlogLikeCount() {

	config.LOGGER.Info("开始初始化点赞量......")

	var maps = b.cache.GetBlogLikeCount()

	for key, value := range maps {
		var bid, _ = strconv.Atoi(key)
		var count, _ = strconv.Atoi(value)
		b.repository.UpdateBlogLikeCount(bid, int64(count))
	}

	config.LOGGER.Info("点赞量初始化完毕")

	b.cache.RemoveKeys([]string{common.BLOG_LIKE_COUNT_MAP_KEY})

}

func (b BlogServiceImpl) UpdateBlog(id int, user models.User, blogRequest request.BlogRequest) {

	var err2 = blogRequest.Check()

	myerr.IError(err2, common.FAIL_CODE)

	var blog = b.repository.FindBlogById(id)

	if blog.Id == 0 {
		myerr.PanicError(common.NOT_FOUNT_ERROR)
	}

	if user.Role.Name != common.SUPER_ROLE && user.Id != blog.User.Id {
		myerr.PanicError(common.UPDATE_BLOG_AUTH_ERROR)
	}

	if blogRequest.Topic == 0 {
		blog.CategoryId = &blogRequest.Category
		var tags []models.Tag
		for _, tag := range blogRequest.Tags {
			tags = append(tags, models.Tag{
				Id: tag,
			})
		}
		blog.Tags = tags
	} else {
		blog.TopicId = &blogRequest.Topic
	}

	blog.UpdateAt = time.Now()

	blog.Title = blogRequest.Title

	blog.Description = blogRequest.Description

	blog.Content = blogRequest.Content

	blog.CoverImage = blogRequest.CoverImage

	blog.Markdown = blogRequest.Markdown

	blog.SourceUrl = blogRequest.SourceUrl

	var err = b.repository.UpdateBlog(blog)

	if err != nil {
		myerr.PanicError(common.UPDATE_BLOG_ERROR)
	}

	var keys []string

	if blogRequest.Topic == 0 {
		for _, tag := range blogRequest.Tags {
			keys = append(keys, common.FIRST_TAG_BLOG_PAGE_KEY+strconv.Itoa(tag))
		}
	} else {
		keys = append(keys, common.FIRST_TOPIC_BLOG_PAGE_KEY+strconv.Itoa(blogRequest.Topic))
	}

	keys = append(keys, common.FIRST_PAGE_BLOG_PAGE_KEY)

	var simpleBlog = vo.SimpleBlogVo{
		Id:    blog.Id,
		Title: blogRequest.Title,
	}

	b.cache.AddRandomBlog(simpleBlog)

	simpleBlog.Description = blogRequest.Description

	b.AddBlogToSearch([]vo.SimpleBlogVo{simpleBlog})

	b.cache.RemoveBlogFromMap(strconv.Itoa(simpleBlog.Id))

	b.cache.RemoveKeys(keys)

	config.LOGGER.Info("用户修改博客成功=>", zap.Int("user_id", blog.UserId), zap.String("title", blog.Title))
}

func (b BlogServiceImpl) AddBlogToSearch(blog []vo.SimpleBlogVo) {

	var buff, err = json.Marshal(&blog)

	if err == nil {
		b.search.SaveDocuments(common.INDEX, string(buff))
	}

}

func (b BlogServiceImpl) SaveBlog(uid int, blogRequest request.BlogRequest) {

	var err2 = blogRequest.Check()

	myerr.IError(err2, common.FAIL_CODE)

	var blog = blogRequest.ToDo(uid)

	var result = b.repository.AddBlogToDb(blog)

	if result.Id == 0 {
		myerr.PanicError(common.RELEASE_BLOG_ERROR)
	}

	var keys []string

	if result.TopicId == nil {
		for _, tag := range result.Tags {
			keys = append(keys, common.FIRST_TAG_BLOG_PAGE_KEY+strconv.Itoa(tag.Id))
		}
	} else {
		keys = append(keys, common.FIRST_TOPIC_BLOG_PAGE_KEY+strconv.Itoa(*blog.TopicId))
	}

	keys = append(keys, common.FIRST_PAGE_BLOG_PAGE_KEY)

	var simpleBlog = vo.SimpleBlogVo{
		Id:    blog.Id,
		Title: blog.Title,
	}

	b.cache.AddRandomBlog(simpleBlog)

	simpleBlog.Description = result.Description

	b.AddBlogToSearch([]vo.SimpleBlogVo{simpleBlog})

	b.cache.RemoveKeys(keys)

	config.LOGGER.Info("用户添加博客成功=>", zap.Int("user_id", blog.UserId), zap.String("title", blog.Title))
}

func (b BlogServiceImpl) SaveUserEditorContent(uid int, content string) {
	var err = b.cache.SaveUserEditorContent(strconv.Itoa(uid), content)

	if err != nil {
		myerr.PanicError(common.Fail("保存失败"))
	}

}

func (b BlogServiceImpl) GetUserEditorContent(uid int) string {
	return b.cache.GetUserEditorContent(strconv.Itoa(uid))
}

func (b BlogServiceImpl) Chat(token string, message string) io.ReadCloser {

	var id = uuid.NewString()

	var request_json = "{\"action\":\"next\",\"messages\":[{\"id\":\"" + id + "\",\"role\":\"user\",\"content\":{\"parts\":[\"" + message + "\"],\"content_type\":\"text\"}}],\"model\":\"" + "text-davinci-002-render-sha-mobile" + "\",\"parent_message_id\":\"" + uuid.NewString() + "\"}"

	var client = http.Client{}

	var request, _ = http.NewRequest("POST", "https://ai.fakeopen.com/api/conversation", strings.NewReader(request_json))

	request.Header.Add("Content-Type", "application/json;charset=utf-8")

	request.Header.Add("host", "ai.fakeopen.com")

	request.Header.Add("Authorization", token)

	var resp, _ = client.Do(request)

	if resp.StatusCode == 401 || resp.StatusCode == 403 {
		return nil
	}

	var inputStream = resp.Body

	config.LOGGER.Info("GPT提问", zap.String("message", message))

	return inputStream

}

func (b BlogServiceImpl) CurrentIpIsLikeBlog(ip string, id int) bool {
	var count = b.repository.CurrentIpIsLikeBlog(ip, id)

	if count > 0 {
		return true
	} else {
		return false
	}
}

func (b BlogServiceImpl) AddLikeBlog(ip string, id int) int64 {

	var count = b.repository.CurrentIpIsLikeBlog(ip, id)

	if count > 0 {
		myerr.PanicError(common.REPEAT_BLOG_LIKE)
		return 0
	}

	var err = b.repository.SaveLikeBlog(models.BlogLike{
		Ip:     ip,
		BlogId: id,
	})

	if err != nil {
		myerr.PanicError(common.LIKE_BLOG_FAIL)
		return 0
	}

	return b.cache.IncreaseInLike(strconv.Itoa(id))
}

func (b BlogServiceImpl) GetBlogById(id int) (blogVo vo.BlogContentVo) {

	var idStr = strconv.Itoa(id)

	var result = b.cache.GetBlogFromMap(idStr)

	if result == "" {
		var blog = b.repository.FindBlogById(id)
		if blog.Id == 0 {
			myerr.PanicError(common.BLOG_NOT_FOUND)
			return vo.BlogContentVo{}
		} else {
			blogVo = blog.ToContentVo()
			b.cache.SaveBlogToMap(blogVo)
		}
	} else {
		json.Unmarshal([]byte(result), &blogVo)
	}

	blogVo.EyeCount = b.cache.IncreaseInView(blogVo.EyeCount, strconv.Itoa(blogVo.Id))

	var likeCount = b.cache.GetLikeCount(idStr)

	if likeCount > 0 {
		blogVo.LikeCount = likeCount
	}

	return blogVo
}

func (b BlogServiceImpl) GetBlogByUser(uid int, page int) response.PageInfo {

	var blogs, count = b.repository.FindBlogByUserId(uid, page)

	var blogVos []vo.BlogVo

	for _, blog := range blogs {
		blogVos = append(blogVos, blog.ToVo())
	}

	return response.PageInfo{
		Page:  page,
		Size:  common.PAGE_SIZE,
		Total: count,
		Data:  blogVos,
	}
}

func (b BlogServiceImpl) GetUserBlogTop(uid int) []vo.SimpleBlogVo {
	return b.repository.FindBlogByUserTop(uid)
}

func (b BlogServiceImpl) GetRangeBlog(page int, startStamp int64, endStamp int64) response.PageInfo {
	if startStamp > endStamp {
		myerr.PanicError(common.RANGE_TIME_ERROR)
	}

	var start = utils.MillTimeStampToTime(startStamp)

	var end = utils.MillTimeStampToTime(endStamp)

	var blogs, count = b.repository.FindRangeBlog(page, start, end)

	var pageInfo = response.PageInfo{
		Page:  page,
		Size:  common.ARCHIVE_PAGE_SIZE,
		Total: count,
		Data:  blogs,
	}

	return pageInfo
}

func (b BlogServiceImpl) GetHotKeywords() []string {
	return b.cache.GetTop10SearchKeyword()
}

func (b BlogServiceImpl) CountSearch(keyword string, page int) response.PageInfo {
	var searchResponse, err = b.search.Search(common.INDEX, keyword, page)

	myerr.IError(err, common.ERROR_CODE)

	var pageInfo = response.PageInfo{
		Page:  page,
		Size:  common.PAGE_SIZE,
		Total: searchResponse.EstimatedTotalHits,
		Data:  searchResponse.Hits,
	}

	b.cache.SearchCountPlusOne(keyword)

	return pageInfo
}

func (b BlogServiceImpl) SearchBlog(keyword string) interface{} {
	var searchResponse, err = b.search.Search(common.INDEX, keyword, 1)

	myerr.IError(err, common.ERROR_CODE)

	return searchResponse.Hits
}

func (b BlogServiceImpl) RandomBlogs() (blogs []vo.SimpleBlogVo) {
	var result = b.cache.GetRandomBlog()
	if len(result) == 0 {
		b.cache.SaveRandomBlog(b.repository.FindAllSimpleBlog())
	}
	return b.cache.GetRandomBlog()
}

func (b BlogServiceImpl) SetRecommend(ids []int) {
	if len(ids) < 4 {
		myerr.PanicError(common.SAVE_RECOMMEND_ERROR)
	}
	var blogs = b.repository.FindByIdIn(ids)

	b.cache.SetRecommendBlog(blogs)
}

func (b BlogServiceImpl) GetRecommend() (blogs []vo.SimpleBlogVo) {
	return b.cache.GetRecommendBlog()
}

func (b BlogServiceImpl) GetHotBlog() []vo.SimpleBlogVo {
	var blogs = b.cache.GetTopBlog()

	if len(blogs) == 0 {
		blogs = b.repository.GetHotBlog()
		b.cache.SaveTopBlog(blogs)
	} else {
		blogs = b.cache.GetTopBlog()
	}

	return blogs
}

func (b BlogServiceImpl) FindBlogByCidPage(pageRequest request.BlogPageRequest) response.PageInfo {

	if pageRequest.Page == 1 && pageRequest.Cid == -1 && pageRequest.Sort == "CREATE" {
		var result = b.cache.GetHomeFirstPageBlog()

		if result != "" {
			var pageInfo response.PageInfo
			json.Unmarshal([]byte(result), &pageInfo)
			return pageInfo
		}

	}

	var blogs, count = b.repository.PaginatedBlogQueries(pageRequest)

	var blogVos []vo.BlogVo

	for _, blog := range blogs {
		blogVos = append(blogVos, blog.ToVo())
	}

	var pageInfo = response.PageInfo{
		Page:  pageRequest.Page,
		Size:  common.PAGE_SIZE,
		Total: count,
		Data:  blogVos,
	}

	if pageInfo.Page == 1 && pageRequest.Cid == -1 && pageRequest.Sort == "CREATE" {
		b.cache.SaveHomeFirstPageBlog(pageInfo)
	}

	return pageInfo
}

func NewBlogService() BlogService {
	var repository = repository.NewBlogRepository(config.DB)
	var cache = cache.NewBlogCache(config.REDIS)
	return BlogServiceImpl{repository: repository, cache: cache, search: config.SEARCH}
}
