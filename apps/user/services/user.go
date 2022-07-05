package services

import (
	repositories "boilerplate/apps/user/repositories/gorm"
	"boilerplate/core/infrastructures"
	"boilerplate/core/interfaces"
	"boilerplate/core/models"
	"boilerplate/core/utils"
)

// UserService -> struct
type UserService struct {
	repositories repositories.UserRepository
	db           infrastructures.GormDB
	logger       interfaces.Logger
}

// NewUserService -> creates a new Userservice
func NewUserService(repositories repositories.UserRepository, db infrastructures.GormDB, logger *infrastructures.Logger) UserService {
	return UserService{
		repositories: repositories,
		db:           db,
		logger:       logger,
	}
}

// GetAllUser -> call to get all the User
func (us UserService) GetAllUsers(pagination utils.Pagination) (users []models.User, count int64, err error) {
	users, count, err = us.repositories.GetAllUsers(pagination)
	if err != nil {
		us.logger.Fatal("Failed to get users", err.Error())
		return
	}
	return
}
