package core

import (
	"boilerplate/apps/user/services"
	"go.uber.org/fx"
)

var ServiceModule = fx.Options(
	fx.Provide(services.FxNewUserService),
	fx.Provide(services.NewAuthService),
)
