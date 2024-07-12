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
		for _, routePath := range noAuthRoutePaths {
			if routePath == path {
				c.Next()
				return
			}
		}
		fmt.Println("token: ", cookie)
		fmt.Println(cookie == "")
		if cookie == "" && (path != "/signup" && path != "/login") {
			utils.TemplateRenderer(c, 302, view.Base(view.PageNotfound("User not Loggedin:UnAuthorized"), false))
			return
		}
		claims, err := utils.ValidateToken(cookie)
		if err != "" && (path != "/signup" && path != "/login") {
			utils.TemplateRenderer(c, 302, view.Base(view.PageNotfound("Something went wrong on the server side, Please Reload"), false))
			c.JSON(http.StatusBadRequest, "Bearer token not found")
			return
		}
		c.Set("email", claims.Email)
		c.Set("uid", claims.Uid)
		c.Set("username", claims.Username)
		c.Set("isLoggedIn", true)

		if path == "/" {
			c.Redirect(http.StatusFound, "/")
		}

		c.Next()
	}
}
