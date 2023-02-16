package modal

import (
	"fmt"
	"strconv"
	"time"
	"vs-blog-api/common"
)

type EsFile struct {
	Id           string    `json:"id"`
	UserId       int       `json:"user_id"`
	OldName      string    `json:"old_name"`
	NewName      string    `json:"new_name"`
	Date         time.Time `json:"date"`
	Size         int64     `json:"size"`
	Suffix       string    `json:"suffix"`
	Url          string    `json:"url"`
	AbsolutePath string    `json:"absolute_path"`
	Path         string    `json:"path"`
	Public       bool      `json:"public"`
	First        bool      `json:"first"`
}

type FileInfoVo struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	DateStr string `json:"dateStr"`
	Suffix  string `json:"suffix"`
	SizeStr string `json:"sizeStr"`
	Url     string `json:"url"`
}

func (file EsFile) ToVo() FileInfoVo {

	return FileInfoVo{
		Id:      file.Id,
		Name:    file.OldName,
		Suffix:  file.Suffix,
		DateStr: file.Date.Format("2006-01-02 15:04"),
		SizeStr: GetSizeStr(float64(file.Size)),
		Url:     file.Url,
	}
}

func GetSizeStr(size float64) string {
	if size == 0 {
		return "0 B"
	}
	var sizeStr = strconv.Itoa(int(size))

	if size < 1024 {
		return sizeStr + " BIT"
	} else if size > 1024 && size < common.MB {
		return fmt.Sprintf("%.2f", size/float64(1024)) + " KB"
	} else if size > common.MB && size < common.GB {
		return fmt.Sprintf("%.2f", size/float64(1024*1024)) + " MB"
	} else if size > common.GB && size < common.TB {
		return fmt.Sprintf("%.2f", size/float64(1024*1024*1024)) + " GB"
	} else if size > common.TB {
		return fmt.Sprintf("%.2f", size/float64(1024*1024*1024*1024)) + " TB"
	} else {
		return "未知"
	}
	return sizeStr
}
