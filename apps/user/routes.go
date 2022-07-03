package user

import (
	"boilerplate/core/infrastructures"
	"boilerplate/core/interfaces"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserRoutes -> utility routes struct
type UserRoutes struct {
	router *infrastructures.Router
	Logger interfaces.Logger
}

//NewProfileRoute -> returns new utility route
func NewUserRoutes(
	logger *infrastructures.Logger,
	env *infrastructures.Env,
	router *infrastructures.Router,
) UserRoutes {
	return UserRoutes{
		Logger: logger,
		router: router,
	}
}

//Setup -> sets up route for util entities
func (pr UserRoutes) Setup() {
	g := pr.router.Gin.Group("/api/users")
	{
		g.GET("/", func(context *gin.Context) {
			a := 0
			fmt.Println(2 / a)
			context.String(http.StatusOK, "Hello world")
		})
	}
}
