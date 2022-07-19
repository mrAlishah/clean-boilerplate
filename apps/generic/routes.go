package generic

import (
	"boilerplate/apps/generic/controllers"
	"boilerplate/core/infrastructures"
	"boilerplate/core/interfaces"
)

// GenericRoutes -> utility routes struct
type GenericRoutes struct {
	Router            *infrastructures.Router
	Logger            interfaces.Logger
	Env               *infrastructures.Env
	GenericController *controllers.GenericController
}

//NewGenericRoute -> returns new utility route
func NewGenericRoutes(
	logger *infrastructures.Logger,
	router *infrastructures.Router,
	env *infrastructures.Env,
	genericController *controllers.GenericController,
) GenericRoutes {
	return GenericRoutes{
		Logger:            logger,
		Router:            router,
		GenericController: genericController,
		Env:               env,
	}
}

//Setup -> sets up route for util entities
func (gr GenericRoutes) Setup() {
	gr.Router.Gin.GET("/api/ping", gr.GenericController.Ping)
	gr.Router.Gin.Static("/media", gr.Env.BasePath+"/media")
}
