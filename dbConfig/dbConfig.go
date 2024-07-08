package dbconfig

import (
	"fmt"
	"os"

	"github.com/Praveenkusuluri08/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnect() {
	dsn := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	fmt.Println("Connection Opened to Database")
	DB = db
	db.AutoMigrate(&models.User{})
}
