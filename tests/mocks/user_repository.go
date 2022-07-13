package mocks

import (
	"boilerplate/core/models"
	"boilerplate/core/utils"
)

type UserRepository struct {
	CreateFn       func(user *models.User) error
	FindByFieldFn  func(field string, value interface{}) (user models.User, err error)
	DeleteByIDFn   func(id uint) error
	IsExistFn      func(field string, value string) (bool, error)
	GetAllUsersFn  func(pagination utils.Pagination) ([]models.User, int64, error)
	UpdateColumnFn func(user *models.User, column string, value interface{}) error
	DeleteUserFn   func(id uint64) (err error)
}

func (r *UserRepository) Create(user *models.User) error {
	return r.CreateFn(user)
}

func (r *UserRepository) FindByField() {

}
