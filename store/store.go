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
	    priority TEXT,
	    due_date TEXT,
	    completed BOOLEAN DEFAULT FALSE,
	    completed_at TEXT,
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

func CreateTask(task *models.Tasks) (string, error) {
	if task == nil {
		return "", errors.New("task cannot be nil")
	}

	log.Printf("Creating task: %+v", *task)

	task.ID = utils.GenerateUUID()
	task.CompletedAt = ""

	log.Printf("Task after processing: %+v", *task)

	insertQuery := "INSERT INTO tasks (id, title, description, user_id, priority, due_date, completed, completed_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"
	result, err := dbconfig.DB.Exec(insertQuery, task.ID, task.Title, task.Description, task.UserID, task.Priority, task.DueDate, false, task.CompletedAt)
	if err != nil {
		return "", fmt.Errorf("error executing query: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return "", fmt.Errorf("error getting rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return "", errors.New("no rows were inserted")
	}
	return fmt.Sprintf("Task %s created successfully", task.Title), nil
}

func GetTasks(uid string) ([]*models.Tasks, error) {
	fmt.Println("The uid is", uid)
	tasks := []*models.Tasks{}
	query := "SELECT id, title, description, user_id, priority, due_date, completed, completed_at, created_at FROM tasks WHERE user_id = $1"
	rows, err := dbconfig.DB.Query(query, uid)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		task := &models.Tasks{}
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.UserID, &task.Priority, &task.DueDate, &task.Completed, &task.CompletedAt, &task.CreatedAt)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		log.Println(err)
		return nil, err
	}
	// Print the tasks in a human-readable format
	fmt.Printf("Tasks: %+v\n", tasks)
	return tasks, nil
}
