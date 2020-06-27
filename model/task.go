package model

// Status da tarefa
type Status string

// Enum status
const (
	StatusNotInit  Status = "NOT_INIT"
	StatusProgess  Status = "PROGRESS"
	StatusFinished Status = "FINISHED"
	StatusPending  Status = "PENDING"
)

// Task ...
type Task struct {
	ID          int
	Name        string
	Description string
	Status      Status
}

// TaskDTO ...
type TaskDTO struct {
	Name        string
	Description string
	Status      Status
}
