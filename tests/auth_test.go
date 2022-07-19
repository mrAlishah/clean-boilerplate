package tests

import (
	"boilerplate/apps/user/controllers"
	"boilerplate/apps/user/services"
	"boilerplate/core/infrastructures"
	"boilerplate/tests/mocks"
)

func test_register_good_data() {
	env := infrastructures.GetEnv()
	userRepository := mocks.NewUserRepository()
	logger := infrastructures.NewLogger(env)
	encryption := infrastructures.NewEncryption(logger, env)
	db := infrastructures.GetDB(env)

	userService := services.NewUserService(userRepository, db, logger, encryption)
	userController := controllers.NewUserController(logger, env, userService)
	c := userController.CreateUser()
}
