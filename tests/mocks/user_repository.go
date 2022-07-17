package mocks

import (
	"boilerplate/core/infrastructures"
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

func (r *UserRepository) FindByField(field string, value interface{}) (user models.User, err error) {
	return r.FindByFieldFn(field, value)
}

func (r *UserRepository) DeleteByID(id uint) error {
	return r.DeleteByID(id)
}

func (r *UserRepository) IsExist(field string, value string) (bool, error) {
	return r.IsExistFn(field, value)
}

func (r *UserRepository) GetAllUser(pagination utils.Pagination) ([]models.User, int64, error) {
	return r.GetAllUsersFn(pagination)
}

func (r *UserRepository) UpdateColumn(user *models.User, column string, value interface{}) error {
	return r.UpdateColumnFn(user, column, value)
}

func (r *UserRepository) DeleteUser(id uint64) (err error) {
	return r.DeleteUserFn(id)
}

func NewUserRepository() UserRepository {
	env := infrastructures.NewEnv()
	logger := infrastructures.NewLogger(env)
	encryption := infrastructures.NewEncryption(logger, env)
	return UserRepository{
		CreateFn: func(user *models.User) error {
			return nil
		},
		FindByFieldFn: func(field string, value interface{}) (user models.User, err error) {
			return models.User{
				FirstName: "dgsgd",
				LastName:  "dsgdssdg",
				Email:     "mahdi@gmail.com",
				Password:  encryption.SaltAndSha256Encrypt("m1234567", "m12345567"),
			}, err
		},
	}
}
