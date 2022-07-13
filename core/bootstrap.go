package core

import (
	"boilerplate/core/infrastructures"
	"boilerplate/core/responses"
	"boilerplate/core/responses/validators"
	"context"
	"fmt"
	"net/http"
	"runtime"

	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"

	"go.uber.org/fx"
)

var BootstrapModule = fx.Options(
	infrastructures.Module,
	RoutesModule,
	validators.Module,
	ServiceModule,
	RepositoryModule,
	ControllerModule,
	fx.Invoke(bootstrap),
)

func bootstrap(lifecycle fx.Lifecycle,
	logger *infrastructures.Logger,
	router *infrastructures.Router,
	env *infrastructures.Env,
	routes Routes,
	validators validators.Validators,
) {
	//recover unwanted 500 errors
	router.Gin.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if recovered != nil {
			switch e := recovered.(type) {
			case string:
				logger.Warning("recovered (string) panic:", e)
			case runtime.Error:
				logger.Warning("recovered (runtime.Error) panic:", e.Error())
			case error:
				logger.Warning("recovered (error) panic:", e.Error())
			default:
				logger.Warning("recovered (default) panic:", e)
			}
			responses.ErrorJSON(c, http.StatusInternalServerError, gin.H{}, "Sorry an error occured!")
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	}))

	//recover unwanted 500 errors to sentery
	router.Gin.Use(sentrygin.New(sentrygin.Options{
		Repanic: true,
	}))

	fmt.Println("xxxxxxvsdgfgds")
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Info("Starting Application🔥💝😈")
			logger.Info("------------------------")
			logger.Info(fmt.Sprintf("------ %s  ------", env.AppName))
			logger.Info("------------------------")
			fmt.Println(env.ServerPort)
			go func() {
				validators.Setup()
				routes.Setup()
				//docs.SwaggerInfo.BasePath = "/api"
				//router.Gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

				//middlewares.Setup()
				if env.ServerPort == "" {
					router.Gin.Run(":8000")
				} else {
					router.Gin.Run(":" + env.ServerPort)
				}
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			fmt.Println("Stopping Application 📛") //log info
			return nil
		},
	})
}
