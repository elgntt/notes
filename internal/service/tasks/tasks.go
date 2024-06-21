package tasks

import (
	"context"

	"task-manager/internal/model/domain"
	"task-manager/internal/model/dto"
	"task-manager/internal/service"
)

type tasksRepo interface {
	Create(ctx context.Context, task dto.NewTask) (int, error)
	Update(ctx context.Context, note domain.Task) error
	Delete(ctx context.Context, noteId int) error
	Get(ctx context.Context, categoryID int) ([]domain.Task, error)
	GetByID(ctx context.Context, taskId int) (*domain.Task, error)
}

type categoriesRepo interface {
	GetByID(ctx context.Context, id int) (*domain.Category, error)
}

type projectRepo interface {
	GetByID(ctx context.Context, id int) (*domain.Project, error)
}

type Service struct {
	tasksRepo      tasksRepo
	categoriesRepo categoriesRepo
	projectRepo    projectRepo
}

func New(tasksRepo tasksRepo, categoriesRepo categoriesRepo, projectRepo projectRepo) *Service {
	return &Service{
		tasksRepo:      tasksRepo,
		categoriesRepo: categoriesRepo,
		projectRepo:    projectRepo,
	}
}

func (s *Service) Create(ctx context.Context, task dto.NewTask) (domain.Task, error) {
	project, err := s.projectRepo.GetByID(ctx, task.ProjectID)
	if err != nil {
		return domain.Task{}, err
	}
	if project == nil {
		return domain.Task{}, service.ErrProjectNotFound
	}

	category, err := s.categoriesRepo.GetByID(ctx, task.CategoryID)
	if err != nil {
		return domain.Task{}, err
	}
	if category == nil {
		return domain.Task{}, service.ErrCategoryNotFound
	}

	id, err := s.tasksRepo.Create(ctx, task)
	if err != nil {
		return domain.Task{}, err
	}

	return domain.Task{
		ID:          id,
		Title:       task.Title,
		Description: task.Description,
		Status:      string(task.Status),
		CategoryID:  task.CategoryID,
		ProjectID:   task.ProjectID,
	}, nil
}

func (s *Service) Update(ctx context.Context, updateTask dto.UpdateTask) error {
	task, err := s.tasksRepo.GetByID(ctx, updateTask.ID)
	if err != nil {
		return err
	}
	if task == nil {
		return service.ErrTaskNotExists
	}

	task.Update(updateTask)

	return s.tasksRepo.Update(ctx, *task)
}

func (s *Service) Delete(ctx context.Context, taskID int) error {
	return s.tasksRepo.Delete(ctx, taskID)
}

func (s *Service) Get(ctx context.Context, categoryID int) ([]domain.Task, error) {
	return s.tasksRepo.Get(ctx, categoryID)
}

func (s *Service) GetByID(ctx context.Context, taskID int) (domain.Task, error) {
	task, err := s.tasksRepo.GetByID(ctx, taskID)
	if err != nil {
		return domain.Task{}, err
	}
	if task == nil {
		return domain.Task{}, service.ErrTaskNotExists
	}

	return *task, nil
}
