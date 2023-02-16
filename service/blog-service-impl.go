package service

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/olivere/elastic/v7"
	"math"
	"strconv"
	"time"
	"vs-blog-api/common"
	"vs-blog-api/config"
	"vs-blog-api/dto"
	"vs-blog-api/modal"
	"vs-blog-api/repository"
	"vs-blog-api/response"
	"vs-blog-api/utils"
)

type BlogServiceImpl struct {
}

func (b BlogServiceImpl) GetUserLikeBlog(userId int, page int) response.PageInfoResponse {
	if userId <= 0 || page <= 0 {
		panic(response.NewGlobalException(response.ParamsError, "非法参数"))
	}

	var key = common.UserLikeSet + ":" + strconv.Itoa(userId)

	result := config.Redis.ZRevRangeByScore(key, redis.ZRangeBy{
		Min:    "0",
		Max:    strconv.Itoa(math.MaxInt),
		Offset: int64((page - 1) * common.PageSize),
		Count:  common.PageSize,
	}).Val()

	err, blogs := blogRepository.FindBlogIdIn(result)

	if err != nil {
		panic(response.NewGlobalException(response.ERROR, "查询失败"))
	}

	var count = config.Redis.ZCard(key).Val()

	var blogVos = BlogArrayToBlogVoArray(blogs)

	return response.PageInfoResponse{
		Page:  page,
		Size:  common.PageSize,
		Total: count,
		Data:  blogVos,
	}
}

func (b BlogServiceImpl) AddBlog(userId int, blogRequest modal.BlogRequest) {

	blogRequest.Validator()

	var tags = make([]modal.Tag, 0)

	if len(blogRequest.Tags) > 0 {
		for _, id := range blogRequest.Tags {
			tags = append(tags, modal.Tag{
				Id: id,
			})
		}
	}

	var now = time.Now()

	var blog = modal.Blog{
		Description: blogRequest.Description,
		Title:       blogRequest.Title,
		SourceUrl:   blogRequest.SourceUrl,
		CoverImage:  blogRequest.CoverImage,
		EyeCount:    0,
		LikeCount:   0,
		TopicId:     blogRequest.Topic,
		CategoryId:  blogRequest.Category,
		Tags:        tags,
		UserId:      userId,
		CreateTime:  now,
		UpdateTime:  now,
	}

	if blogRequest.Markdown == "true" {
		blog.Content = blogRequest.Content
		blog.Markdown = true
	} else {
		blog.Content = blogRequest.Content
		blog.Markdown = false
	}

	err, blog2 := blogRepository.SaveBlog(blog)

	config.Redis.SAdd(common.RandomBlog, utils.ObjectToJson(modal.SimpleBlog{
		Id:    blog2.Id,
		Title: blog2.Title,
	}))

	if err != nil {
		panic(response.NewGlobalException(response.ERROR, "添加博客失败!"))
	}

}

func (b BlogServiceImpl) SaveBlog(userId int, content string) {

	if userId <= 0 {
		panic(response.NewGlobalException(response.ParamsError, "非法参数"))
	}

	if content == "" {
		panic(response.NewGlobalException(response.ParamsError, "内容不能为空"))
	}

	var base64Content = base64.StdEncoding.EncodeToString([]byte(content))

	config.Redis.HSet(common.UserSaveBlogMap, strconv.Itoa(userId), base64Content)
}

func (b BlogServiceImpl) GetSaveBlog(userId int) string {
	content := config.Redis.HGet(common.UserSaveBlogMap, strconv.Itoa(userId)).Val()

	if content == "" {
		panic(response.NewGlobalException(response.NOTFOUND, "您之前没有保存博客哦!"))
	}

	decoderContent, err := base64.StdEncoding.DecodeString(content)

	if err != nil {
		panic(response.NewGlobalException(response.ERROR, "解码失败"))
	}

	return string(decoderContent)
}

// 将blog数组转为blogVo数组
func BlogArrayToBlogVoArray(blogs []modal.Blog) []modal.BlogVo {

	var blogVos []modal.BlogVo

	for _, blog := range blogs {
		blogVos = append(blogVos, blog.ToBlogVo())
	}

	return blogVos

}

func (b BlogServiceImpl) GetSimilarBlog(keyword string, blogId int) []modal.SimpleBlog {

	if keyword == "" {
		panic(response.NewGlobalException(response.ParamsError, "缺少关键字"))
	}

	if blogId <= 0 {
		panic(response.NewGlobalException(response.ParamsError, "非法ID"))
	}

	query := elastic.NewMoreLikeThisQuery().Field("title", "description").LikeText(keyword).MinWordLength(2).MinTermFreq(1)

	do, err := config.ES.Search().Index(common.BlogIndex).From(0).Size(common.SimilarBlogCount).Query(query).Do(context.Background())

	var blogVos = make([]modal.SimpleBlog, 0)

	if err != nil {
		return blogVos
	}

	for _, blog := range do.Hits.Hits {

		var simpleBlog modal.SimpleBlog

		json.Unmarshal(blog.Source, &simpleBlog)

		if simpleBlog.Id == blogId {
			continue
		}

		blogVos = append(blogVos, simpleBlog)

	}

	return blogVos
}

func (b BlogServiceImpl) GetUserTopBlog(id int) []modal.SimpleBlog {
	if id <= 0 {
		panic(response.NewGlobalException(response.ParamsError, "非法的用户ID"))
	}

	err, blogs, _ := blogRepository.FindBlogByUserId(id, 1, dto.SortMap[dto.EYE])

	if err != nil {
		return make([]modal.SimpleBlog, 0)
	}

	var simpleBlogs []modal.SimpleBlog

	for _, blog := range blogs {
		simpleBlogs = append(simpleBlogs, modal.SimpleBlog{
			Id:    blog.Id,
			Title: blog.Title,
		})
	}

	return simpleBlogs
}

func (b BlogServiceImpl) FindUserBlog(id int, page int) response.PageInfoResponse {
	if id <= 0 {
		panic(response.NewGlobalException(response.ParamsError, "非法的用户ID"))
	}

	err, blogs, count := blogRepository.FindBlogByUserId(id, page, dto.SortMap[dto.CREATE])

	if err != nil {
		return response.PageInfoResponse{}
	}

	var blogVos = BlogArrayToBlogVoArray(blogs)

	return response.PageInfoResponse{
		Page:  page,
		Size:  common.PageSize,
		Total: count,
		Data:  blogVos,
	}
}

func (b BlogServiceImpl) SearchBlog(keyword string, page int) response.PageInfoResponse {

	multiQuery := elastic.NewMultiMatchQuery(keyword, "title", "description")

	highlight := elastic.NewHighlight()

	titleHig := elastic.NewHighlighterField("title").PreTags("<span style='color:#ff2d51'>").PostTags("</span>")

	descHig := elastic.NewHighlighterField("description").PreTags("<span style='color:#ff2d51'>").PostTags("</span>")

	highlight.Fields(titleHig, descHig)

	result, err := config.ES.Search().Index(common.BlogIndex).From((page - 1) * 10).Size(10).Query(multiQuery).Highlight(highlight).Do(context.Background())

	if err != nil {
		panic(response.NewGlobalException(response.ERROR, "后台接口出错"))
	}

	hits := result.Hits

	if len(hits.Hits) == 0 {
		panic(response.NewGlobalException(response.NOTFOUND, "未找到相关数据"))
	}

	var blogs []modal.EsBlog

	for _, hit := range hits.Hits {

		var esBlog modal.EsBlog

		json.Unmarshal(hit.Source, &esBlog)

		titleHtml := hit.Highlight["title"]

		if len(hit.Highlight["title"]) > 0 {
			esBlog.Title = titleHtml[0]
		}

		descHtml := hit.Highlight["description"]
		if len(descHtml) > 0 {
			esBlog.Description = hit.Highlight["description"][0]
		}
		blogs = append(blogs, esBlog)
	}

	return response.PageInfoResponse{
		Page:  page,
		Size:  10,
		Total: hits.TotalHits.Value,
		Data:  blogs,
	}

}

func (b BlogServiceImpl) RangeDateBlog(start int64, end int64, page int) []modal.RangeBlog {
	err, blogs := blogRepository.FindRangeDate(start, end, page, dto.SortMap[dto.CREATE])

	if err != nil {
		return []modal.RangeBlog{}
	}

	return blogs
}

func (b BlogServiceImpl) GetRandomBlogs() []modal.SimpleBlog {
	var KEY = common.RandomBlog

	var exists = config.Redis.Exists(KEY).Val()

	if exists == 0 {

		err, blogs := blogRepository.FindAllIdAndTitle()

		if err != nil {
			return []modal.SimpleBlog{}
		}

		for _, blog := range blogs {
			config.Redis.SAdd(KEY, utils.ObjectToJson(blog))
		}

	}

	var simpleBlog []modal.SimpleBlog

	blogs := config.Redis.SRandMemberN(KEY, 10).Val()

	for _, blogStr := range blogs {

		var blog modal.SimpleBlog

		json.Unmarshal([]byte(blogStr), &blog)

		simpleBlog = append(simpleBlog, modal.SimpleBlog{
			Id:    blog.Id,
			Title: blog.Title,
		})
	}

	return simpleBlog

}

func NewBlogService() BlogService {

	blogRepository = repository.NewBlogRepositoryImpl()

	return BlogServiceImpl{}
}

func (b BlogServiceImpl) FindByIdBlog(id int) modal.Blog {

	if id <= 0 {
		panic(response.NewGlobalException(response.ParamsError, "非法参数"))
	}

	var key = strconv.Itoa(id)

	var redisBlog = config.Redis.HGet(common.BlogMap, key).Val()

	if redisBlog == "" {
		err, blog := blogRepository.FindById(id)

		config.Redis.HSet(common.BlogMap, key, utils.ObjectToJson(blog))

		if err != nil || blog.Id == 0 {
			panic(response.NewGlobalException(response.NOTFOUND, "该博客不存在或者已被删除"))
		}

		return blog

	} else {

		var blog modal.Blog

		if err := json.Unmarshal([]byte(redisBlog), &blog); err == nil {
			if blog.Id == 0 {
				panic(response.NewGlobalException(response.NOTFOUND, "该博客不存在或者已被删除"))
			}
			return blog
		}

	}

	return modal.Blog{}

}

func (b BlogServiceImpl) GetHotBlogs() []modal.SimpleBlog {

	list := config.Redis.LRange(common.HotBlog, 0, -1).Val()

	var blogs []modal.SimpleBlog

	if len(list) == 0 {
		var sortRequest = dto.NewBlogPageSortDto(1, dto.EYE, -1)

		result, err, _ := blogRepository.FindAll(sortRequest)

		if err != nil {
			panic(response.NewGlobalException(response.ERROR, "后台接口错误,获取热门博客失败"))
		}
		var blogsStr []string

		for _, blog := range result {
			simpleBlog := modal.SimpleBlog{
				Id:    blog.Id,
				Title: blog.Title,
			}

			blogs = append(blogs, simpleBlog)

			blogsStr = append(blogsStr, utils.ObjectToJson(simpleBlog))
		}
		config.Redis.Del(common.HotBlog)
		config.Redis.RPush(common.HotBlog, blogsStr)

	} else {
		for _, str := range list {
			var simpleBlog modal.SimpleBlog
			if err := json.Unmarshal([]byte(str), &simpleBlog); err == nil {
				blogs = append(blogs, simpleBlog)
			}
		}

	}

	return blogs
}

var blogRepository repository.BlogRepository

func (b BlogServiceImpl) FindBlogs(request dto.BlogPageSortDto) response.PageInfoResponse {

	fmt.Println(request)

	if request.Sort == "" {
		request.Sort = dto.SortMap[dto.CREATE]
	}

	blogs, err, count := blogRepository.FindAll(request)

	if err != nil {
		panic(response.NewGlobalException(response.ERROR, "后台出现错误,查询博客失败"))
	}

	var blogVos = BlogArrayToBlogVoArray(blogs)

	return response.PageInfoResponse{
		Page:  request.Page,
		Size:  request.Size,
		Total: count,
		Data:  blogVos,
	}
}
