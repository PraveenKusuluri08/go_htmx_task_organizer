package store

import (
	"errors"
	"fmt"

	dbconfig "github.com/Praveenkusuluri08/dbConfig"
	"github.com/Praveenkusuluri08/models"
	"github.com/Praveenkusuluri08/utils"
)

func CreateTables() error {
	userTableQuery := `
	CREATE TABLE IF NOT EXISTS users(
		id TEXT PRIMARY KEY,	
		username TEXT UNIQUE,
		email TEXT UNIQUE,
    	password TEXT,
		created_at TIMESTAMP DEFAULT NOW()
	)
	`
	tableInfo := dbconfig.DB.Exec(userTableQuery)
	if tableInfo.Error != nil {
		return tableInfo.Error
	}
	fmt.Println(tableInfo)

	tasksTableQuery := `
	CREATE TABLE IF NOT EXISTS tasks(
	    id TEXT PRIMARY KEY,
	    title TEXT,
	    description TEXT,
	    user_id TEXT,
	    priority BOOLEAN,
	    due_date TEXT,
	    completed BOOLEAN DEFAULT FALSE,
	    completed_at TIMESTAMP,
	    created_at TIMESTAMP DEFAULT NOW(),
	    FOREIGN KEY (user_id) REFERENCES users(id)
	)
	`
	tasksTableInfo := dbconfig.DB.Exec(tasksTableQuery)
	if tasksTableInfo.Error != nil {
		return tasksTableInfo.Error
	}
	fmt.Println(tasksTableInfo)
	fmt.Println("Tables created successfully")
	return nil
}

func CreateUser(user *models.User) (msg string, err error) {
	if user == nil {
		return "", errors.New("user cannot be nil")
	}
	hashedPassword := utils.HashPassword(user.Password)
	user.Password = hashedPassword
	user.ID = utils.GenerateUUID()
	insertQuery := "INSERT INTO users (id, username, email, password) VALUES (?, ?, ?, ?)"
	insertInfo := dbconfig.DB.Exec(insertQuery, user.ID, user.Username, user.Email, user.Password)
	if insertInfo.Error != nil {
		return "", insertInfo.Error
	}
	return fmt.Sprintf("User %s created successfully", user.Username), nil
}

func CheckUserExistsByUsername(username string) bool {
	userExistsQuery := "SELECT id FROM users WHERE username = ?"
	var user *models.User
	dbconfig.DB.Exec(userExistsQuery, username).Scan(&user)
	return user != nil
}

func GetUserByUsername(username string) (user *models.User, err error) {
	userExistsQuery := "SELECT * FROM users WHERE username = ?"
	dbconfig.DB.Exec(userExistsQuery, username).Scan(&user)
	return user, nil
}
