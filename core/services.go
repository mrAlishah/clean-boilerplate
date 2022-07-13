package core

import (
	"boilerplate/apps/user/services"
	"go.uber.org/fx"
)

var ServiceModule = fx.Options(
	fx.Provide(services.NewUserService),
	fx.Provide(services.NewAuthService),
)
