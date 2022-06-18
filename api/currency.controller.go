package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addCurrencyRoutes(rg *gin.RouterGroup) {

	user := rg.Group("/user/:userId")
	user.GET("/currencies/:direction", getUserCurrencies)
}

func getUserCurrencies(c *gin.Context) {
	c.JSON(http.StatusOK, GetPage(c.Param("userId"), c.Param("direction"), c.Query("offset"), c.Query("limit")))
}

func exchange(c *gin.Context) {

}
