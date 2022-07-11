package gorm

import (
	"boilerplate/core/infrastructures"
	"github.com/go-playground/validator/v10"
)

type FkValidator struct {
	db *infrastructures.GormDB
}

func NewFkValidator(database *infrastructures.GormDB) FkValidator {
	return FkValidator{
		db: database,
	}
}

//fk validator
//please send destionation table name as foreign key
func (uv FkValidator) Handler() func(fl validator.FieldLevel) bool {
	return func(fl validator.FieldLevel) bool {
		value := fl.Field()
		table := fl.Param()
		var count int64
		uv.db.DB.Table(table).Where("id=?", value.Uint()).Count(&count)
		return count > 0
	}
}
