package services

import (
	"boilerplate/apps/user/DTO"
	repositories "boilerplate/apps/user/repositories/gorm"
	errors2 "boilerplate/core/errors"
	"boilerplate/core/infrastructures"
	"boilerplate/core/interfaces"
	"boilerplate/core/models"
	"boilerplate/core/utils"
	"errors"
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
func (s UserService) GetAllUsers(pagination utils.Pagination) (users []models.User, count int64, err error) {
	users, count, err = s.userRepository.GetAllUsers(pagination)
	if err != nil {
		s.logger.Fatal("Failed to get users", err.Error())
		return
	}
	return
}

func (s UserService) CreateUser(userData DTO.CreateUserRequestAdmin) (err error) {
	encryptedPassword := s.encryption.SaltAndSha256Encrypt(userData.Password, userData.Email)
	user := models.User{
		Password:  encryptedPassword,
		FirstName: userData.FirstName,
		LastName:  userData.LastName,
		Email:     userData.Email,
	}
	err = s.userRepository.Create(&user)
	if err != nil {
		s.logger.Fatal("Failed to create user:%s", err.Error())
		return
	}
	return
}

func (s UserService) DeleteUser(id uint64) (err error) {
	err = s.userRepository.DeleteUser(id)
	if !errors.Is(err, errors2.NotFoundError) && err != nil {
		s.logger.Fatal("Failed to find user:%s", err.Error())
		return err
	}
	return err
}
