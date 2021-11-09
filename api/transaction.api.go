package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getTransaction(c *gin.Context) {
	c.String(http.StatusOK, "List transaction")
}

func createTransaction(c *gin.Context) {
	c.String(http.StatusOK, "Create transaction")

}

func SetupTransactionAPI(router *gin.Engine) {
	transactionAPI := router.Group("/api/v2")
	{
		transactionAPI.GET("/transaction", getTransaction)
		transactionAPI.POST("/transaction", createTransaction)

	}
}
