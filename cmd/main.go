package main

import (
	"embed"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/Praveenkusuluri08/api/routes"
	dbconfig "github.com/Praveenkusuluri08/dbConfig"
	"github.com/Praveenkusuluri08/store"
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

var staticFiles embed.FS

func main() {
	r := gin.Default()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.StaticFS("/static/css/", http.FS(staticFiles))
	r.StaticFS("/static/javascript/", http.FS(staticFiles))
	dbconfig.DBConnect()
	routes.UserRoutes(r)
	routes.TaskRoutes(r)
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

	log.Printf("Starting server on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
	wg.Wait()
}
