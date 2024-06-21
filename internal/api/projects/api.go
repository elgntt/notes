package projects

import (
	"context"

	"task-manager/internal/model/domain"
	"task-manager/internal/model/dto"
	validatorpkg "task-manager/internal/pkg/validator"
)

const (
	InvalidProjectIDErrMsg = "invalid project ID parameter"
)

type service interface {
	Create(ctx context.Context, project dto.Project) (domain.Project, error)
	GetByID(ctx context.Context, id int) (domain.Project, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, proj dto.Project, projectID int) (domain.Project, error)
}

type logger interface {
	Info(text ...any)
	Warn(text ...any)
	Err(text ...any)
}

type API struct {
	logger    logger
	service   service
	validator *validatorpkg.CustomValidator
}

func NewAPI(logger logger, service service, validator *validatorpkg.CustomValidator) *API {
	return &API{
		logger:    logger,
		service:   service,
		validator: validator,
	}
}
