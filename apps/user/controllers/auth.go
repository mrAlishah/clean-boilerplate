package controllers

import (
	"boilerplate/apps/user/DTO"
	repositories "boilerplate/apps/user/repositories/gorm"
	"boilerplate/apps/user/services"
	"boilerplate/core/infrastructures"
	"boilerplate/core/interfaces"
	"boilerplate/core/models"
	"boilerplate/core/responses"
	"boilerplate/core/utils"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type AuthController struct {
	logger         interfaces.Logger
	env            *infrastructures.Env
	encryption     *infrastructures.Encryption
	userService    *services.UserService
	authService    *services.AuthService
	userRepository *repositories.UserRepository
}

func NewAuthController(
	env *infrastructures.Env,
	encryption *infrastructures.Encryption,
	userService *services.UserService,
	authService *services.AuthService,
	userRepository *repositories.UserRepository,
) *AuthController {
	return &AuthController{
		env:            env,
		encryption:     encryption,
		userService:    userService,
		authService:    authService,
		userRepository: userRepository,
	}
}

// @BasePath /api/auth

// @Summary register
// @Schemes
// @Description jwt register
// @Tags auth
// @Accept json
// @Produce json
// @Param email query string true "unique email"
// @Param password query string true "password that have at least 8 length and contain an alphabet and number "
// @Param repeatPassword query string true "repeatPassword that have at least 8 length and contain an alphabet and number "
// @Param firstName query string true "firstName"
// @Param lastName query string true "lastName"
// @Success 200 {object} swagger.SuccessResponse
// @failure 422 {object} swagger.FailedValidationResponse
// @Router /auth/register [post]
func (ac AuthController) Register(c *gin.Context) {

	// Data Parse
	var userData DTO.RegisterRequest
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
	var user models.User
	encryptedPassword := ac.encryption.SaltAndSha256Encrypt(userData.Password, userData.Email)
	user.Password = encryptedPassword

	user.FirstName = userData.FirstName
	user.LastName = userData.LastName
	user.Email = userData.Email
	err := ac.userRepository.Create(&user)
	if err != nil {
		ac.logger.Fatal("Failed to create registered user ", err.Error())
		responses.ErrorJSON(c, http.StatusInternalServerError, gin.H{}, "Sorry an error occurred in registering your account!")
		return
	}
	responses.JSON(c, http.StatusOK, gin.H{}, "Your account created successfully, an verification link sent to your email use that to verify your account")
}

// @Summary login
// @Schemes
// @Description jwt login
// @Tags auth
// @Accept json
// @Produce json
// @Param email query string true "email"
// @Param deviceName query string true "send user operating system + browser name in this param"
// @Param password query string true "password"
// @Success 200 {object} swagger.LoginResponse
// @failure 422 {object} swagger.FailedValidationResponse
// @failure 401 {object} swagger.FailedLoginResponse
// @Router /auth/login [post]
func (ac AuthController) Login(c *gin.Context) {
	// Data Parse
	var loginRequest DTO.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		responses.ValidationErrorsJSON(c, err, "", map[string]string{})
		return
	}
	var user models.User
	user, err := ac.userRepository.FindByField("Email", loginRequest.Email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		responses.ErrorJSON(c, http.StatusUnauthorized, gin.H{}, "No user found with entered credentials")
		return
	}
	if err != nil {
		ac.logger.Fatal("Error to find user:%s", err.Error())
		responses.ErrorJSON(c, http.StatusInternalServerError, gin.H{}, "An error occured")
		return
	}
	encryptedPassword := ac.encryption.SaltAndSha256Encrypt(loginRequest.Password, loginRequest.Email)
	if user.Password == encryptedPassword {
		tokensData, err := ac.authService.CreateTokens(user)
		if err != nil {
			ac.logger.Fatal("Failed generate jwt tokens:%s", err.Error())
			responses.ErrorJSON(c, http.StatusInternalServerError, gin.H{}, "An error occured")
			return
		}
		var loginResult DTO.LoginResponse
		loginResult.AccessToken = tokensData["accessToken"]
		loginResult.RefreshToken = tokensData["refreshToken"]
		loginResult.ExpRefreshToken = tokensData["expRefreshToken"]
		loginResult.ExpAccessToken = tokensData["expAccessToken"]
		var userResponse DTO.UserResponse
		userResponse.FromModel(user)
		loginResult.User = userResponse

		responses.JSON(c, http.StatusOK, loginResult, "Hello "+user.FirstName+" wellcome back")
		return
	} else {
		responses.ErrorJSON(c, http.StatusUnauthorized, gin.H{}, "No user found with entered credentials")
		return
	}
}

// @Summary access token verify
// @Schemes
// @Description jwt access token verify
// @Tags auth
// @Accept json
// @Produce json
// @Param accessToken query string true "accessToken"
// @Success 200 {object} swagger.SuccessResponse
// @failure 422 {object} swagger.FailedValidationResponse
// @failure 401 {object} swagger.FailedResponse
// @Router /auth/access-token-verify [post]
func (ac AuthController) AccessTokenVerify(c *gin.Context) {
	at := DTO.AccessTokenReq{}
	if err := c.ShouldBindJSON(&at); err != nil {
		responses.ValidationErrorsJSON(c, err, "", map[string]string{})
		return
	}

	accessToken := at.AccessToken
	accessSecret := "access" + ac.env.Secret
	valid, _, err := ac.authService.DecodeToken(accessToken, accessSecret)
	if err != nil {
		responses.ErrorJSON(c, http.StatusUnauthorized, gin.H{}, "Access token is not valid")
		return
	}

	if valid {
		responses.JSON(c, http.StatusOK, gin.H{}, "Access token is valid")
		return
	} else {
		responses.ErrorJSON(c, http.StatusUnauthorized, gin.H{}, "Access token is not valid")
		return
	}
}

// @Summary renew access token
// @Schemes
// @Description jwt renew access token
// @Tags auth
// @Accept json
// @Produce json
// @Param refreshToken query string true "refreshToken"
// @Success 200 {object} swagger.SuccessVerifyAccessTokenResponse
// @failure 422 {object} swagger.FailedValidationResponse
// @failure 400 {object} swagger.FailedResponse
// @Router /auth/renew-access-token [post]
func (ac AuthController) RenewToken(c *gin.Context) {
	rtr := DTO.RefreshTokenRequest{}
	if err := c.ShouldBindJSON(&rtr); err != nil {
		responses.ValidationErrorsJSON(c, err, "", map[string]string{})
		return
	}

	//Parse and extract claims
	refreshToken := rtr.RefreshToken
	var valid bool
	var atClaims jwt.MapClaims
	refreshSecret := "refresh" + ac.env.Secret
	valid, atClaims, err := ac.authService.DecodeToken(refreshToken, refreshSecret)
	if err != nil {
		responses.ErrorJSON(c, http.StatusBadRequest, gin.H{}, "Refresh token is not valid")
		return
	}

	uid, ok := atClaims["userId"].(float64)
	if !ok {
		responses.ErrorJSON(c, http.StatusBadRequest, gin.H{}, "Refresh token is not valid")
		return
	}
	userID := int(uid)

	user, err := ac.userRepository.FindByField("id", strconv.Itoa(userID))
	//don't allow deleted user renew access token
	if errors.Is(err, gorm.ErrRecordNotFound) {
		responses.ErrorJSON(c, http.StatusBadRequest, gin.H{}, "Refresh token is not valid")
		return
	}
	if err != nil {
		ac.logger.Fatal("Error in finding user:", err)
		responses.ErrorJSON(c, http.StatusBadRequest, gin.H{}, "Refresh token is not valid")
		return
	}

	if valid {
		var exp int64
		accessSecret := "access" + ac.env.Secret
		exp = time.Now().Add(time.Hour * 2).Unix()
		accessToken, _ := ac.authService.CreateAccessToken(user, exp, accessSecret)
		responses.JSON(c, http.StatusOK, DTO.AccessTokenRes{AccessToken: accessToken, ExpAccessToken: strconv.Itoa(int(exp))}, "")
		return
	} else {
		responses.ErrorJSON(c, http.StatusBadRequest, gin.H{}, "Refresh token is not valid")
		return
	}
}

// @Summary logout
// @Schemes
// @Description jwt logout , atuhentication required
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} swagger.SuccessResponse
// @failure 401 {object} swagger.UnauthenticatedResponse
// @Router /auth/logout [post]
//func (ac AuthController) Logout(c *gin.Context) {
//	user, err := ac.userRepository.GetAuthenticatedUser(c)
//	if err != nil {
//		ac.logger.Zap.Error("Failed to change password", err.Error())
//		responses.ErrorJSON(c, http.StatusInternalServerError, gin.H{}, "Sorry an error occoured in changing password!")
//		return
//	}
//	deviceToken := c.MustGet("deviceToken").(string)
//	ac.authService.DeleteDevice(&user, deviceToken)
//
//	responses.JSON(c, http.StatusOK, gin.H{}, "You logged out successfuly")
//}
