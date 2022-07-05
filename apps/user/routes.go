package user

import (
	"boilerplate/apps/user/controllers"
	"boilerplate/core/infrastructures"
	"boilerplate/core/interfaces"
)

// UserRoutes -> utility routes struct
type UserRoutes struct {
	router         *infrastructures.Router
	logger         interfaces.Logger
	userController *controllers.UserController
}

//NewProfileRoute -> returns new utility route
func NewUserRoutes(
	logger *infrastructures.Logger,
	env *infrastructures.Env,
	router *infrastructures.Router,
	userController *controllers.UserController,
) UserRoutes {
	return UserRoutes{
		logger:         logger,
		router:         router,
		userController: userController,
	}
}

//Setup -> sets up route for util entities
func (pr UserRoutes) Setup() {
	g := pr.router.Gin.Group("/api/users")
	{
		g.GET("/", pr.userController.ListUser)
		g.POST("/", pr.userController.CreateUser)
		g.DELETE("/:id", pr.userController.DeleteUser)
	}
}
