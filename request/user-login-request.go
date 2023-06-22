package request

import "errors"

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Code     string `json:"code"`
}

func (user UserLoginRequest) Check() error {
	if user.Username == "" {
		return errors.New("账号不能为空")
	} else if user.Password == "" {
		return errors.New("密码不能为空")
	} else if user.Code == "" {
		return errors.New("验证码不能为空")
	}
	return nil
}
