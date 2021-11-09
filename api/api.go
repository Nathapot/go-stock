package api

import (
	"github.com/Nathapot/go-stock/db"
	"github.com/gin-gonic/gin"
)

// Setup - call this method to setup routes
func Setup(router *gin.Engine) {

	db.SetupDB()
	SetupAuthenAPI(router)
	SetupProductAPI(router)
	SetupTransactionAPI(router)
}
