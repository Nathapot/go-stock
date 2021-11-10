package interceptor

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Nathapot/go-stock/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secretKey = "87654321"

func JwtSign(payload models.User) string {
	atClaims := jwt.MapClaims{}
	// Payload begin
	atClaims["id"] = payload.ID
	atClaims["username"] = payload.Username
	atClaims["level"] = payload.Level
	// token expire in 15 mins
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	// Payload end
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, _ := at.SignedString([]byte(secretKey))
	return token
}

func JwtVerify(c *gin.Context) {
	tokenString := strings.Split(c.Request.Header["Authorization"][0], " ")[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims)
		c.Set("jwt_username", claims["username"])
		c.Set("jwt_level", claims["level"])
		c.Next()
	} else {
		c.JSON(http.StatusOK, gin.H{"result": "nok", "error": err})
		c.Abort()
	}
}
