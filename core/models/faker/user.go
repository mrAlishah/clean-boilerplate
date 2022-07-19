package faker

import (
	"boilerplate/core/infrastructures"
	"boilerplate/core/models"
	"boilerplate/core/utils"
	"github.com/brianvoe/gofakeit/v6"
)

type User struct{}

func (u *User) CreateOne() models.User {
	env := infrastructures.NewEnv()
	logger := infrastructures.NewLogger(env)
	email := utils.GenerateRandomCode(10) + "@gmail.com"
	encryption := infrastructures.NewEncryption(logger, env)
	return models.User{
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Email:     email,
		Password:  encryption.SaltAndSha256Encrypt("m1234567", email),
	}
}

func (u *User) CreateMany(count int) (users []models.User) {
	env := infrastructures.NewEnv()
	logger := infrastructures.NewLogger(env)
	email := utils.GenerateRandomCode(10) + "@gmail.com"
	encryption := infrastructures.NewEncryption(logger, env)
	for i := 0; i < count; i++ {
		users = append(users, models.User{
			FirstName: gofakeit.FirstName(),
			LastName:  gofakeit.LastName(),
			Email:     email,
			Password:  encryption.SaltAndSha256Encrypt("m1234567", email),
		})
	}
	return
}
