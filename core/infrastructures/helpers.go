package infrastructures

import (
	gormLogger "gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

//get env for no fx dependency injection environment
func GetEnv() *Env {
	return NewEnv()
}

//get db for no fx dependency injection environment
func GetDBNoFX(env *Env) *GormDB {
	logger := NewLogger(env)
	gLogger := gormLogger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		gormLogger.Config{
			SlowThreshold: time.Second,     // Slow SQL threshold
			LogLevel:      gormLogger.Info, // Log level
			Colorful:      true,            // Disable color
		},
	)
	return GetDB(gLogger, logger, env)
}
