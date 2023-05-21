package api

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/elgntt/notes/internal/model"
	businessErr "github.com/elgntt/notes/internal/pkg/errors"
)

type NoteService interface {
	CreateNote(ctx context.Context, note model.Note) error
	UpdateNote(ctx context.Context, note model.Note) error
	DeleteNote(ctx context.Context, noteId int) error
	GetAllNotes(ctx context.Context) ([]model.NoteInfo, error)
	GetNote(ctx context.Context, noteId int) (model.NoteInfo, error)
}

type Logger interface {
	Info(text ...any)
	Warn(text ...any)
	Err(text ...any)
}

type Handler struct {
	noteService NoteService
	logger      Logger
}

const (
	noteIsEmptyErr   = "Заметка не может быть пустой!"
	invalidNoteIdErr = "Невалидный параметр noteId"
	noteIdNotExists  = "Не указан noteId"
)

func New(noteService NoteService, logger Logger) *gin.Engine {
	h := Handler{
		noteService: noteService,
		logger:      logger,
	}

	r := gin.New()

	api := r.Group("/api/note")

	api.POST("/create", h.CreateNote)
	api.POST("/update", h.UpdateNote)
	api.DELETE("/delete", h.DeleteNote)
	api.GET("/getAll", h.GetAllNotes)
	api.GET("/get", h.GetNote)

	return r
}

func parseNoteId(noteIdQuery string) (int, error) {
	if noteIdQuery == "" {
		return 0, businessErr.NewBusinessError(noteIdNotExists)
	}

	noteId, err := strconv.Atoi(noteIdQuery)
	if err != nil {
		return 0, businessErr.NewBusinessError(invalidNoteIdErr)
	}

	return noteId, nil
}
