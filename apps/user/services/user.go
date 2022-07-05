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
	userRepository *repositories.UserRepository
	db             *infrastructures.GormDB
	logger         interfaces.Logger
	encryption     *infrastructures.Encryption
}

// NewUserService -> creates a new Userservice
func NewUserService(userRepository *repositories.UserRepository,
	db *infrastructures.GormDB, logger *infrastructures.Logger,
	encryption *infrastructures.Encryption) *UserService {
	return &UserService{
		userRepository: userRepository,
		db:             db,
		logger:         logger,
		encryption:     encryption,
	}
}

// GetAllUser -> call to get all the User
func (us UserService) GetAllUsers(pagination utils.Pagination) (users []models.User, count int64, err error) {
	users, count, err = us.userRepository.GetAllUsers(pagination)
	if err != nil {
		us.logger.Fatal("Failed to get users", err.Error())
		return
	}
	return
}

func (us UserService) CreateUser(userData models.CreateUserRequestAdmin) (err error) {
	encryptedPassword := us.encryption.SaltAndSha256Encrypt(userData.Password)
	user := models.User{
		Password:  encryptedPassword,
		FirstName: userData.FirstName,
		LastName:  userData.LastName,
		Email:     userData.Email,
	}
	err = us.userRepository.Create(&user)
	if err != nil {
		us.logger.Fatal("Failed to create user:%s", err.Error())
		return
	}
	return
}
