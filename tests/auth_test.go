package tests

import (
	"boilerplate/apps/user/controllers"
	"boilerplate/apps/user/services"
	"boilerplate/core/infrastructures"
	"boilerplate/core/utils"
	"boilerplate/tests/mocks"
)

func test_register_good_data() {
	env := utils.GetEnv()
	userRepository := mocks.NewUserRepository()
	logger := infrastructures.NewLogger(env)
	encryption := infrastructures.NewEncryption(logger, env)
	db := utils.GetDB(env)

	userService := services.NewUserService(userRepository, db, logger, encryption)
	userController := controllers.NewUserController(logger, env, userService)
	c := userController.CreateUser()
}
