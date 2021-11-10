package api

import (
	"net/http"
	"strconv"

	"github.com/Nathapot/go-stock/interceptor"
	"github.com/Nathapot/go-stock/models"
	"github.com/gin-gonic/gin"
)

func getProduct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": "get product", "username": c.GetString("jwt_username"), "level": c.GetString("jwt_level")})
}

func createProduct(c *gin.Context) {
	product := models.Product{}
	product.Name = c.PostForm("name")
	product.Stock, _ = strconv.ParseInt(c.PostForm("stock"), 10, 64)
	product.Price, _ = strconv.ParseFloat(c.PostForm("price"), 64)
	c.JSON(http.StatusOK, gin.H{"result": product})
}

// SetupProductAPI - call this method to setup product route group
func SetupProductAPI(router *gin.Engine) {
	productAPI := router.Group("/api/v2")
	{
		productAPI.GET("/product", interceptor.JwtVerify, getProduct)
		productAPI.POST("/product", createProduct)
	}
}
