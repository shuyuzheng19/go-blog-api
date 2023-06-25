package utils

import (
	"encoding/json"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"io"
	"net/http"
	"strings"
)

func GetIPAddress(request *http.Request) string {
	ipAddress := request.Header.Get("X-Forwarded-For")
	if ipAddress == "" || strings.ToLower(ipAddress) == "unknown" {
		ipAddress = request.Header.Get("Proxy-Client-IP")
	}
	if ipAddress == "" || strings.ToLower(ipAddress) == "unknown" {
		ipAddress = request.Header.Get("WL-Proxy-Client-IP")
	}
	if ipAddress == "" || strings.ToLower(ipAddress) == "unknown" {
		ipAddress = request.RemoteAddr
	}
	return ipAddress
}

var ip_file_path = "ip2region.xdb"

var searcher *xdb.Searcher

func getSearcher() *xdb.Searcher {
	var search, err = xdb.NewWithFileOnly(ip_file_path)
	if err != nil {
		return nil
	}
	return search
}

func init() {
	searcher = getSearcher()
}

func GetIpCity(ip string) string {
	region, err := searcher.SearchByStr(ip)
	if err != nil {
		return "未知"
	}

	var split = strings.Split(region, "|")

	return strings.ReplaceAll(split[0]+" "+split[2]+" "+split[3], "0", "")
}

type Location struct {
	Country     string `json:"country"`
	CountryCode string `json:"countryCode"`
	Region      string `json:"region"`
	RegionName  string `json:"regionName"`
	City        string `json:"city"`
	Timezone    string `json:"timezone"`
	Query       string `json:"query"`
}

func GetIpCityApi(ip string) Location {
	var api = "http://ip-api.com/json/" + ip + "?lang=zh-CN"

	var response, _ = http.Get(api)

	var result Location

	var body = response.Body

	var buff, _ = io.ReadAll(body)

	json.Unmarshal(buff, &result)

	return result
}
