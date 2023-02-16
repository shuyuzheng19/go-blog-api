package utils

import (
	"encoding/json"
	"github.com/gocolly/colly"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
	"vs-blog-api/common"
	"vs-blog-api/response"
)

func PanicError(err error, message string) {

	if err != nil {
		panic(response.GlobalException{Code: response.ERROR, Message: message})
	}

}

// 验证是否是邮箱
func IsEmail(email string) bool {

	result, _ := regexp.MatchString(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`, email)

	return result
}

// 生成随机数
func CreateRandomNumber(size int) string {
	var numbers string
	for i := 0; i < size; i++ {
		numbers += strconv.Itoa(rand.Intn(9))
	}
	return numbers
}

// 获取文件名的后缀
func GetFileNameExt(filename string) string {

	ext := filename[strings.LastIndex(filename, ".")+1:]

	return ext
}

// 判断文件是否属于数组的某个类型
func IsAssignFile(filename string, fileTypes []string) bool {
	if filename == "" {
		return false
	}

	ext := GetFileNameExt(filename)

	for _, imageType := range fileTypes {
		if strings.ToLower(ext) == imageType {
			return true
		}
	}

	return false
}

const FORMAT_DATE = "2006-01-02"

const FORMAT_DATE_SIMPLE_TIME = "2006年 01月 02日 15:04"

const FORMAT_DATE_TIME = "2006-01-02 15:04"

// 将时间格式化
func FormatDate(date time.Time, format string) string {

	if format == "" {
		format = FORMAT_DATE_TIME
	}

	return date.Format(format)
}

func FormatDate2(s int64) string {
	format := time.Unix(s, 0).Format(FORMAT_DATE_TIME)
	return format
}

// 将数据序列化string类型
func ObjectToJson(data interface{}) string {

	result, err := json.Marshal(&data)

	if err != nil {
		return ""
	}

	return string(result)

}

// 将时间转成字符串
func DateToString(date time.Time) string {

	gap := date.Unix()

	now := time.Now().Unix()

	var second = now - gap

	var dateStr string

	if second <= 60 {
		dateStr = "刚刚"
	} else if second > 60 && second <= 60*60 {
		dateStr = strconv.Itoa(int(second/60)) + "分钟前"
	} else if second > 60*60 && second <= 60*60*24 {
		dateStr = strconv.Itoa(int(second/60/60)) + "小时前"
	} else if second > 60*60*24 && second <= 60*60*24*30 {
		dateStr = strconv.Itoa(int(second/60/60/24)) + "天前"
	} else if second > 60*60*24*30 && second <= 60*60*24*30*12 {
		dateStr = strconv.Itoa(int(second/60/60/24/30)) + "月前"
	} else {
		dateStr = strconv.Itoa(int(second/60/60/24/(30*12))) + "年前"
	}

	return dateStr
}

func JsonToObject(str string) interface{} {
	var result interface{}
	err := json.Unmarshal([]byte(str), &result)

	if err != nil {
		return nil
	}

	return result
}

func GetColly(url string) *colly.Collector {
	c := colly.NewCollector(colly.UserAgent(common.UserAgent))
	c.Async = true
	c.AllowURLRevisit = true
	c.Visit(url)
	return c
}

// 字符串转int
func ToInt(str string) int {

	result, err := strconv.Atoi(str)

	if err != nil {
		return -1
	}

	return result

}

// 字符串转int64
func ToInt64(str string) int64 {

	result, err := strconv.ParseInt(str, 10, 64)

	if err != nil {
		return -1
	}

	return result

}
