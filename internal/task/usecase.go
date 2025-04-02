package task

type UseCase interface {
	CreateTask(task Task) (Task, error)
	GetTasks() ([]Task, error)
	GetTaskByID(id int64) (Task, error)
	UpdateTask(id int64, task Task) (Task, error)
	DeleteTask(id int64) error
}

type taskUseCase struct {
	repo Repository
}

func NewTaskUseCase(r Repository) UseCase {
	return &taskUseCase{repo: r}
}

// Create task
func (u *taskUseCase) CreateTask(task Task) (Task, error) {
	return u.repo.Create(task)
}

// List all tasks
func (u *taskUseCase) GetTasks() ([]Task, error) {
	return u.repo.GetAll()
}

// Retrieves a task by its ID
func (u *taskUseCase) GetTaskByID(id int64) (Task, error) {
	return u.repo.GetByID(id)
}

// Update an existing task
func (u *taskUseCase) UpdateTask(id int64, task Task) (Task, error) {
	return u.repo.Update(id, task)
}

// Deletes a task by its ID
func (u *taskUseCase) DeleteTask(id int64) error {
	return u.repo.Delete(id)
}
