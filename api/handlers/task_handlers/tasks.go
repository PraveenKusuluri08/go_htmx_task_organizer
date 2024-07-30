package task_handlers

import (
	"fmt"

	"github.com/Praveenkusuluri08/models"
	"github.com/Praveenkusuluri08/store"
	"github.com/Praveenkusuluri08/utils"
	"github.com/Praveenkusuluri08/view"
	"github.com/Praveenkusuluri08/view/components"
	"github.com/gin-gonic/gin"
)

func HandleTaskPage(c *gin.Context) {
	isLoggedIn := c.GetBool("isLoggedIn")
	fmt.Println("isLoggedIn", isLoggedIn)

	if isLoggedIn {
		utils.TemplateRenderer(c, 200, view.Base(view.Home(), isLoggedIn))
	}
}

func HandleGetTasks() gin.HandlerFunc {
	return func(c *gin.Context) {

		isLoggedIn := c.GetBool("isLoggedIn")
		if !isLoggedIn {
			utils.TemplateRenderer(c, 302, components.Error("Unauthorized access"))
			return
		}
		userID := c.GetString("uid")
		fmt.Println("userID", userID)

		tasks, err := store.GetTasks(userID)
		if err != nil {
			fmt.Println(err)
			utils.TemplateRenderer(c, 302, components.Error("Failed to retrieve tasks! Something is wrong"))
			return
		}
		tasksList := make([]models.Tasks, 0, len(tasks))
		for _, task := range tasks {
			tasksList = append(tasksList, *task)
		}
		fmt.Println("tasks", tasksList)
		for _, val := range tasksList {
			fmt.Println("tasks🚀", val.Title)
		}
		utils.TemplateRenderer(c, 200, view.Base(view.Tasks(tasksList), isLoggedIn))
	}
}

func HandleCreateTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := c.Request.ParseForm()
		if err != nil {
			utils.TemplateRenderer(c, 302, components.Error("Invalid form data"))
			return
		}
		title := c.PostForm("title")
		description := c.PostForm("description")
		priority := c.PostForm("priority")
		dueDate := c.PostForm("due_date")
		userID := c.GetString("uid")
		fmt.Println(title, description, priority, dueDate, userID)
		if title == "" || description == "" || priority == "" || dueDate == "" || userID == "" {
			utils.TemplateRenderer(c, 302, components.Error("All fields are required"))
			return
		}
		// Create task
		tasks := &models.Tasks{}
		tasks.Title = title
		tasks.Description = description
		tasks.Priority = priority
		tasks.DueDate = string(dueDate)
		tasks.UserID = string(userID)

		fmt.Printf("The type of duedate is %T", dueDate)

		msg, err := store.CreateTask(tasks)
		if err != nil {
			utils.TemplateRenderer(c, 302, components.Error(err.Error()))
			return
		}
		fmt.Println(msg)
		utils.TemplateRenderer(c, 302, components.Alert("Task created successfully"))
	}
}
