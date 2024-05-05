package service

import (
	"context"
	"errors"

	"github.com/elgntt/notes/internal/model/domain"
	"github.com/elgntt/notes/internal/model/dto"
)

type tasksRepo interface {
	Create(ctx context.Context, task dto.NewTask) error
	Update(ctx context.Context, note domain.Task) error
	Delete(ctx context.Context, noteId int) error
	GetAll(ctx context.Context) ([]domain.Task, error)
	GetByID(ctx context.Context, taskId int) (*domain.Task, error)
}

type Service struct {
	tasks tasksRepo
}

var (
	ErrTaskNotExists = errors.New("задача не найдена")
)

func New(r tasksRepo) *Service {
	return &Service{
		tasks: r,
	}
}

func (s *Service) Create(ctx context.Context, task dto.NewTask) error {
	return s.tasks.Create(ctx, task)
}

func (s *Service) Update(ctx context.Context, updateTask dto.UpdateTask) error {
	task, err := s.tasks.GetByID(ctx, updateTask.ID)
	if err != nil {
		return err
	}
	if task == nil {
		return ErrTaskNotExists
	}

	task.Update(updateTask)

	return s.tasks.Update(ctx, *task)
}

func (s *Service) Delete(ctx context.Context, taskID int) error {
	return s.tasks.Delete(ctx, taskID)
}

func (s *Service) GetAll(ctx context.Context) ([]domain.Task, error) {
	return s.tasks.GetAll(ctx)
}

func (s *Service) GetByID(ctx context.Context, taskID int) (domain.Task, error) {
	task, err := s.tasks.GetByID(ctx, taskID)
	if err != nil {
		return domain.Task{}, err
	}
	if task == nil {
		return domain.Task{}, ErrTaskNotExists
	}

	return *task, nil
}
