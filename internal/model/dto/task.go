package dto

type UpdateTask struct {
	ID          int
	Title       *string
	Description *string
	Status      *TaskStatus
}

type NewTask struct {
	Title       string
	Description *string
	Status      TaskStatus
	CategoryID  int
	ProjectID   int
}
