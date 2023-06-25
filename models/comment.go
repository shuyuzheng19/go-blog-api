package models

import (
	"gin-demo/response"
	"gin-demo/vo"
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	Id         int64           `gorm:"primary_key"`
	ParentId   *int64          `gorm:"parent_id"`
	Ip         string          `gorm:"column:ip;not null"`
	Address    string          `gorm:"column:address;not null"`
	Content    string          `gorm:"column:content;not null"`
	Likes      int64           `gorm:"default:0"`
	CreateTime time.Time       `gorm:"column:create_time"`
	BlogId     int             `gorm:"column:blog_id"`
	UserAgent  string          `gorm:"column:user_agent"`
	ContentImg string          `gorm:"column:content_img"`
	DeletedAt  *gorm.DeletedAt `gorm:"index"`
	UserId     int
	User       User
	Reply      []Comment `gorm:"foreignkey:ParentId"`
}

func (comment Comment) ToVo() vo.CommentVo {

	var reply vo.ReplyVo

	var replyCount = len(comment.Reply)

	if replyCount > 0 {
		reply.Total = replyCount
		for _, c := range comment.Reply {
			reply.List = append(reply.List, c.ToVo())
		}
	} else {
		reply = vo.ReplyVo{
			Total: 0,
			List:  make([]vo.CommentVo, 0),
		}
	}

	return vo.CommentVo{
		ID:         comment.Id,
		ParentId:   comment.ParentId,
		User:       comment.User.ToCommentUserVo(),
		Address:    comment.Address,
		Content:    comment.Content,
		UID:        comment.User.Id,
		Likes:      comment.Likes,
		CreateTime: response.FormatTime(comment.CreateTime),
		ContentImg: comment.ContentImg,
		Reply:      reply,
	}
}
