package middleware

import (
	"fmt"
	"net/http"

	"github.com/Praveenkusuluri08/utils"
	"github.com/Praveenkusuluri08/view"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, _ := c.Cookie("token")
		noAuthRoutePaths := []string{"/login", "/signup", "/forgotPassword"}
		path := c.Request.URL.Path
		fmt.Println("path: ", path)

		// Allow unauthenticated access to noAuthRoutePaths
		for _, routePath := range noAuthRoutePaths {
			if routePath == path {
				c.Next()
				return
			}
		}
		fmt.Println("token: ", cookie)
		fmt.Println(cookie == "")

		// If the user is not authenticated, redirect to login or show an error
		if cookie == "" {
			if path != "/login" && path != "/signup" {
				utils.TemplateRenderer(c, 302, view.Base(view.PageNotfound("User not Logged in: Unauthorized"), false))
				c.Abort()
				return
			}
		} else {
			claims, err := utils.ValidateToken(cookie)
			if err != "" {
				utils.TemplateRenderer(c, 302, view.Base(view.PageNotfound("Something went wrong on the server side, Please Reload"), false))
				c.JSON(http.StatusBadRequest, "Bearer token not found")
				c.Abort()
				return
			}

			c.Set("email", claims.Email)
			c.Set("uid", claims.Uid)
			c.Set("isLoggedIn", true)

			if path == "/login" || path == "/signup" {
				c.Header("HX-Redirect", "/")
				c.Status(http.StatusOK)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
func AuthMiddleware1() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, _ := c.Cookie("token")
		fmt.Println("Token: ", cookie)
		if cookie == "" {
			utils.TemplateRenderer(c, 302, view.Base(view.PageNotfound("User not Logged in: Unauthorized"), false))
			c.Abort()
		}
		claims, err := utils.ValidateToken(cookie)
		if err != "" {
			utils.TemplateRenderer(c, 302, view.Base(view.PageNotfound("Something went wrong on the server side, Please Reload"), false))
			c.JSON(http.StatusBadRequest, "COOKIE not found")
			c.Abort()
		}
		fmt.Println("Login User: ", claims.Uid)
		c.Set("email", claims.Email)
		c.Set("isLoggedIn", true)
		c.Set("uid", claims.Uid)
		c.Next()
	}
}
