package infrastructures

import (
	"boilerplate/core/interfaces"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

// GormDB modal
type GormDB struct {
	DB *gorm.DB
}

// NewGormDB creates a new database instance
func NewGormDB(logger *Logger, env *Env) *GormDB {
	gLogger := gormLogger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		gormLogger.Config{
			SlowThreshold: time.Second,     // Slow SQL threshold
			LogLevel:      gormLogger.Info, // Log level
			Colorful:      true,            // Disable color
		},
	)
	if env.Environment == "test" {
		RemoveDB(logger, env, "")
		CreateDB(logger, env, "")
	}
	return GetDB(gLogger, logger, env)
}

// connect to database and return GormDB object
func GetDB(gormLogger gormLogger.Interface, logger interfaces.Logger, env *Env) *GormDB {
	url := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=Europe/London",
		env.DBHost, env.DBUsername, env.DBPassword, env.DBName,
		env.DBPort)

	if env.Environment != "development" {
		url = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/London",
			env.DBHost, env.DBUsername, env.DBPassword, env.DBName,
			env.DBPort)
	}

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{Logger: gormLogger})
	if err != nil {
		logger.Info("Url: ", url)
		logger.Fatal("Error get db:", err.Error())
	}

	logger.Info("GormDB connection established ✔️")

	return &GormDB{
		DB: db,
	}
}

//create database if no DBName passed , it automaticaly use dbname of environment varible
func CreateDB(logger interfaces.Logger, env *Env, DBName string) {
	db, err := ConnectDBSQL(env)
	if err != nil {
		logger.Fatal("Error create db:", err.Error())
	}
	defer db.Close()

	if DBName == "" {
		DBName = env.DBName
	}

	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s;", DBName))
	if err != nil {
		logger.Fatal("Error create db:", err)
	}
}

//remove passed database if no DBName passed , it automaticaly use dbname of environment varible
func RemoveDB(logger interfaces.Logger, env *Env, DBName string) {
	db, err := ConnectDBSQL(env)
	if err != nil {
		logger.Fatal("Error remove db:", err.Error())
	}
	defer db.Close()

	if DBName == "" {
		DBName = env.DBName
	}

	_, err = db.Exec(fmt.Sprintf("DROP DATABASE %s;", DBName))
	if err != nil {
		logger.Fatal("Error remove db:", err.Error())
	}
}

//connect to database without passing database name and with database/sql package
//useful for doing general database sql statements that not related to a specefic database
//be sure to defer db.Close() in using function
func ConnectDBSQL(env *Env) (*sql.DB, error) {
	url := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable",
		env.DBHost, env.DBPort, env.DBUsername, env.DBPassword)
	return sql.Open("postgres", url)
}
