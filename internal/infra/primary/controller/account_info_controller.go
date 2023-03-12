package controller

import (
	"fmt"
	"net/http"

	"cleverbank/internal/core/usecase"
	"github.com/gin-gonic/gin"
)

type AccountInfoController struct {
	accountInfoUseCase usecase.AccountInfoUseCase
}

func (e *AccountInfoController) RunController(r *gin.Engine) {

	r.GET("/v1/balance", func(gc *gin.Context) {
		queryParams := gc.Request.URL.Query()

		if queryParams.Get("account") == "" {
			fmt.Errorf("Error in AccountInfoController /v1/balance - missing account number")
			gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "Error", "message": "missing account number"})
			return
		}

		response, err := e.accountInfoUseCase.Handle(queryParams.Get("account"))

		if err == nil {
			gc.JSON(http.StatusOK, response)
		} else {
			gc.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "Error", "message": err.Error()})
		}
	})
}
