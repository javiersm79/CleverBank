package controller

import (
	"net/http"

	"cleverbank/internal/core/usecase"
	"github.com/gin-gonic/gin"
)

type AccountInfoController struct {
	accountInfoUseCase usecase.AccountInfoUseCase
}

func (e *AccountInfoController) RunController(r *gin.Engine) {

	r.GET("/v1/balance", func(gc *gin.Context) {
		response, err := e.accountInfoUseCase.Handle("12345")

		if err == nil {
			gc.JSON(http.StatusOK, response)
		} else {
			gc.AbortWithStatus(http.StatusNotFound)
		}
	})
}
