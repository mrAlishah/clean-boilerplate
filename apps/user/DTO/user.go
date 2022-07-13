package DTO

import (
	"boilerplate/core/infrastructures"
	"boilerplate/core/models"
)

func UsersToResponses(users []models.User) []UserResponse {
	userResponses := make([]UserResponse, 0)
	for _, v := range users {
		var userResponse UserResponse
		userResponse.FromModel(v)
		userResponses = append(userResponses, userResponse)
	}
	return userResponses
}

type RegisterRequest struct {
	Email          string `json:"email" binding:"required,uniqueGorm=users&email,email"`
	FirstName      string `json:"firstName" binding:"required"`
	LastName       string `json:"lastName" binding:"required"`
	Password       string `json:"password" binding:"required"`
	RepeatPassword string `json:"repeatPassword" binding:"required,eqfield=Password"`
}

type UserResponse struct {
	models.BaseResponse
	Email     string `json:"email"`
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Password  string `json:"-"`
}

func (r *UserResponse) FromModel(userModel models.User) {
	r.Email = userModel.Email
	r.FirstName = userModel.FirstName
	r.LastName = userModel.LastName
	r.Password = userModel.Password
}

func (r *RegisterRequest) ToModel(encryption infrastructures.Encryption, m models.User) {
	r.Email = m.Email
	r.FirstName = m.FirstName
	r.LastName = m.LastName
	r.Password = encryption.SaltAndSha256Encrypt(m.Password, m.Email)
}

type CreateUserRequestAdmin struct {
	Email          string `json:"email" binding:"required,uniqueGorm=users&email,email"`
	FirstName      string `json:"firstName" binding:"required"`
	LastName       string `json:"lastName" binding:"required"`
	Password       string `json:"password" binding:"required"`
	RepeatPassword string `json:"repeatPassword" binding:"required,eqfield=Password"`
}
