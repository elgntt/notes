package handlers

import (
	"context"

	"github.com/elgntt/notes/internal/model/domain"
	"github.com/elgntt/notes/internal/model/dto"
)

type taskService interface {
	Create(ctx context.Context, newTask dto.NewTask) error
	Update(ctx context.Context, updateTask dto.UpdateTask) error
	Delete(ctx context.Context, taskId int) error
	GetAll(ctx context.Context) ([]domain.Task, error)
	GetByID(ctx context.Context, taskId int) (domain.Task, error)
}

type logger interface {
	Info(text ...any)
	Warn(text ...any)
	Err(text ...any)
}
