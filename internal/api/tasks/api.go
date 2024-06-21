package tasks

import (
	"context"

	"task-manager/internal/model/domain"
	"task-manager/internal/model/dto"
	validationpkg "task-manager/internal/pkg/validator"
)

const (
	InvalidTaskIDErrMsg = "невалидный параметр taskId"
)

type logger interface {
	Info(text ...any)
	Warn(text ...any)
	Err(text ...any)
}

type service interface {
	Create(ctx context.Context, task dto.NewTask) (domain.Task, error)
	Update(ctx context.Context, updateTask dto.UpdateTask) error
	Delete(ctx context.Context, taskId int) error
	Get(ctx context.Context, categoryID int) ([]domain.Task, error)
	GetByID(ctx context.Context, taskId int) (domain.Task, error)
}

type API struct {
	logger    logger
	service   service
	validator *validationpkg.CustomValidator
}

func NewAPI(logger logger, service service, validator *validationpkg.CustomValidator) *API {
	return &API{
		logger:    logger,
		service:   service,
		validator: validator,
	}
}
