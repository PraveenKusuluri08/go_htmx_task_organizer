package routes

import (
	userhandlers "github.com/Praveenkusuluri08/api/handlers/user_handlers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.GET("/login", userhandlers.HandleLoginPage)

	router.GET("/signup", userhandlers.HandleSignupPage)

	//api fucns
	router.POST("/login", userhandlers.HandleLoginApi())

	router.POST("/signup", userhandlers.HandleSignupApi())

	router.GET("/logout", userhandlers.HandleLogout)
}
