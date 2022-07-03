package core

import (
	"boilerplate/apps/user"
	"go.uber.org/fx"
)

// Module exports dependency to container
var RoutesModule = fx.Options(
	fx.Provide(NewRoutes),
	fx.Provide(user.NewUserRoutes),
)

// Routes contains multiple routes
type Routes []Route

// Route interface
type Route interface {
	Setup()
}

// NewRoutes sets up routes
func NewRoutes(
	userRoutes user.UserRoutes,
) Routes {
	return Routes{
		userRoutes,
	}
}

// Setup all the route
func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
