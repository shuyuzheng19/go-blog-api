package request

import (
	"errors"
	"gin-demo/models"
	"gin-demo/utils"
	"time"
	"unicode/utf8"
)

type TopicRequest struct {
	Name        string `json:"name"`
	Cover       string `json:"cover"`
	Description string `json:"desc"`
}

func (t TopicRequest) Check() error {
	var nameLen = utf8.RuneCountInString(t.Name)
	var descriptionLen = utf8.RuneCountInString(t.Description)
	if nameLen < 1 || nameLen > 15 {
		return errors.New("专题名称不能小于1个字符或则和大于15个字符")
	} else if descriptionLen < 1 || descriptionLen > 150 {
		return errors.New("专题简介不能小于1个字符或则和大于150个字符")
	} else if !utils.IsImageURL(t.Cover) {
		return errors.New("这不是一个图片链接")
	}
	return nil
}

func (t TopicRequest) ToDo(uid int) models.Topic {
	return models.Topic{
		Name:        t.Name,
		Description: t.Description,
		Cover:       t.Cover,
		CreateAt:    time.Now(),
		UserId:      uid,
	}
}
