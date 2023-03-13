package controller

import (
	Account "cleverbank/internal/core/domain/account"
	"cleverbank/internal/infra/primary/middleware"
	"net/http"

	"cleverbank/internal/core/usecase"
	"github.com/gin-gonic/gin"
)

type AccountCreationController struct {
	accountCreationUseCase usecase.AccountCreationUseCase
}

func (e *AccountCreationController) RunController(r *gin.Engine) {

	r.POST("/v1/createaccount", func(gc *gin.Context) {
		authorized, err := middleware.IsAuthorized(gc.GetHeader("Authorization"))

		if !authorized {
			gc.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "Error", "message": err.Error()})
			return
		}

		var req Account.NewAccountReq
		gc.BindJSON(&req)

		response, err := e.accountCreationUseCase.Handle(req)

		if err == nil {
			gc.JSON(http.StatusOK, response)
		} else {
			gc.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "Error", "message": err.Error()})
		}
	})
}
