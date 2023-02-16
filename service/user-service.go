package service

import (
	"vs-blog-api/modal"
	"vs-blog-api/response"
)

type UserService interface {
	RegisteredUser(user modal.User)
	Login(user modal.UserLoginRequest) response.Token
	Logout(username string)
}
