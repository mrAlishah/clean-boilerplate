package middlewares

import (
	repositories2 "boilerplate/apps/user/repositories/gorm"
	services2 "boilerplate/apps/user/services"
	"boilerplate/core/infrastructures"
	"boilerplate/core/interfaces"
	"boilerplate/core/responses"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

//AuthMiddleware -> struct for transaction
type AuthMiddleware struct {
	logger         interfaces.Logger
	authService    *services2.AuthService
	env            *infrastructures.Env
	userRepository *repositories2.UserRepository
}

//NewAuthMiddleware -> new instance of transaction
func NewAuthMiddleware(
	logger *infrastructures.Logger,
	authService *services2.AuthService,
	env *infrastructures.Env,
	userRepository *repositories2.UserRepository,
) *AuthMiddleware {
	return &AuthMiddleware{
		authService:    authService,
		logger:         logger,
		env:            env,
		userRepository: userRepository,
	}
}

type authHeader struct {
	Authorization string `header:"Authorization"`
}

func (m *AuthMiddleware) AuthHandle() gin.HandlerFunc {
	m.logger.Info("setting up auth middleware")

	return func(c *gin.Context) {
		ah := authHeader{}
		if err := c.ShouldBindHeader(&ah); err == nil {
			strs := strings.Split(ah.Authorization, " ")
			bearer := strs[0]
			if bearer != "Bearer" {
				responses.ErrorJSON(c, http.StatusUnauthorized, gin.H{}, "your token dosen't start with 'Bearer '")
				c.Abort()
				return
			}
			accessToken := strs[1]
			valid, claims, _ := m.authService.DecodeToken(accessToken, "access"+m.env.Secret)

			id, ok := claims["userId"].(float64)
			if ok {
				userId := strconv.Itoa(int(id))
				deviceToken := claims["deviceToken"].(string)
				if valid && err == nil {
					c.Set("userId", userId)
					c.Set("deviceToken", deviceToken)
					c.Next()
					return
				}
			}
		}
		responses.ErrorJSON(c, http.StatusUnauthorized, gin.H{}, "You must login to access this page 😥")
		c.Abort()
	}
}
