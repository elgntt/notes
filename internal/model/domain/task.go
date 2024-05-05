package domain

import (
	"time"

	"github.com/elgntt/notes/internal/model/dto"
)

type Task struct {
	ID          int
	Title       string
	Description *string
	DueDate     *time.Time
	Status      string
}

func (t *Task) Update(task dto.UpdateTask) {
	if task.Title != nil {
		t.Title = *task.Title
	}
	if task.Description != nil {
		t.Description = task.Description
	}
	if task.DueDate != nil {
		t.DueDate = task.DueDate
	}
	if task.Status != nil {
		t.Status = string(*task.Status)
	}
}
