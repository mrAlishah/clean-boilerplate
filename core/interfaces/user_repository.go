package interfaces

import (
	"boilerplate/core/models"
	"boilerplate/core/utils"
)

type UserRepository interface {
	Create(user *models.User) error
	FindByField(field string, value interface{}) (user models.User, err error)
	DeleteByID(id uint) error
	IsExist(field string, value string) (bool, error)
	GetAllUsers(pagination utils.Pagination) ([]models.User, int64, error)
	UpdateColumn(user *models.User, column string, value interface{}) error
	UpdateModel(user *models.User, id uint64) error
}
