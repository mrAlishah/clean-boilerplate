package infrastructures

import (
	"os"
)

// Env has environment stored
type Env struct {
	ServerPort  string
	Environment string
	AppName     string
	BasePath    string
	SiteUrl     string
	DBUsername  string
	DBPassword  string
	DBHost      string
	DBPort      string
	DBName      string
	Secret      string
}

// NewEnv creates a new environment
func NewEnv() *Env {
	env := Env{}
	env.LoadEnv()
	return &env
}

// LoadEnv loads environment
func (env *Env) LoadEnv() {
	env.ServerPort = os.Getenv("ServerPort")
	env.Environment = os.Getenv("Environment")
	env.AppName = os.Getenv("AppName")
	env.BasePath = os.Getenv("BasePath")
	env.SiteUrl = os.Getenv("SiteUrl")

	env.DBUsername = os.Getenv("DBUsername")
	env.DBPassword = os.Getenv("DBPassword")
	env.DBHost = os.Getenv("DBHost")
	env.DBPort = os.Getenv("DBPort")
	env.DBName = os.Getenv("DBName")

	env.Secret = os.Getenv("Secret")
}
