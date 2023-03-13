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
	for _, o := range c.controllers {
		o.RunController(router)
	}
	configuration(router).ListenAndServe()
}

func configuration(router *gin.Engine) *http.Server {
	/*port := ":" + strconv.Itoa(prop.GetServerPortProperties().Port)
	ReadTimeout := time.Duration(prop.GetServerPortProperties().ReadTimeout) * time.Second
	WriteTimeout := time.Duration(prop.GetServerPortProperties().WriteTimeout) * time.Second*/
	port := ":" + strconv.Itoa(3000)
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
