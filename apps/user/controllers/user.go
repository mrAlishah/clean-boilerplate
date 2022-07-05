package controllers

import (
	"boilerplate/apps/user/services"
	"boilerplate/core/infrastructures"
	"boilerplate/core/interfaces"
	"boilerplate/core/responses"
	"boilerplate/core/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	logger      interfaces.Logger
	env         infrastructures.Env
	userService services.UserService
}

func NewUserController(logger infrastructures.Logger,
	env infrastructures.Env,
	userService services.UserService,
) UserController {
	return UserController{
		logger: &logger,
		env:    env,
	}
}

// @Summary get users list
// @Schemes
// @Description list of paginated response , authentication required
// @Tags admin
// @Accept json
// @Produce json
// @Success 200 {object} swagger.UsersListResponse
// @failure 401 {object} swagger.UnauthenticatedResponse
// @failure 403 {object} swagger.AccessForbiddenResponse
// @Router /users [get]
func (uc UserController) ListUser(c *gin.Context) {
	pagination := utils.BuildPagination(c)
	users, count, err := uc.userService.GetAllUsers(pagination)
	if err != nil {
		responses.ErrorJSON(c, http.StatusInternalServerError, gin.H{}, "Sorry an error occurred 😢")
		return
	}
	responses.JSON(c, http.StatusOK, gin.H{
		"users": users,
		"count": count,
	}, "")
}
