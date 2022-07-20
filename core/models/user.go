package models

type User struct {
	Base
	Email     string `json:"email" binding:"required"`
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Password  string `json:"password" binding:"required"`
	IsAdmin   bool   `json:"isAdmin" binding:"required"`
}

// TableName gives table name of model
func (m User) TableName() string {
	return "users"
}
