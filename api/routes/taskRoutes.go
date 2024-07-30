package routes

import (
	"github.com/Praveenkusuluri08/api/handlers/task_handlers"
	"github.com/Praveenkusuluri08/middleware"
	"github.com/gin-gonic/gin"
)

func TaskRoutes(r *gin.Engine) {
	r.Use(middleware.AuthMiddleware())
	r.GET("/", task_handlers.HandleTaskPage)
	r.Use(middleware.AuthMiddleware1())
	r.GET("/tasks", task_handlers.HandleGetTasks())

	r.POST("/create", task_handlers.HandleCreateTask())
}
