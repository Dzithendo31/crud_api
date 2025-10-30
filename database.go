package main

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// InitDB initializes the SQLite database and creates the tasks table
func InitDB(filepath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		return nil, err
	}

	// Create tasks table
	createTableSQL := `CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT,
		status TEXT DEFAULT 'pending',
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// CreateTask inserts a new task into the database
func CreateTask(db *sql.DB, task *Task) error {
	query := `INSERT INTO tasks (title, description, status, created_at, updated_at) 
			  VALUES (?, ?, ?, ?, ?)`

	now := time.Now()
	result, err := db.Exec(query, task.Title, task.Description, task.Status, now, now)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	task.ID = int(id)
	task.CreatedAt = now
	task.UpdatedAt = now
	return nil
}

// GetAllTasks retrieves all tasks from the database
func GetAllTasks(db *sql.DB) ([]Task, error) {
	query := `SELECT id, title, description, status, created_at, updated_at FROM tasks`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

// GetTaskByID retrieves a single task by ID
func GetTaskByID(db *sql.DB, id int) (*Task, error) {
	query := `SELECT id, title, description, status, created_at, updated_at FROM tasks WHERE id = ?`

	var task Task
	err := db.QueryRow(query, id).Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

// UpdateTask updates an existing task
func UpdateTask(db *sql.DB, task *Task) error {
	query := `UPDATE tasks SET title = ?, description = ?, status = ?, updated_at = ? WHERE id = ?`

	task.UpdatedAt = time.Now()
	_, err := db.Exec(query, task.Title, task.Description, task.Status, task.UpdatedAt, task.ID)
	return err
}

// DeleteTask removes a task from the database
func DeleteTask(db *sql.DB, id int) error {
	query := `DELETE FROM tasks WHERE id = ?`
	_, err := db.Exec(query, id)
	return err
}
