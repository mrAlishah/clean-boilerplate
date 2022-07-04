package models

type User struct {
	Base
	IsAdmin   bool   `json:"-"`
	Email     string `json:"email" binding:"required"`
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

// TableName gives table name of model
func (m User) TableName() string {
	return "users"
}

// ToMap convert User to map
func (m User) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"Email":     m.Email,
		"FirstName": m.FirstName,
		"LastName":  m.LastName,
		"IsAdmin":   m.IsAdmin,
	}
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
	Email          string `json:"email" binding:"required,uniqueDB=users&email,email"`
	FirstName      string `json:"firstName" binding:"required"`
	LastName       string `json:"lastName" binding:"required"`
	Password       string `json:"password" binding:"required"`
	RepeatPassword string `json:"repeatPassword" binding:"required,eqfield=Password"`
}
