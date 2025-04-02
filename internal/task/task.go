package task

type Task struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      Status `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type Status string

// Define the allowed status values as constants
const (
	New       Status = "NEW"
	Started   Status = "STARTED"
	Pending   Status = "PENDING"
	Completed Status = "COMPLETED"
)
