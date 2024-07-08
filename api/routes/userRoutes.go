package routes

import (
	userhandlers "github.com/Praveenkusuluri08/api/handlers/user_handlers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.GET("/login", userhandlers.HandleLoginPage)
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the API",
		})
	})
}
