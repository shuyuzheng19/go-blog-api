package repository

import (
	"gorm.io/gorm"
)

type BaseDao[T any] struct {
	t  *T
	db *gorm.DB
}

func (dao BaseDao[T]) Save(t T) error {
	var err = dao.db.Model(dao.t).Save(&t).Error
	return err
}

func (dao BaseDao[T]) UpdateById(id int, t T) error {
	var err = dao.db.Model(dao.t).Where("id = ?", id).UpdateColumns(&t).Error
	return err
}

func (dao BaseDao[T]) DeleteById(id int) error {
	var err = dao.db.Model(dao.t).Delete("id = ?", id).Error
	return err
}

func (dao BaseDao[T]) FindById(id int) T {
	var result T

	dao.db.Model(dao.t).Find(&result, "id = ?", id)

	return result
}

func (dao BaseDao[T]) FindAll() []T {
	var result = make([]T, 0)

	dao.db.Model(dao.t).Find(&result)

	return result
}

func NewBaseDao[T any](db *gorm.DB) *BaseDao[T] {
	var t T
	return &BaseDao[T]{t: &t, db: db}
}
