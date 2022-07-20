package core

import (
	"boilerplate/apps/admin"
	genericApp "boilerplate/apps/generic"
	"boilerplate/apps/user"
	"go.uber.org/fx"
)

// Module exports dependency to container
var RoutesModule = fx.Options(
	fx.Provide(NewRoutes),
	fx.Provide(user.NewUserRoutes),
	fx.Provide(admin.NewAdminRoutes),
	fx.Provide(genericApp.NewGenericRoutes),
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
	genericRoutes genericApp.GenericRoutes,
	adminRoutes admin.AdminRoutes,
) Routes {
	return Routes{
		userRoutes,
		genericRoutes,
		adminRoutes,
	}
}

// Setup all the route
func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
