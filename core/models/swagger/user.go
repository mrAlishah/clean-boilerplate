package swagger

import (
	"boilerplate/apps/user/DTO"
)

type PaginateUsersData struct {
	Count int                `json:"count" example:"10"`
	List  []DTO.UserResponse `json:"list"`
}

type UsersListResponse struct {
	SuccessResponse
	Data PaginateUsersData
}

type SingleUserResponse struct {
	SuccessResponse
	Data DTO.UserResponse
}
