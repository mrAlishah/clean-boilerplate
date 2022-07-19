package utils

import (
	"boilerplate/core/infrastructures"
	gormLogger "gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

//get env for no fx dependency injection environment
func GetEnv() *infrastructures.Env {
	return infrastructures.NewEnv()
}

//get db for no fx dependency injection environment
func GetDB(env *infrastructures.Env) *infrastructures.GormDB {
	logger := infrastructures.NewLogger(env)
	gLogger := gormLogger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		gormLogger.Config{
			SlowThreshold: time.Second,     // Slow SQL threshold
			LogLevel:      gormLogger.Info, // Log level
			Colorful:      true,            // Disable color
		},
	)
	return infrastructures.GetDB(gLogger, logger, env)
}
