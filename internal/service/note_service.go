package service

import (
	"context"

	"github.com/elgntt/notes/internal/model"
)

type Repository interface {
	CreateNote(ctx context.Context, note model.Note) error
}

type Service struct {
	repo Repository
}

func New(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) CreateNote(ctx context.Context, note model.Note) error {
	return s.repo.CreateNote(ctx, note)
}
