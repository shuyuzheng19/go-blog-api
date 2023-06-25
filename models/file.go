package models

import (
	"fmt"
	"gin-demo/response"
	"gin-demo/vo"
	"gorm.io/gorm"
	"time"
)

type FileInfo struct {
	ID           int             `gorm:"primaryKey"`
	UserID       int             `gorm:"column:user_id"`
	OldName      string          `gorm:"column:old_name"`
	NewName      string          `gorm:"column:new_name"`
	Date         time.Time       `gorm:"column:date"`
	Size         int64           `gorm:"column:size"`
	Suffix       string          `gorm:"column:suffix"`
	URL          string          `gorm:"column:url"`
	AbsolutePath string          `gorm:"column:absolute_path"`
	IsPublic     bool            `gorm:"column:is_public"`
	DeletedAt    *gorm.DeletedAt `gorm:"index"`
	FileMd5Id    string          `gorm:"column:md5;type:text"`
	FileMd5      FileMd5         `gorm:"foreignKey:FileMd5Id"`
}

type FileMd5 struct {
	Md5 string `gorm:"primary_key;not null;unique"`
	Url string `gorm:"column:url;not null"`
}

func (file FileInfo) ToVo() vo.FileVo {
	return vo.FileVo{
		ID:      file.ID,
		Name:    file.OldName,
		DateStr: response.FormatTime(file.Date),
		Suffix:  file.Suffix,
		SizeStr: GetSizeStr(float64(file.Size)),
		MD5:     file.FileMd5.Md5,
		URL:     file.FileMd5.Url,
	}
}

func GetSizeStr(size float64) string {
	if size == 0 {
		return "0 B"
	}

	var sizeStr string

	if size < 1024 {
		sizeStr = fmt.Sprintf("%.0f", size)
		return sizeStr + " BIT"
	} else if size > 1024 && size < 1024*1024 {
		sizeStr = fmt.Sprintf("%.2f", size/1024)
		return sizeStr + " KB"
	} else if size > 1024*1024 && size < 1024*1024*1024 {
		sizeStr = fmt.Sprintf("%.2f", size/(1024*1024))
		return sizeStr + " MB"
	} else if size > 1024*1024*1024 && size < 1024*1024*1024*1024 {
		sizeStr = fmt.Sprintf("%.2f", size/(1024*1024*1024))
		return sizeStr + " GB"
	} else if size > 1024*1024*1024*1024 {
		sizeStr = fmt.Sprintf("%.2f", size/(1024*1024*1024*1024))
		return sizeStr + " TB"
	} else {
		return "未知"
	}
}

func (FileInfo) TableName() string {
	return "files"
}
