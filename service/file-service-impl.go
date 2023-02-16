package service

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"os"
	"reflect"
	"vs-blog-api/common"
	"vs-blog-api/config"
	"vs-blog-api/modal"
	"vs-blog-api/response"
)

type FileServiceImpl struct {
}

func (f FileServiceImpl) DeleteFiles(ids []string) {
	f.DeleteUserFiles(-1, ids)
}

func (f FileServiceImpl) DeleteUserFiles(userId int, ids []string) {

	if len(ids) == 0 {
		panic(response.NewGlobalException(response.ParamsError, "找不到该ID"))
	}

	type Path struct {
		Id           string `json:"id"`
		AbsolutePath string `json:"absolute_path"`
	}

	if len(ids) == 1 {
		do, err := config.ES.Get().Index(common.FileIndex).Id(ids[0]).Do(context.Background())

		if err != nil {
			panic(response.NewGlobalException(response.ERROR, "删除失败"))
		}

		var file Path

		err = json.Unmarshal(do.Source, &file)

		if file.Id != "" {

			err := os.Remove(file.AbsolutePath)

			if err == nil {
				_, err := config.ES.Delete().Index(common.FileIndex).Id(file.Id).Refresh("true").Do(context.Background())

				if err != nil {
					panic(response.NewGlobalException(response.ERROR, "删除失败"))
				}

				return
			}
		} else {
			panic(response.NewGlobalException(response.NOTFOUND, "该文件不存在"))
		}

	}

	boolQuery := elastic.NewBoolQuery()

	if userId > 0 {
		boolQuery.Must(elastic.NewTermQuery("user_id", userId))
	}

	boolQuery.Must(elastic.NewIdsQuery().Ids(ids...))

	do, err := config.ES.Search().Index(common.FileIndex).Query(boolQuery).Do(context.Background())

	if err != nil {
		panic(response.NewGlobalException(response.ERROR, "删除失败"))
	}

	if do.TotalHits() == 0 {
		panic(response.NewGlobalException(response.NOTFOUND, "未找到相关文件"))
	}

	bulk := config.ES.Bulk()

	var deleteRequest []elastic.BulkableRequest

	for _, result := range do.Each(reflect.TypeOf(Path{})) {

		var file = result.(Path)

		if file.Id != "" {

			err := os.Remove(file.AbsolutePath)

			if err == nil {
				deleteRequest = append(deleteRequest, elastic.NewBulkDeleteRequest().Id(file.Id))
			}
		}
	}

	_, err3 := bulk.Index(common.FileIndex).Add(deleteRequest...).Refresh("true").Do(context.Background())

	if err3 != nil {
		panic(response.NewGlobalException(response.ERROR, "删除失败"))
	}
}

func (f FileServiceImpl) GetPublicFiles(page int, sort string, flag bool, keyword string) response.PageInfoResponse {
	return f.GetCurrentFiles(-1, page, sort, flag, keyword)
}

func (f FileServiceImpl) GetCurrentFiles(userId int, page int, sort string, flag bool, keyword string) response.PageInfoResponse {

	boolQuery := elastic.NewBoolQuery()

	if userId == -1 {
		boolQuery.Must(elastic.NewTermQuery("public", true))
	} else {
		boolQuery.Must(elastic.NewTermQuery("user_id", userId))
	}

	if keyword != "" {
		if flag {
			boolQuery.Must(elastic.NewTermQuery("suffix", keyword))
		} else {
			boolQuery.Must(elastic.NewMatchQuery("old_name", keyword))
		}
	}

	query := config.ES.Search().Index(common.FileIndex)

	if sort == "date" {
		query.SortBy(elastic.NewFieldSort("date").Desc())
	} else if sort == "size" {
		query.SortBy(elastic.NewFieldSort("size").Desc())
	} else {
		query.SortBy(elastic.NewFieldSort("date").Desc())
	}

	result, err := query.From(page - 1).Size(common.FileCount).Query(boolQuery).Do(context.Background())

	if err != nil {
		panic(response.NewGlobalException(response.ParamsError, "搜索文件失败,后台接口出错"))
	}

	hits := result.Hits

	var fileVo []modal.FileInfoVo

	var pageInfo response.PageInfoResponse

	pageInfo.Page = page

	pageInfo.Size = common.FileCount

	pageInfo.Total = hits.TotalHits.Value

	for _, hit := range hits.Hits {
		var fileInfo modal.EsFile
		json.Unmarshal(hit.Source, &fileInfo)
		fileInfo.Id = hit.Id
		fileVo = append(fileVo, fileInfo.ToVo())
	}

	pageInfo.Data = fileVo

	return pageInfo
}

func NewFileServiceImpl() FileService {

	return FileServiceImpl{}
}
