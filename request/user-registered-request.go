package request

import (
	"errors"
	"gin-demo/models"
	"gin-demo/utils"
	"time"
)

type UserRequest struct {
	Username string `json:"username"`
	NickName string `json:"nickName"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Icon     string `json:"icon"`
	Code     string `json:"code"`
}

func (user UserRequest) Check() error {
	if user.Username == "" || len(user.Username) < 8 || len(user.Username) > 16 {
		return errors.New("用户名要在8-16个字符之间")
	} else if user.NickName == "" || (len(user.NickName) < 1 && len(user.NickName) > 20) {
		return errors.New("用户名要在1-20个字符之间")
	} else if user.Password == "" || len(user.Password) < 8 || len(user.Password) > 16 {
		return errors.New("密码要在8-16个字符之间")
	} else if user.Email == "" || !utils.IsEmailValid(user.Email) {
		return errors.New("邮箱格式不正确")
	} else if user.Icon == "" || !utils.IsImageURL(user.Icon) {
		return errors.New("这不是一个有效的图片链接")
	}
	return nil
}

func (user UserRequest) ToUserDo() models.User {
	var password = utils.EncryptPassword(user.Password)
	var now = time.Now()
	return models.User{
		Username: user.Username,
		Nickname: user.NickName,
		Password: password,
		Email:    user.Email,
		CreateAt: now,
		UpdateAt: now,
		RoleId:   models.ROLE_USER.Id,
		Role:     models.Role{},
	}
}
