package projects

import (
	"context"

	"task-manager/internal/model/domain"
	"task-manager/internal/model/dto"
	"task-manager/internal/service"
)

type projectsRepo interface {
	Create(ctx context.Context, project dto.Project) (domain.Project, error)
	GetByID(ctx context.Context, id int) (*domain.Project, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, project domain.Project) (domain.Project, error)
}

type Service struct {
	projectRepo projectsRepo
}

func New(projectRepo projectsRepo) *Service {
	return &Service{
		projectRepo: projectRepo,
	}
}

func (s *Service) Create(ctx context.Context, project dto.Project) (domain.Project, error) {
	return s.projectRepo.Create(ctx, project)
}

func (s *Service) GetByID(ctx context.Context, id int) (domain.Project, error) {
	project, err := s.projectRepo.GetByID(ctx, id)
	if err != nil {
		return domain.Project{}, err
	}
	if project == nil {
		return domain.Project{}, service.ErrProjectNotFound
	}

	return *project, nil
}

func (s *Service) Delete(ctx context.Context, id int) error {
	return s.projectRepo.Delete(ctx, id)
}

func (s *Service) Update(ctx context.Context, proj dto.Project, projectID int) (domain.Project, error) {
	project, err := s.projectRepo.GetByID(ctx, projectID)
	if err != nil {
		return domain.Project{}, err
	}
	if project == nil {
		return domain.Project{}, service.ErrProjectNotFound
	}

	project.Update(proj)

	return s.projectRepo.Update(ctx, *project)
}
