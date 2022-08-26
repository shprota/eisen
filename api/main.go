package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func Run() {
	getRoutes()
	err := router.Run(":5000")
	if err != nil {
		fmt.Printf("Error occured: %s", err.Error())
	}
}

func getRoutes() {
	v1 := router.Group("/v1")
	addCurrencyRoutes(v1)
}
