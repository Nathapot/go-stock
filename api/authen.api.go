package api

import (
	"github.com/Nathapot/go-stock/db"
	"github.com/Nathapot/go-stock/interceptor"
	"github.com/Nathapot/go-stock/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"time"
)

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func login(c *gin.Context) {
	var user models.User

	if c.ShouldBind(&user) == nil {
		var queryUser models.User
		if err := db.GetDB().First(&queryUser, "username = ?", user.Username).Error; err != nil {
			c.JSON(200, gin.H{"result": "nok", "error": err})
		} else if checkPasswordHash(user.Password, queryUser.Password) == false {
			c.JSON(200, gin.H{"result": "nok", "error": "invalid password"})
		} else {
			token := interceptor.JwtSign(queryUser)
			c.JSON(200, gin.H{"result": "ok", "token": token})
		}

	} else {
		c.JSON(401, gin.H{"status": "unable to bind data"})
	}
}

func register(c *gin.Context) {
	var user models.User
	if c.ShouldBind(&user) == nil {
		user.Password, _ = hashPassword(user.Password)
		user.CreatedAt = time.Now()
		err := db.GetDB().Create(&user).Error
		if err != nil {
			c.JSON(200, gin.H{"result": "nok", "error": err})
		} else {
			c.JSON(200, gin.H{"result": "ok", "data": user})
		}
	} else {
		c.JSON(401, gin.H{"status": "unable to bind data"})
	}
}

func SetupAuthenAPI(router *gin.Engine) {
	authenAPI := router.Group("/api/v2")
	{
		authenAPI.POST("/login", login)
		authenAPI.POST("/register", register)
	}
}
