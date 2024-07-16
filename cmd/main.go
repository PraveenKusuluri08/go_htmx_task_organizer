package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/Praveenkusuluri08/api/routes"
	dbconfig "github.com/Praveenkusuluri08/dbConfig"
	"github.com/Praveenkusuluri08/middleware"
	"github.com/Praveenkusuluri08/store"
	"github.com/Praveenkusuluri08/utils"
	"github.com/Praveenkusuluri08/view"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}
	dbconfig.DBConnect()
}

var wg = sync.WaitGroup{}

func main() {
	r := gin.Default()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	dbconfig.DBConnect()
	routes.UserRoutes(r)
	r.Use(gin.Logger())

	// wg.Add(1)
	// go func() {
	err := store.CreateTables()
	if err != nil {
		log.Fatalf("Failed to create tables: %v", err)
	}
	log.Println("Tables created successfully.")
	// 	wg.Done()
	// }()

	r.Use(middleware.AuthMiddleware())
	r.GET("/", func(c *gin.Context) {
		isLoggedIn := c.GetBool("isLoggedIn")
		fmt.Println("isLoggedIn", isLoggedIn)

		if isLoggedIn {
			utils.TemplateRenderer(c, 200, view.Base(view.Home(), isLoggedIn))
		}
	})

	log.Printf("Starting server on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
	wg.Wait()
}
