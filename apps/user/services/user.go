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
	"github.com/gin-gonic/gin"
)

// UserService -> struct
type UserService struct {
	userRepository interfaces.UserRepository
	db             *infrastructures.GormDB
	logger         interfaces.Logger
	encryption     *infrastructures.Encryption
}

func NewUserService(userRepository interfaces.UserRepository,
	db *infrastructures.GormDB, logger *infrastructures.Logger,
	encryption *infrastructures.Encryption) *UserService {
	return &UserService{
		userRepository: userRepository,
		db:             db,
		logger:         logger,
		encryption:     encryption,
	}
}

// FxNewUserService -> creates a new Userservice
func FxNewUserService(userRepository *repositories.UserRepository,
	db *infrastructures.GormDB, logger *infrastructures.Logger,
	encryption *infrastructures.Encryption) *UserService {
	return NewUserService(userRepository, db, logger, encryption)
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
	}
	return
}

func (s UserService) DeleteUser(id uint64) (err error) {
	err = s.userRepository.DeleteByID(uint(id))
	if !errors.Is(err, errors2.NotFoundError) && err != nil {
		s.logger.Fatal("Failed to find user:%s", err.Error())
	}
	return
}

func (s UserService) UpdateUser(userData DTO.UpdateUserRequestAdmin, userID uint64) (user models.User, err error) {
	userData.ToModel(&user)
	err = s.userRepository.UpdateModel(&user, userID)
	if !errors.Is(err, errors2.NotFoundError) && err != nil {
		s.logger.Fatal("Failed to find user:%s", err.Error())
		return
	}
	if err != nil {
		s.logger.Fatal("Failed to update user:%s", err.Error())
		return
	}
	return
}

func (s UserService) DetailUser(id uint64) (user models.User, err error) {
	user, err = s.userRepository.FindByField("id", uint(id))
	if !errors.Is(err, errors2.NotFoundError) && err != nil {
		s.logger.Fatal("Failed to find user:%s", err.Error())
	}
	return
}

//get authenticated user
func (s UserService) GetAuthenticatedUser(c *gin.Context) (models.User, error) {
	userId := c.MustGet("userId").(string)
	if userId == "" {
		return models.User{}, errors.New("user didn't logged in")
	}
	return s.userRepository.FindByField("id", userId)
}
