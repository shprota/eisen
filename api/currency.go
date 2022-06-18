package api

import (
	"fmt"
	"net/http"
	"strconv"

	"eisen/api/data"

	"github.com/gin-gonic/gin"
)

func addCurrencyRoutes(rg *gin.RouterGroup) {

	user := rg.Group("/user/:userId")
	user.GET("/currencies", getUserCurrencies)
}

func getUserCurrencies(c *gin.Context) {
	user := c.Param("userId")
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 10
	}

	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		offset = 0
	}

	res := fmt.Sprintf("WIP. User: %s, limit: %d, offset: %d", user, limit, offset)

	c.JSON(http.StatusOK, gin.H{"error": res, "currencies": data.Currencies[offset:limit]})
}
