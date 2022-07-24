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
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	IsAdmin   bool   `json:"isAdmin"`
}

func (r *UserResponse) FromModel(userModel models.User) {
	r.ID = userModel.ID
	r.CreatedAt = userModel.CreatedAt.Unix()
	r.UpdatedAt = userModel.UpdatedAt.Unix()
	r.Email = userModel.Email
	r.FirstName = userModel.FirstName
	r.LastName = userModel.LastName
	r.IsAdmin = userModel.IsAdmin
}

func (r *RegisterRequest) ToModel(encryption infrastructures.Encryption, m *models.User) {
	m.Email = r.Email
	m.FirstName = r.FirstName
	m.LastName = r.LastName
	m.Password = encryption.SaltAndSha256Encrypt(r.Password, r.Email)
}

type CreateUserRequestAdmin struct {
	Email          string `json:"email" binding:"required,uniqueGorm=users&email,email"`
	FirstName      string `json:"firstName" binding:"required"`
	LastName       string `json:"lastName" binding:"required"`
	Password       string `json:"password" binding:"required"`
	RepeatPassword string `json:"repeatPassword" binding:"required,eqfield=Password"`
	IsAdmin        bool   `json:"isAdmin" binding:"required"`
}

type UpdateUserRequestAdmin struct {
	Email     string `json:"email" binding:"required,uniqueGorm=users&email,email"`
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	IsAdmin   bool   `json:"isAdmin" binding:"required"`
	ID        uint64 `json:"id"`
}

func (r *UpdateUserRequestAdmin) ToModel(m *models.User) {
	m.Email = r.Email
	m.FirstName = r.FirstName
	m.LastName = r.LastName
	m.IsAdmin = r.IsAdmin
}
