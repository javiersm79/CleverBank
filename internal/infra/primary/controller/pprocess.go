package controller

import (
	"fmt"
	"github.com/rendis/lightms"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ControllerRunnable interface {
	RunController(r *gin.Engine)
}

type GinController struct {
	controllers []ControllerRunnable
}

func GetControllerInstance(c []ControllerRunnable) lightms.PrimaryProcess {
	return &GinController{c}
}

func (c *GinController) Start() {
	router := gin.Default()
	router.Use(corsMiddleware())
	for _, o := range c.controllers {
		o.RunController(router)
	}
	configuration(router).ListenAndServe()
}

func configuration(router *gin.Engine) *http.Server {

	port := ":" + strconv.Itoa(4000)
	ReadTimeout := time.Duration(1000) * time.Second
	WriteTimeout := time.Duration(1000) * time.Second
	fmt.Printf("Start Server. Port = '%s', ReadTimeout = %s, WriteTimeout = %s. \n\n", port, ReadTimeout, WriteTimeout)

	return &http.Server{
		Addr:         port,
		Handler:      router,
		ReadTimeout:  ReadTimeout,
		WriteTimeout: WriteTimeout,
	}
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
