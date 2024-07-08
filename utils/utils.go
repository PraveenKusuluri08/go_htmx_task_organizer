package utils

import (
	"reflect"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func TemplateRenderer(c *gin.Context, status int, template templ.Component) error {
	c.Status(status)
	return template.Render(c.Request.Context(), c.Writer)
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
