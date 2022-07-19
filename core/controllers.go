package core

import (
	genericControllers "boilerplate/apps/generic/controllers"
	"boilerplate/apps/user/controllers"
	"go.uber.org/fx"
)

var ControllerModule = fx.Options(
	fx.Provide(controllers.NewUserController),
	fx.Provide(controllers.NewAuthController),
	fx.Provide(genericControllers.NewGenericController),
)
