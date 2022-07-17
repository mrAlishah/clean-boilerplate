package mocks

import (
	"boilerplate/core/models"
	"boilerplate/core/models/faker"
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

func (r *UserRepository) FindByField(field string, value interface{}) (user models.User, err error) {
	return r.FindByFieldFn(field, value)
}

func (r *UserRepository) DeleteByID(id uint) error {
	return r.DeleteByID(id)
}

func (r *UserRepository) IsExist(field string, value string) (bool, error) {
	return r.IsExistFn(field, value)
}

func (r *UserRepository) GetAllUsers(pagination utils.Pagination) ([]models.User, int64, error) {
	return r.GetAllUsersFn(pagination)
}

func (r *UserRepository) UpdateColumn(user *models.User, column string, value interface{}) error {
	return r.UpdateColumnFn(user, column, value)
}

func (r *UserRepository) DeleteUser(id uint64) (err error) {
	return r.DeleteUserFn(id)
}

func NewUserRepository() *UserRepository {
	userFaker := faker.User{}
	return &UserRepository{
		CreateFn: func(user *models.User) error {
			return nil
		},
		FindByFieldFn: func(field string, value interface{}) (user models.User, err error) {
			return userFaker.CreateOne(), nil
		},
		DeleteByIDFn: func(id uint) error {
			return nil
		},
		UpdateColumnFn: func(user *models.User, column string, value interface{}) error {
			return nil
		},
		IsExistFn: func(field string, value string) (bool, error) {
			return true, nil
		},
		GetAllUsersFn: func(pagination utils.Pagination) ([]models.User, int64, error) {
			return userFaker.CreateMany(5), 5, nil
		},
	}
}
