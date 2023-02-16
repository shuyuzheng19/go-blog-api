package dto

import "vs-blog-api/common"

type BlogPageSortDto struct {
	Page   int
	Size   int
	Sort   string
	SortId int
}

func NewBlogPageSortDto(page int, sort string, sortId int) BlogPageSortDto {

	return BlogPageSortDto{
		Page:   page,
		Size:   common.PageSize,
		Sort:   sort,
		SortId: sortId,
	}

}

const (
	//创建时间排序
	CREATE = "CREATE"
	//修改时间排序
	UPDATE = "UPDATE"
	//浏览量排序
	EYE = "EYE"
	//点赞量排序
	LIKE = "LIKE"
	//时间倒序排序
	SORT = "SORT"
)

var SortMap = map[string]string{
	CREATE: "create_time desc",
	UPDATE: "update_time desc",
	EYE:    "eye_count desc",
	LIKE:   "like_count desc",
	SORT:   "create_time asc",
}
