package service

import (
	"context"

	"github.com/elgntt/notes/internal/model"
)

type Repository interface {
	CreateNote(ctx context.Context, note model.Note) error
	UpdateNote(ctx context.Context, note model.Note) error
	DeleteNote(ctx context.Context, noteId int) error
	GetAllNotes(ctx context.Context) ([]model.NoteInfo, error)
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

func (s *Service) UpdateNote(ctx context.Context, note model.Note) error {
	return s.repo.UpdateNote(ctx, note)
}

func (s *Service) DeleteNote(ctx context.Context, noteId int) error {
	return s.repo.DeleteNote(ctx, noteId)
}

func (s *Service) GetAllNotes(ctx context.Context) ([]model.NoteInfo, error) {
	return s.repo.GetAllNotes(ctx)
}
