package store

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	dbconfig "github.com/Praveenkusuluri08/dbConfig"
	"github.com/Praveenkusuluri08/models"
	"github.com/Praveenkusuluri08/utils"
)

func CreateTables() error {
	log.Println("Attempting to create users table...")

	userTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id TEXT PRIMARY KEY,
		firstname TEXT,
		lastname TEXT,
		email TEXT UNIQUE,
		password TEXT,
		created_at TIMESTAMP DEFAULT NOW()
	)
	`
	_, err := dbconfig.DB.Exec(userTableQuery)
	if err != nil {
		log.Printf("Error creating users table: %v", err)
		return err
	}
	log.Println("Users table created successfully.")

	log.Println("Attempting to create tasks table...")

	tasksTableQuery := `
	CREATE TABLE IF NOT EXISTS tasks (
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
	_, err = dbconfig.DB.Exec(tasksTableQuery)
	if err != nil {
		log.Printf("Error creating tasks table: %v", err)
		return err
	}
	log.Println("Tasks table created successfully.")

	log.Println("Tables created successfully.")
	return nil
}

func CreateUser(user *models.User) (msg string, err error) {
	if user == nil {
		return "", errors.New("user cannot be nil")
	}

	// Log user information (except the password) for debugging
	log.Printf("Creating user: %+v", *user)

	user.ID = utils.GenerateUUID()

	// Log user information after hashing the password and generating UUID
	log.Printf("User after processing: %+v", *user)

	insertQuery := "INSERT INTO users (id, firstname, lastname, email, password) VALUES ($1, $2, $3, $4, $5)"
	result, err := dbconfig.DB.Exec(insertQuery, user.ID, user.Firstname, user.Lastname, user.Email, user.Password)

	rowsAffected, rowsAffectedErr := result.RowsAffected()
	if err != nil {
		return "", rowsAffectedErr
	}

	if rowsAffected == 0 {
		return "", errors.New("no rows were inserted")
	}

	return fmt.Sprintf("User %s successfully", user.Email), nil
}

func CheckUserExistsByEmail(email string) bool {
	userExistsQuery := "SELECT 1 FROM users WHERE email = $1"
	var exists bool
	err := dbconfig.DB.QueryRow(userExistsQuery, email).Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		fmt.Println("Error retrieving user:", err)
		return false
	}
	return true
}
func GetUserByUsername(email string) (user *models.User, err error) {
	user = &models.User{}
	userExistsQuery := "SELECT id,firstname,lastname,email,password FROM users WHERE email = $1"
	rows := dbconfig.DB.QueryRow(userExistsQuery, email)
	fmt.Println(rows)
	err = rows.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	fmt.Println(*user)
	return user, nil
}
