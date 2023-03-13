package controller

import (
	"cleverbank/internal/core/domain/account"
	"cleverbank/internal/infra/primary/middleware"
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
		authorized, err := middleware.IsAuthorized(gc.GetHeader("Authorization"))

		if !authorized {
			gc.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "Error", "message": err.Error()})
			return
		}
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

	r.POST("/v1/withdrawal", func(gc *gin.Context) {
		authorized, err := middleware.IsAuthorized(gc.GetHeader("Authorization"))

		if !authorized {
			gc.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "Error", "message": err.Error()})
			return
		}

		var req account.AccountMovementRequest
		req.Type = "withdrawal"
		gc.BindJSON(&req)

		if req.SourceAccountNumber == "" {
			fmt.Errorf("Error in AccountInfoController /v1/withdrawal - missing account number")
			gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "Error", "message": "missing account number"})
			return
		}

		if req.Amount <= 0 {
			fmt.Errorf("Error in AccountInfoController /v1/withdrawal - Amount must be mayor than 0")
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

	r.POST("/v1/transfer", func(gc *gin.Context) {
		authorized, err := middleware.IsAuthorized(gc.GetHeader("Authorization"))

		if !authorized {
			gc.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "Error", "message": err.Error()})
			return
		}

		var req account.AccountMovementRequest
		req.Type = "transfer"
		gc.BindJSON(&req)

		if req.SourceAccountNumber == "" {
			fmt.Errorf("Error in AccountInfoController /v1/transfer - missing destiny account number")
			gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "Error", "message": "missing destiny account number"})
			return
		}

		if req.DestinyAccountNumber == "" {
			fmt.Errorf("Error in AccountInfoController /v1/transfer - missing destiny account number")
			gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "Error", "message": "missing destiny account number"})
			return
		}

		if req.Amount <= 0 {
			fmt.Errorf("Error in AccountInfoController /v1/transfer - Amount must be mayor than 0")
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
