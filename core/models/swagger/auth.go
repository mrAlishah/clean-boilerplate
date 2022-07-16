package swagger

import (
	"boilerplate/apps/user/DTO"
)

type LoginResponse struct {
	SuccessResponse
	Data DTO.LoginResponse `json:"data"`
}

type FailedLoginResponse struct {
	FailedResponse
	Msg string `json:"msg" example:"No user found with entered credentials"`
}

type SuccessVerifyAccessTokenResponse struct {
	SuccessResponse
	Data DTO.AccessTokenRes `json:"data"`
}

type UnauthenticatedResponse struct {
	FailedResponse
	Msg string `json:"msg" example:"You must login first!"`
}
