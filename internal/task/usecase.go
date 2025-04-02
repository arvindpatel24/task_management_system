package task

type UseCase interface {
	CreateTask(task Task) (Task, error)
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
