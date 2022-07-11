package validators

import (
	"boilerplate/core/responses/validators/gorm"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.uber.org/fx"
)

// Module Middleware exported
var Module = fx.Options(
	fx.Provide(gorm.NewUniqueValidator),
	fx.Provide(gorm.NewFkValidator),
	fx.Provide(NewTimestampValidator),
	fx.Provide(NewValidators),
)

type Validators struct {
	uv  gorm.UniqueValidator
	fkv gorm.FkValidator
	ts  TimestampValidator
}

func NewValidators(uv gorm.UniqueValidator, fkv gorm.FkValidator, ts TimestampValidator) Validators {
	return Validators{
		uv:  uv,
		fkv: fkv,
		ts:  ts,
	}
}

// Setup sets up middlewares
func (val Validators) Setup() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("uniqueGorm", val.uv.Handler())
		v.RegisterValidation("fkGorm", val.fkv.Handler())
		v.RegisterValidation("timestamp", val.ts.Handler())
	}
}
