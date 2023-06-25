package search

import (
	"bytes"
	"encoding/json"
	"gin-demo/common"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"strings"
)

type MeiliSearch struct {
	ApiHost string
	Hearers map[string]string
}

func (m MeiliSearch) CreateIndex(index string) error {
	var requestBody = "{ \"uid\": \"" + index + "\", \"primaryKey\": \"id\" }"
	var url = m.ApiHost + "/indexes/" + index
	var request, err = http.NewRequest("POST", url, strings.NewReader(requestBody))
	if err != nil {
		return errors.New("请求URL错误")
	}

	for header, value := range m.Hearers {
		request.Header.Add(header, value)
	}

	var client = &http.Client{}

	var response, err3 = client.Do(request)

	if err3 != nil {
		return errors.New("获取响应失败")
	}

	var statusCode = response.StatusCode

	if statusCode == 200 {
		return nil
	} else {
		return errors.New("请求响应错误")
	}
}

func (m MeiliSearch) DropIndex(index string) error {
	var url = m.ApiHost + "/indexes/" + index

	var request, err = http.NewRequest("DELETE", url, nil)

	if err != nil {
		return errors.New("请求URL错误")
	}

	for header, value := range m.Hearers {
		request.Header.Add(header, value)
	}

	var client = &http.Client{}

	var response, err3 = client.Do(request)

	if err3 != nil {
		return errors.New("获取响应失败")
	}

	var statusCode = response.StatusCode

	if statusCode == 200 {
		return nil
	} else {
		return errors.New("请求响应错误")
	}
}

func (m MeiliSearch) DeleteAllDocument(index string) error {
	var url = m.ApiHost + "/indexes/" + index + "/documents"

	var request, err = http.NewRequest("DELETE", url, nil)

	if err != nil {
		return errors.New("请求URL错误")
	}

	for header, value := range m.Hearers {
		request.Header.Add(header, value)
	}

	var client = &http.Client{}

	var response, err3 = client.Do(request)

	if err3 != nil {
		return errors.New("获取响应失败")
	}

	var statusCode = response.StatusCode

	if statusCode == 200 {
		return nil
	} else {
		return errors.New("请求响应错误")
	}

}

func (m MeiliSearch) SaveDocuments(index string, arrayJson string) error {
	var url = m.ApiHost + "/indexes/" + index + "/documents"

	var request, err = http.NewRequest("POST", url, strings.NewReader(arrayJson))

	if err != nil {
		return errors.New("请求URL错误")
	}

	for header, value := range m.Hearers {
		request.Header.Add(header, value)
	}

	var client = &http.Client{}

	var response, err3 = client.Do(request)

	if err3 != nil {
		return errors.New("获取响应失败")
	}

	var statusCode = response.StatusCode

	if statusCode == 200 {
		return nil
	} else {
		return errors.New("请求响应错误")
	}
}

func (m MeiliSearch) Search(index string, keyword string, page int) (result *SearchResponse, err error) {
	var jsonBody = NewSearchQueryBuilder().SetQ(keyword).SetOffset((page - 1) * common.PAGE_SIZE).SetAttributesToHighlight([]string{"*"}).SetLimit(common.PAGE_SIZE).SetShowMatchesPosition(false).SetHighlightPreTag(common.HIG_PRE).SetHighlightPostTag(common.HIG_SUFFIX).Build()

	var url = m.ApiHost + "/indexes/" + index + "/search"

	var buff, err1 = json.Marshal(&jsonBody)

	if err1 != nil {
		return nil, errors.New("错误的请求体")
	}

	var request, err2 = http.NewRequest("POST", url, bytes.NewReader(buff))

	if err2 != nil {
		return nil, errors.New("请求URL错误")
	}

	for header, value := range m.Hearers {
		request.Header.Add(header, value)
	}

	var client = &http.Client{}

	var response, err3 = client.Do(request)

	if err3 != nil {
		return nil, errors.New("获取响应失败")
	}

	var statusCode = response.StatusCode

	if statusCode == 200 {
		var responseBody, err4 = io.ReadAll(response.Body)

		if err4 != nil {
			return nil, errors.New("获取响应数据失败")
		}

		json.Unmarshal(responseBody, &result)

		return result, nil
	} else {
		return nil, errors.New("请求响应错误")
	}

}

func NewMeiliSearch(apiHost string, apiKey string) *MeiliSearch {
	var headers = make(map[string]string)
	headers["Authorization"] = "Bearer " + apiKey
	headers["Content-Type"] = "application/json;charset=utf-8"
	return &MeiliSearch{ApiHost: apiHost, Hearers: headers}
}
