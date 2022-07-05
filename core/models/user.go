package models

type User struct {
	Base
	Email     string `json:"email" binding:"required"`
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

// TableName gives table name of model
func (m User) TableName() string {
	return "users"
}

func (u User) ToResponse() UserResponse {
	return UserResponse{
		BaseResponse: BaseResponse{
			CreatedAt: u.CreatedAt.Unix(),
			UpdatedAt: u.UpdatedAt.Unix(),
			ID:        u.ID,
		},
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Password:  u.Password,
	}
}

func UsersToResponses(users []User) []UserResponse {
	userResponses := make([]UserResponse, len(users))
	for i, v := range users {
		userResponses[i] = v.ToResponse()
	}
	return userResponses
}

type UserResponse struct {
	BaseResponse
	Email     string `json:"email"`
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Password  string `json:"-"`
}

type CreateUserRequestAdmin struct {
	Email          string `json:"email" binding:"required,uniqueGorm=users&email,email"`
	FirstName      string `json:"firstName" binding:"required"`
	LastName       string `json:"lastName" binding:"required"`
	Password       string `json:"password" binding:"required"`
	RepeatPassword string `json:"repeatPassword" binding:"required,eqfield=Password"`
}
