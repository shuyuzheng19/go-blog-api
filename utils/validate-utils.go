package utils

import (
	"net/mail"
	"path/filepath"
	"regexp"
	"strings"
)

func IsEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func IsImageURL(url string) bool {
	imageRegex := `^https?://.*\.(png|jpe?g|gif|svg|ico)$`

	regex := regexp.MustCompile(imageRegex)

	return regex.MatchString(url)
}

func IsImageFile(filename string) bool {
	// 将文件名转换为小写字母，并获取文件扩展名
	ext := strings.ToLower(filepath.Ext(filename))

	// 检查文件扩展名是否为图片格式
	allowedExtensions := []string{".jpg", ".jpeg", ".png", ".gif"}
	for _, allowedExt := range allowedExtensions {
		if ext == allowedExt {
			return true
		}
	}

	return false
}
