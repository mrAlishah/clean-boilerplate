package core

import (
	"boilerplate/apps/user/controllers"
	"go.uber.org/fx"
)

var ControllerModule = fx.Options(
	fx.Provide(controllers.NewUserController),
)
