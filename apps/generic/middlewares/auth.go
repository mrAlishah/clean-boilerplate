package middlewares

import (
	"cotchciti/apps/authApp/services"
	"cotchciti/apps/userApp/repositories"
	"cotchciti/core/infrastructure"
	"cotchciti/core/responses"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

//AuthMiddleware -> struct for transaction
type AuthMiddleware struct {
	logger         infrastructure.Logger
	authService    services.AuthService
	env            infrastructure.Env
	userRepository repositories.UserRepository
}

//NewAuthMiddleware -> new instance of transaction
func NewAuthMiddleware(
	logger infrastructure.Logger,
	authService services.AuthService,
	env infrastructure.Env,
	userRepository repositories.UserRepository,
) AuthMiddleware {
	return AuthMiddleware{
		authService:    authService,
		logger:         logger,
		env:            env,
		userRepository: userRepository,
	}
}

type authHeader struct {
	Authorization string `header:"Authorization"`
}

func (m AuthMiddleware) AuthHandle() gin.HandlerFunc {
	m.logger.Zap.Info("setting up auth middleware")

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
			valid, claims, _ := services.DecodeToken(accessToken, "access"+m.env.Secret)

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
