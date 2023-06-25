package request

import (
	"errors"
	"gin-demo/models"
	"strconv"
	"time"
)

type CommentRequest struct {
	BlogId     int    `json:"blogId"`
	ParentId   string `json:"parentId"`
	Content    string `json:"content"`
	ContentImg string `json:"contentImg"`
}

func (comment CommentRequest) Check() error {
	if comment.BlogId <= 0 {
		return errors.New("缺少博客ID")
	} else if comment.Content == "" {
		return errors.New("请输入评论内容")
	}
	return nil
}

func (comment CommentRequest) ToDo() models.Comment {
	var parentId, err = strconv.Atoi(comment.ParentId)
	var comment2 = models.Comment{
		Content:    comment.Content,
		Likes:      0,
		CreateTime: time.Now(),
		BlogId:     comment.BlogId,
		ContentImg: comment.ContentImg,
	}
	if err != nil {
		comment2.ParentId = nil
	} else {
		var parent = int64(parentId)
		comment2.ParentId = &parent
	}
	return comment2
}
