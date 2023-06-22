package repository

import (
	"gin-demo/models"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func (u UserRepositoryImpl) ExistsByUsername(username string) (count int64) {
	u.db.Model(&models.User{}).Where("username = ?", username).Count(&count)
	return count
}

func (u UserRepositoryImpl) FindByUsername(username string) (user models.User) {
	u.db.Model(&models.User{}).Preload("Role").First(&user, "username = ?", username)
	return user
}

func (u UserRepositoryImpl) FindAll() (users []models.User) {
	u.db.Model(&models.User{}).Preload("Role").Find(&users)
	return users
}

func (u UserRepositoryImpl) FindById(id int) (user models.User) {
	u.db.Model(&models.User{}).First(&user, "id = ?", id)
	return user
}

func (u UserRepositoryImpl) Update(user models.User) error {
	return u.db.Model(&models.User{}).UpdateColumns(&user).Error
}

func (u UserRepositoryImpl) Save(user models.User) error {
	return u.db.Model(&models.User{}).Save(&user).Error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepositoryImpl{db: db}
}
