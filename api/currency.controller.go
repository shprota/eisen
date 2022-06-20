package api

import (
	"eisen/api/data"
	"net/http"

	"github.com/gin-gonic/gin"
)

func addCurrencyRoutes(rg *gin.RouterGroup) {

	user := rg.Group("/user/:userId")
	user.GET("/currencies/:direction", getUserCurrencies)
	user.POST("/exchange/:direction", exchange)
	user.GET("/balances", getBalances)
}

func getUserCurrencies(c *gin.Context) {
	var params data.DirectionUrlParams
	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, GetPage(c.Param("userId"), c.Param("direction"), c.Query("offset"), c.Query("limit")))
}

func getBalances(c *gin.Context) {
	c.JSON(http.StatusOK, GetBalances(c.Param("userId"), c.Query("offset"), c.Query("limit")))
}

func exchange(c *gin.Context) {
	var params data.DirectionUrlParams
	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	data := data.ExchangeDto{}
	c.BindJSON(data)
	c.JSON(http.StatusOK, Exchange(c.Param("userId"), c.Param("Direction"), data))
}
