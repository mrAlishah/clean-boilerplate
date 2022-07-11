package core

import (
	repositories "boilerplate/apps/user/repositories/gorm"
	"go.uber.org/fx"
)

var RepositoryModule = fx.Options(
	fx.Provide(repositories.NewUserRepository),
)
