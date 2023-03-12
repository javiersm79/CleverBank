package controller

import (
	"net/http"

	"cleverbank/internal/core/usecase"
	"github.com/gin-gonic/gin"
)

type AccountMovementController struct {
	accountMovementUseCase usecase.AccountMovementUseCase
}

func (e *AccountMovementController) RunController(r *gin.Engine) {

	r.GET("/v1/deposit", func(gc *gin.Context) {
		response, err := e.accountMovementUseCase.Handle("12345")

		if err == nil {
			gc.JSON(http.StatusOK, response)
		} else {
			gc.AbortWithStatus(http.StatusNotFound)
		}
	})
}
