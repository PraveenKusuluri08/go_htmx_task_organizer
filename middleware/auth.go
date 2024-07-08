package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var Secret_Key = os.Getenv("SECRET_KEY")

type SignInDetails struct {
	Email    string
	Uid      string
	Username string
	jwt.StandardClaims
}

func validateToken(token string) (claims *SignInDetails, msg string) {
	var message string
	tokenString, err := jwt.ParseWithClaims(
		token,
		&SignInDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(Secret_Key), nil
		},
	)
	if err != nil {
		log.Fatal(err)
		return
	}
	claims, ok := tokenString.Claims.(*SignInDetails)

	if !ok {
		message = "token is expired"
		return nil, message
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		message = "Token is expired please check"
		return nil, message
	}
	fmt.Println(message)
	return claims, message
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, _ := c.Cookie("token")
		fmt.Println("token: ", cookie)
		if cookie == "" {
			c.Redirect(http.StatusFound, "/login")
			return
		}
		claims, err := validateToken(cookie)
		if err != "" {
			c.Redirect(http.StatusFound, "/login")

			return
		}
		c.Set("email", claims.Email)
		c.Set("uid", claims.Uid)
		c.Set("username", claims.Username)
		c.Set("isLoggedIn", true)
		c.Redirect(http.StatusFound, "/")
		c.Next()
	}
}
