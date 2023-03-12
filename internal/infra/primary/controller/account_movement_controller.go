package controller

import (
	"cleverbank/internal/core/domain/account"
	"fmt"
	"net/http"

	"cleverbank/internal/core/usecase"
	"github.com/gin-gonic/gin"
)

type AccountMovementController struct {
	accountMovementUseCase usecase.AccountMovementUseCase
}

func (e *AccountMovementController) RunController(r *gin.Engine) {

	r.POST("/v1/deposit", func(gc *gin.Context) {
		var req account.AccountMovementRequest
		req.Type = "deposit"
		gc.BindJSON(&req)

		if req.DestinyAccountNumber == "" {
			fmt.Errorf("Error in AccountInfoController /v1/deposit - missing destiny account number")
			gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "Error", "message": "missing destiny account number"})
			return
		}

		if req.Amount <= 0 {
			fmt.Errorf("Error in AccountInfoController /v1/deposit - Amount must be mayor than 0")
			gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "Error", "message": "Amount must be mayor than 0"})
			return
		}

		response, err := e.accountMovementUseCase.Handle(req)

		if err == nil {
			gc.JSON(http.StatusOK, response)
		} else {
			gc.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "Error", "message": err.Error()})
		}
	})
}
