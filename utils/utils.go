package utils

import (
	"fmt"
	"log"
	"os"
	"regexp"
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
	Email string
	Uid   string
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
		Email: user.Email,
		Uid:   user.ID,
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

func HashPassword(password string) (string, error) {
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	hashedPassword := string(hashBytes)
	return hashedPassword, nil
}

func ValidateEmail(email string) bool {
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
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
