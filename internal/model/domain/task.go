package domain

import (
	"task-manager/internal/model/dto"
)

type Task struct {
	ID          int
	Title       string
	Description *string
	Status      string
	CategoryID  int
	ProjectID   int
}

func (t *Task) Update(task dto.UpdateTask) {
	if task.Title != nil {
		t.Title = *task.Title
	}
	if task.Description != nil {
		t.Description = task.Description
	}
	if task.Status != nil {
		t.Status = string(*task.Status)
	}
}
