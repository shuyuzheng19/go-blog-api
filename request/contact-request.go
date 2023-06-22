package request

import (
	"errors"
	"gin-demo/utils"
)

type ContactRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Subject string `json:"subject"`
	Content string `json:"content"`
}

func (contact ContactRequest) Check() error {
	if contact.Email == "" || !utils.IsEmailValid(contact.Email) {
		return errors.New("错误的邮箱格式")
	} else if contact.Name == "" {
		return errors.New("请输入你的名称")
	} else if contact.Subject == "" {
		return errors.New("请输入主题内容")
	} else if contact.Content == "" {
		return errors.New("请输入消息内容")
	}
	return nil
}
