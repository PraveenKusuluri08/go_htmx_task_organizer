package userhandlers

import (
	"fmt"
	"net/http"

	"github.com/Praveenkusuluri08/models"
	"github.com/Praveenkusuluri08/store"
	"github.com/Praveenkusuluri08/utils"
	"github.com/Praveenkusuluri08/view"
	"github.com/Praveenkusuluri08/view/components"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func HandleLoginPage(c *gin.Context) {
	cookie, _ := c.Cookie("token")
	fmt.Println(cookie)
	if cookie != "" {
		c.Redirect(http.StatusFound, "/")
	} else {

		utils.TemplateRenderer(c, 302, view.Base(view.Login(), false))
	}
}

func HandleLogout(c *gin.Context) {
	c.Header("HX-Redirect", "/login")
	c.Status(http.StatusOK)
	c.SetCookie("token", "", -1, "/", "", false, true)
}

func HandleSignupPage(c *gin.Context) {
	utils.TemplateRenderer(c, 302, view.Base(view.Signup(), false))
}

func HandleLoginApi() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := c.Request.ParseForm()
		if err != nil {
			utils.TemplateRenderer(c, 302, components.Error("Invalid form data"))
			return
		}
		email := c.PostForm("email")
		password := c.PostForm("password")

		fmt.Println(email, password)

		if email == "" || password == "" {
			badRequestResponse := utils.ErrorHandler{
				StatusCode: 400,
				Message:    "Username and password are required",
			}
			c.JSON(badRequestResponse.StatusCode, badRequestResponse)
			return
		}

		//check if the email is valid or not

		isValidEmail := utils.ValidateEmail(email)
		if !isValidEmail {
			utils.TemplateRenderer(c, 302, components.Error("Invalid email address"))
			return
		}

		isExists := store.CheckUserExistsByEmail(email)
		fmt.Println("isExists", isExists)
		if !isExists {
			c.Writer.WriteHeader(http.StatusUnauthorized)
			utils.TemplateRenderer(c, 302, components.Error("User does not exists"))
			return
		}
		userData, _ := store.GetUserByUsername(email)
		matchstring, isMatch := utils.CompareHashAndPassword(userData.Password, password)
		fmt.Println("matchstring", matchstring, isMatch)
		if !isMatch {
			c.Writer.WriteHeader(http.StatusUnauthorized)
			utils.TemplateRenderer(c, 302, components.Error("Password does not match"))
			return
		}
		// check the current passwor with the password present in the database if the both passwords are same then go to the next process if not make the user to the same page.
		token, _ := utils.GenerateToken(*userData)
		c.SetCookie("token", token, 3600, "/", "", true, true)
		c.Header("HX-Redirect", "/")
		c.Status(http.StatusOK)
	}
}

func HandleSignupApi() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := c.Request.ParseForm()
		if err != nil {
			utils.TemplateRenderer(c, 302, view.Base(components.Error("Invalid form data"), false))
			return
		}
		// values := c.Request.Form
		// fmt.Println("values", values)
		// for i, v := range values {
		//     fmt.Printf("%s: %s\n", i, v[0])
		// }
		firstname := c.PostForm("firstname")
		lastname := c.PostForm("lastname")
		email := c.PostForm("email")
		password := c.PostForm("password")
		confirmPassword := c.PostForm("confirmPassword")

		fmt.Println(firstname, lastname, email, password, confirmPassword)

		if email == "" || password == "" || confirmPassword == "" || firstname == "" || lastname == "" {
			badRequestResponse := utils.ErrorHandler{
				StatusCode: 400,
				Message:    "FirstName, LastName, Email, password and confirm password are required",
			}
			c.JSON(badRequestResponse.StatusCode, badRequestResponse)
		}
		isExists := store.CheckUserExistsByEmail(email)
		if isExists {
			c.Writer.WriteHeader(http.StatusUnauthorized)
			utils.TemplateRenderer(c, 302, components.Error("User already exists"))
			return
		}
		if password != confirmPassword {
			c.Writer.WriteHeader(http.StatusUnauthorized)
			utils.TemplateRenderer(c, 302, components.Error("Passwords do not match"))
			return
		}

		hashedPassword, err := utils.HashPassword(password)
		if err != nil {
			c.Writer.WriteHeader(http.StatusInternalServerError)
			utils.TemplateRenderer(c, 302, components.Error("Failed to hash password"))
			return
		}

		users := &models.User{}
		users.ID = uuid.New().String()
		users.Firstname = firstname
		users.Lastname = lastname
		users.Email = email
		users.Password = hashedPassword
		msg, err := store.CreateUser(users)
		if err != nil {
			c.Writer.WriteHeader(http.StatusInternalServerError)
			utils.TemplateRenderer(c, 302, components.Error(err.Error()))
			return
		}
		fmt.Println(msg)
		c.JSON(200, gin.H{"success": true, "message": "User created successfully"})
	}
}
