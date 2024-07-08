package userhandlers

import (
	"github.com/Praveenkusuluri08/utils"
	"github.com/Praveenkusuluri08/view"
	"github.com/gin-gonic/gin"
)

func HandleLoginPage(c *gin.Context) {
	utils.TemplateRenderer(c, 302, view.Base(view.Login(), false))
}

func HandleLoginApi() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		if username == "" || password == "" {
			baedReequestResponse := utils.ErrorHandler{
				StatusCode: 400,
				Message:    "Username and password are required",
			}
			c.JSON(baedReequestResponse.StatusCode, baedReequestResponse)
			return
		}
		// Check if user exists and password is correct

	}
}
