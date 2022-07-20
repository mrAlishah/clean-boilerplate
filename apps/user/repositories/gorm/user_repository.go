package repositories

import (
	errors2 "boilerplate/core/errors"
	"boilerplate/core/infrastructures"
	"boilerplate/core/models"
	"boilerplate/core/utils"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

// UserRepository -> database structure
type UserRepository struct {
	db     *infrastructures.GormDB
	logger *infrastructures.Logger
}

// NewUserRepository -> creates a new User repository
func NewUserRepository(db *infrastructures.GormDB, logger *infrastructures.Logger) *UserRepository {
	return &UserRepository{
		db:     db,
		logger: logger,
	}
}

// Save -> User
func (r UserRepository) Create(User *models.User) error {
	return r.db.DB.Create(User).Error
}

func (r UserRepository) FindByField(field string, value interface{}) (user models.User, err error) {
	err = r.db.DB.Where(fmt.Sprintf("%s= ?", field), value).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors2.NotFoundError
	}
	return
}

func (r UserRepository) DeleteByID(id uint) (err error) {
	var user models.User
	err = r.db.DB.Where("id=?", id).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors2.NotFoundError
		return
	}
	if err != nil {
		return
	}
	err = r.db.DB.Where("id=?", id).Delete(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors2.NotFoundError
	}
	return
}

func (r UserRepository) IsExist(field string, value string) (bool, error) {
	_, err := r.FindByField(field, value)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	return false, err
}

// GetAllUser -> Get All users
func (r UserRepository) GetAllUsers(pagination utils.Pagination) ([]models.User, int64, error) {
	var users []models.User
	var totalRows int64 = 0
	queryBuilder := r.db.DB.Limit(pagination.PageSize).Offset(pagination.Offset).Order("created_at desc")
	queryBuilder = queryBuilder.Model(&models.User{})

	if pagination.Keyword != "" {
		searchQuery := "%" + pagination.Keyword + "%"
		queryBuilder.Where(r.db.DB.Where("`users`.`name` LIKE ?", searchQuery))
	}

	err := queryBuilder.
		Find(&users).
		Offset(-1).
		Limit(-1).
		Count(&totalRows).Error
	return users, totalRows, err
}

//update a single column by user model
func (r UserRepository) UpdateColumn(user *models.User, column string, value interface{}) error {
	return r.db.DB.Model(user).Update(column, value).Error
}

func (r UserRepository) UpdateModel(user *models.User, id uint64) error {
	return r.db.DB.Where("id=?", id).Updates(&user).Error
}
