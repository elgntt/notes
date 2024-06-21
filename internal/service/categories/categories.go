package categories

import (
	"context"

	"task-manager/internal/model/domain"
	"task-manager/internal/model/dto"
	"task-manager/internal/service"
)

type categoryRepo interface {
	Create(ctx context.Context, draft dto.NewCategory) (domain.Category, error)
	GetByID(ctx context.Context, id int) (*domain.Category, error)
	FindByProjectID(ctx context.Context, projectId int) ([]domain.Category, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, draft domain.Category) (domain.Category, error)
}

type projectRepo interface {
	GetByID(ctx context.Context, id int) (*domain.Project, error)
}

type Service struct {
	categoryRepo categoryRepo
	projectRepo  projectRepo
}

func New(categoryRepo categoryRepo, projectRepo projectRepo) *Service {
	return &Service{
		categoryRepo: categoryRepo,
		projectRepo:  projectRepo,
	}
}

func (s *Service) Create(ctx context.Context, draft dto.NewCategory) (domain.Category, error) {
	project, err := s.projectRepo.GetByID(ctx, draft.ProjectID)
	if err != nil {
		return domain.Category{}, err
	}
	if project == nil {
		return domain.Category{}, service.ErrProjectNotFound
	}

	return s.categoryRepo.Create(ctx, draft)
}

func (s *Service) GetByID(ctx context.Context, id int) (domain.Category, error) {
	category, err := s.categoryRepo.GetByID(ctx, id)
	if err != nil {
		return domain.Category{}, err
	}
	if category == nil {
		return domain.Category{}, service.ErrCategoryNotFound
	}

	return *category, nil
}

func (s *Service) FindByProjectID(ctx context.Context, projectId int) ([]domain.Category, error) {
	return s.categoryRepo.FindByProjectID(ctx, projectId)
}

func (s *Service) Delete(ctx context.Context, id int) error {
	return s.categoryRepo.Delete(ctx, id)
}

func (s *Service) Update(ctx context.Context, draft dto.UpdateCategory) (domain.Category, error) {
	category, err := s.categoryRepo.GetByID(ctx, draft.ID)
	if err != nil {
		return domain.Category{}, err
	}
	if category == nil {
		return domain.Category{}, service.ErrCategoryNotFound
	}

	category.Update(draft)

	return s.categoryRepo.Update(ctx, *category)
}
