package repository

import "gin-demo/models"

type UserRepository interface {
	Save(user models.User) error
	FindAll() (users []models.User)
	FindById(id int) (user models.User)
	FindByUsername(username string) (user models.User)
	ExistsByUsername(username string) (count int64)
	Update(user models.User) error
}
