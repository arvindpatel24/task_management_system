package task

import (
	"database/sql"
	"fmt"
)

type Repository interface {
	Create(task Task) (Task, error)
	GetAll() ([]Task, error)
	GetByID(id int64) (Task, error)
	Update(id int64, task Task) (Task, error)
	Delete(id int64) error
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

	query := "INSERT INTO " + DB_NAME + " (title, description, status) VALUES (?, ?, ?)"
	result, err := r.db.Exec(query, task.Title, task.Description, parsedStatus)
	if err != nil {
		fmt.Println(err)
		return Task{}, err
	}

	id, _ := result.LastInsertId()
	task.ID = id
	return task, nil
}

// Fetch all tasks from database
func (r *taskRepository) GetAll() ([]Task, error) {
	query := "SELECT id, title, description, status, created_at, updated_at FROM " + DB_NAME
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

// Get task by its ID
func (r *taskRepository) GetByID(id int64) (Task, error) {
	query := "SELECT id, title, description, status, created_at, updated_at FROM " + DB_NAME + " WHERE id = ?"
	row := r.db.QueryRow(query, id)
	var task Task
	if err := row.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt); err != nil {
		return Task{}, err
	}
	return task, nil
}

// Update task in the database
func (r *taskRepository) Update(id int64, task Task) (Task, error) {
	query := "UPDATE " + DB_NAME + " SET title = ?, description = ?, status = ? WHERE id = ?"
	_, err := r.db.Exec(query, task.Title, task.Description, task.Status, id)
	if err != nil {
		return Task{}, err
	}
	task.ID = id
	return task, nil
}

// Delete task by its ID in the database
func (r *taskRepository) Delete(id int64) error {
	query := "DELETE FROM " + DB_NAME + " WHERE id = ?"
	_, err := r.db.Exec(query, id)
	return err
}
