package task

import (
	"database/sql"
	"fmt"
)

type Repository interface {
	Create(task Task) (Task, error)
}

type taskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) Repository {
	return &taskRepository{db: db}
}

const DB_NAME = "tasks"

// --------------------------- Status Validation -------------------------------------

// Validate method to check if a status is valid
func (s Status) IsValid() bool {
	return s == New || s == Started || s == Pending || s == Completed
}

// String method for displaying the status
func (s Status) String() string {
	return string(s)
}

// ParseStatus function to convert a string to a Status
func ParseStatus(status string) (Status, error) {
	s := Status(status)
	if !s.IsValid() {
		return "", fmt.Errorf("invalid status: %s", status)
	}
	return s, nil
}

// ----------------------------- DB Operations -----------------------------------------

// Create a new task in the database
func (r *taskRepository) Create(task Task) (Task, error) {
	parsedStatus, err := ParseStatus(string(task.Status))
	if err != nil {
		return Task{}, err
	}

	query := "INSERT INTO tasks (title, description, status) VALUES (?, ?, ?)"
	result, err := r.db.Exec(query, task.Title, task.Description, parsedStatus)
	if err != nil {
		fmt.Println(err)
		return Task{}, err
	}

	id, _ := result.LastInsertId()
	task.ID = id
	return task, nil
}
