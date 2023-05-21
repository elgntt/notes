package note

import (
	"context"

	"github.com/elgntt/notes/internal/model"
	businessErr "github.com/elgntt/notes/internal/pkg/errors"
)

type Repository interface {
	CreateNote(ctx context.Context, note model.Note) error
	UpdateNote(ctx context.Context, note model.Note) error
	DeleteNote(ctx context.Context, noteId int) error
	GetAllNotes(ctx context.Context) ([]model.NoteInfo, error)
	GetNote(ctx context.Context, noteId int) (*model.NoteInfo, error)
}

type Service struct {
	repo Repository
}

const (
	noteNotExistsErr = "Заметка не найдена"
)

func New(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) CreateNote(ctx context.Context, note model.Note) error {
	return s.repo.CreateNote(ctx, note)
}

func (s *Service) UpdateNote(ctx context.Context, data model.Note) error {
	note, err := s.repo.GetNote(ctx, data.Id)
	if err != nil {
		return err
	}

	if note == nil {
		return businessErr.NewBusinessError(noteNotExistsErr)
	}

	return s.repo.UpdateNote(ctx, data)
}

func (s *Service) DeleteNote(ctx context.Context, noteId int) error {
	return s.repo.DeleteNote(ctx, noteId)
}

func (s *Service) GetAllNotes(ctx context.Context) ([]model.NoteInfo, error) {
	return s.repo.GetAllNotes(ctx)
}

func (s *Service) GetNote(ctx context.Context, noteId int) (model.NoteInfo, error) {
	note, err := s.repo.GetNote(ctx, noteId)
	if err != nil {
		return model.NoteInfo{}, err
	}
	if note == nil {
		return model.NoteInfo{}, businessErr.NewBusinessError(noteNotExistsErr)
	}

	return *note, nil
}
