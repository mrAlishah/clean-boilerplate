package admin

import (
	"boilerplate/apps/admin/controllers"
	"boilerplate/core/infrastructures"
	"boilerplate/core/interfaces"
)

// AdminRoutes -> utility routes struct
type AdminRoutes struct {
	router         *infrastructures.Router
	logger         interfaces.Logger
	userController *controllers.UserController
}

//NewProfileRoute -> returns new utility route
func NewAdminRoutes(
	logger *infrastructures.Logger,
	env *infrastructures.Env,
	router *infrastructures.Router,
	userController *controllers.UserController,
) AdminRoutes {
	return AdminRoutes{
		logger:         logger,
		router:         router,
		userController: userController,
	}
}

//Setup -> sets up route for util entities
func (pr AdminRoutes) Setup() {
	g := pr.router.Gin.Group("/api/admin/users")
	{
		g.GET("/", pr.userController.ListUser)
		g.POST("/", pr.userController.CreateUser)
		g.DELETE("/:id", pr.userController.DeleteUser)
		g.PUT("/:id", pr.userController.UpdateUser)
		g.GET("/:id", pr.userController.DetailUser)
	}

}
