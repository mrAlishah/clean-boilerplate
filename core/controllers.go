package core

import (
	adminControllers "boilerplate/apps/admin/controllers"
	genericControllers "boilerplate/apps/generic/controllers"
	"boilerplate/apps/user/controllers"
	"go.uber.org/fx"
)

var ControllerModule = fx.Options(
	fx.Provide(controllers.NewAuthController),
	fx.Provide(genericControllers.NewGenericController),
	fx.Provide(adminControllers.NewUserController),
)
