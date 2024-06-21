package categories

import (
	"context"

	"task-manager/internal/model/domain"
	"task-manager/internal/model/dto"
	validationpkg "task-manager/internal/pkg/validator"
)

type logger interface {
	Info(text ...any)
	Warn(text ...any)
	Err(text ...any)
}

type service interface {
	Create(ctx context.Context, draft dto.NewCategory) (domain.Category, error)
	GetByID(ctx context.Context, id int) (domain.Category, error)
	FindByProjectID(ctx context.Context, projectId int) ([]domain.Category, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, draft dto.UpdateCategory) (domain.Category, error)
}

const (
	InvalidCategoryIDErrMsg = "невалидный параметр categoryId"
	InvalidProjectIDErrMsg  = "невалидный параметр projectId "
)

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
