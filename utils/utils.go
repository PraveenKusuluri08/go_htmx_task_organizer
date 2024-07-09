package utils

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"time"

	"github.com/Praveenkusuluri08/models"
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var Secret_Key = os.Getenv("SECRET_KEY")

type SignInDetails struct {
	Email    string
	Uid      string
	Username string
	jwt.StandardClaims
}

func TemplateRenderer(c *gin.Context, status int, template templ.Component) error {
	c.Status(status)
	return template.Render(c.Request.Context(), c.Writer)
}

func ValidateToken(token string) (claims *SignInDetails, msg string) {
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

func GenerateToken(user models.User) (string, error) {
	claims := &SignInDetails{
		Email:    user.Email,
		Uid:      user.ID,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(Secret_Key))
	if err != nil {
		log.Panicln(err)
		return "", err
	}
	return token, err
}

func HashPassword(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return reflect.ValueOf(hash).String()
}

func CompareHashAndPassword(hash string, password string) (string, bool) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return "Password does not match", false
	}
	return "Password matches", true
}

func GenerateUUID() string {
	id := uuid.New()
	return id.String()
}
