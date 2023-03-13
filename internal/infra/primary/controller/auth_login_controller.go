package controller

import (
	"cleverbank/internal/core/domain/auth"
	"net/http"

	"cleverbank/internal/core/usecase"
	"github.com/gin-gonic/gin"
)

type AuthLoginController struct {
	loginUseCase usecase.LoginUseCase
}

func (e *AuthLoginController) RunController(r *gin.Engine) {

	r.POST("/v1/login", func(gc *gin.Context) {

		var req auth.LoginReq
		gc.BindJSON(&req)

		response, err := e.loginUseCase.Handle(req)

		if err == nil {
			gc.JSON(http.StatusOK, response)
		} else {
			gc.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "Error", "message": err.Error()})
		}
	})
}
