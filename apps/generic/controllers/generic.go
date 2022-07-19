package controllers

import (
	"boilerplate/core/infrastructures"
	"boilerplate/core/interfaces"
	_ "boilerplate/core/models"
	"boilerplate/core/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GenericController struct {
	logger interfaces.Logger
	env    *infrastructures.Env
}

func NewGenericController(logger *infrastructures.Logger,
	env *infrastructures.Env,
) *GenericController {
	return &GenericController{
		logger: logger,
		env:    env,
	}
}

// @BasePath /api
// @Summary ping
// @Schemes
// @Description do ping
// @Tags generic
// @Accept json
// @Produce json
// @Success 200 {object} swagger.PingResponse
// @Router /ping [get]
func (uc GenericController) Ping(ctx *gin.Context) {
	responses.JSON(ctx, http.StatusOK, gin.H{"pingpong": "🏓🏓🏓🏓🏓🏓"}, "pong")
}
