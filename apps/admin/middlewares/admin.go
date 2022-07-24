package middlewares

import (
	"boilerplate/apps/user/services"
	"boilerplate/core/infrastructures"
	"boilerplate/core/interfaces"
	"boilerplate/core/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

//AdminMiddleware -> struct for transaction
type AdminMiddleware struct {
	logger      interfaces.Logger
	authService *services.AuthService
	env         *infrastructures.Env
	userService *services.UserService
}

//NewAdminMiddleware -> new instance of transaction
func NewAdminMiddleware(
	logger *infrastructures.Logger,
	authService *services.AuthService,
	env *infrastructures.Env,
	userService *services.UserService,
) *AdminMiddleware {
	return &AdminMiddleware{
		authService: authService,
		logger:      logger,
		env:         env,
		userService: userService,
	}
}

func (m *AdminMiddleware) AdminHandle() gin.HandlerFunc {
	m.logger.Info("setting up admin middleware")

	return func(c *gin.Context) {
		user, err := m.userService.GetAuthenticatedUser(c)
		if err != nil {
			m.logger.Fatal("Failed to get user in admin middleware", err.Error())
			responses.ErrorJSON(c, http.StatusInternalServerError, gin.H{}, "Sorry an error occoured 😢")
			c.Abort()
			return
		}
		if !user.IsAdmin {
			responses.ErrorJSON(c, http.StatusForbidden, gin.H{}, "You don't have access to this page 😥")
			c.Abort()
			return

		}
		c.Next()
	}
}
