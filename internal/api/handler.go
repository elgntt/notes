package api

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/elgntt/notes/internal/model"
)

type NoteService interface {
	CreateNote(ctx context.Context, note model.Note) error
	UpdateNote(ctx context.Context, note model.Note) error
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
	noteIsEmptyErr     = "Ошибка. Заметка не может быть пустой!"
	noteIsNotExistsErr = "Ошибка. Нет такой заметки!"
)

func New(noteService NoteService, logger Logger) *gin.Engine {
	h := Handler{
		noteService: noteService,
		logger:      logger,
	}

	r := gin.New()

	api := r.Group("/api")

	api.POST("/note/create", h.CreateNote)
	api.POST("/note/update", h.UpdateNote)

	return r
}
