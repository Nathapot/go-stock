package api

import (
	"github.com/Nathapot/go-stock/interceptor"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getProduct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": "get product", "username": c.GetString("jwt_username"), "level": c.GetString("jwt_level")})
}

func createProduct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": "create product"})

}

// SetupProductAPI - call this method to setup product route group
func SetupProductAPI(router *gin.Engine) {
	productAPI := router.Group("/api/v2")
	{
		productAPI.GET("/product", interceptor.JwtVerify, getProduct)
		productAPI.POST("/product", createProduct)
	}
}
