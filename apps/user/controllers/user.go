package controllers

import (
	"boilerplate/apps/user/DTO"
	"boilerplate/apps/user/services"
	errors2 "boilerplate/core/errors"
	"boilerplate/core/infrastructures"
	"boilerplate/core/interfaces"
	"boilerplate/core/responses"
	"boilerplate/core/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserController struct {
	logger      interfaces.Logger
	env         *infrastructures.Env
	userService *services.UserService
}

func NewUserController(logger *infrastructures.Logger,
	env *infrastructures.Env,
	userService *services.UserService,
) *UserController {
	return &UserController{
		logger:      logger,
		env:         env,
		userService: userService,
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
	uc.paginateUserList(c, "")
}

// @Summary create users
// @Schemes
// @Description create user and admin , admin only
// @Tags admin
// @Accept json
// @Produce json
// @Param email formData string true "unique email"
// @Param password formData string true "password that have at least 8 length and contain an alphabet and number "
// @Param repeatPassword formData string true "repeatPassword that have at least 8 length and contain an alphabet and number "
// @Param firstName formData string true "firstName"
// @Param lastName formData string true "lastName"
// @Param isAdmin formData bool true "isAdmin"
// @Success 200 {object} swagger.UsersListResponse
// @failure 401 {object} swagger.UnauthenticatedResponse
// @failure 403 {object} swagger.AccessForbiddenResponse
// @Router /admin/users [post]
func (uc UserController) CreateUser(c *gin.Context) {
	var userData DTO.CreateUserRequestAdmin
	if err := c.ShouldBindJSON(&userData); err != nil {
		fieldErrors := make(map[string]string, 0)
		if !utils.IsGoodPassword(userData.Password) {
			fieldErrors["password"] = "Password must contain at least one alphabet and one number and its length must be 8 characters or more"

		}
		responses.ValidationErrorsJSON(c, err, "", fieldErrors)
		return
	}
	if !utils.IsGoodPassword(userData.Password) {
		fieldErrors := map[string]string{
			"password": "Password must contain at least one alphabet and one number and its length must be 8 characters or more",
		}
		responses.ManualValidationErrorsJSON(c, fieldErrors, "")
		return
	}

	err := uc.userService.CreateUser(userData)
	if err != nil {
		responses.ErrorJSON(c, http.StatusInternalServerError, gin.H{}, "Sorry an error occurred!")
	}
	uc.paginateUserList(c, "User created successfully.")
}

// @Summary delete user
// @Schemes
// @Description delete user or admin , admin only
// @Tags admin
// @Accept json
// @Produce json
// @Param id path int true "user id"
// @Success 200 {object} swagger.UsersListResponse
// @failure 401 {object} swagger.UnauthenticatedResponse
// @failure 404 {object} swagger.NotFoundResponse
// @failure 403 {object} swagger.AccessForbiddenResponse
// @Router /admin/users/{id} [delete]
func (uc UserController) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responses.ErrorJSON(c, http.StatusNotFound, gin.H{}, "No user found")
		return
	}

	err = uc.userService.DeleteUser(uint64(id))
	if errors.Is(err, errors2.NotFoundError) {
		responses.ErrorJSON(c, http.StatusNotFound, gin.H{}, "No user found")
		return
	}
	if err != nil {
		responses.ErrorJSON(c, http.StatusInternalServerError, gin.H{}, "An error occurred")
		return
	}

	uc.paginateUserList(c, "User deleted successfully !")
}

func (uc UserController) UpdateUser(c *gin.Context) {
	var userData DTO.UpdateUserRequestAdmin
	if err := c.ShouldBindJSON(&userData); err != nil {
		responses.ValidationErrorsJSON(c, err, "", map[string]string{})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responses.ErrorJSON(c, http.StatusNotFound, gin.H{}, "No user found")
		return
	}
	userData.ID = uint64(id)

	_, err = uc.userService.UpdateUser(userData, uint64(id))
	if errors.Is(err, errors2.NotFoundError) {
		responses.ErrorJSON(c, http.StatusNotFound, gin.H{}, "No user found")
		return
	}
	if err != nil {
		responses.ErrorJSON(c, http.StatusInternalServerError, gin.H{}, "Sorry an error occurred!")
	}
	uc.detailUser(c, uint64(id))
}

func (uc *UserController) paginateUserList(c *gin.Context, message string) {
	pagination := utils.BuildPagination(c)
	users, count, err := uc.userService.GetAllUsers(pagination)
	if err != nil {
		responses.ErrorJSON(c, http.StatusInternalServerError, gin.H{}, "Sorry an error occurred 😢")
		return
	}
	responses.JSON(c, http.StatusOK, gin.H{
		"users": users,
		"count": count,
	}, message)
}

func (uc *UserController) detailUser(c *gin.Context, id uint64) {
	user, err := uc.userService.DetailUser(id)
	if errors.Is(err, errors2.NotFoundError) {
		responses.ErrorJSON(c, http.StatusNotFound, gin.H{}, "No user found")
		return
	}
	if err != nil {
		responses.ErrorJSON(c, http.StatusInternalServerError, gin.H{}, "Sorry an error occurred!")
		return
	}
	var userResponse DTO.UserResponse
	userResponse.FromModel(user)
	responses.JSON(c, http.StatusOK, userResponse, "")
}
