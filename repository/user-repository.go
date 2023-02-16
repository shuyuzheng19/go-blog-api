package repository

import "vs-blog-api/modal"

type UserRepository interface {
	SaveUser(user modal.User) error
	FindByUsername(username string) (modal.User, error)
	UpdateUser(user modal.User) error
	DeleteUserById(id int) error
}
