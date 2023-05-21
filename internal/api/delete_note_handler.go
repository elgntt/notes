package api

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	businessErr "github.com/elgntt/notes/internal/pkg/errors"
	response "github.com/elgntt/notes/internal/pkg/http"
)

func (h *Handler) DeleteNote(c *gin.Context) {
	noteId, err := parseNoteId(c.Query("noteId"))
	if err != nil {
		response.WriteErrorResponse(c, h.logger, err)
		return
	}

	ctx := context.Background()

	err = h.noteService.DeleteNote(ctx, noteId)
	if err != nil {
		response.WriteErrorResponse(c, h.logger, err)
		return
	}

	c.Status(http.StatusOK)
}

func parseNoteId(noteIdQuery string) (int, error) {
	if noteIdQuery == "" {
		return 0, businessErr.NewBusinessError("Ошибка. Параметр noteId пустой")
	}

	noteId, err := strconv.Atoi(noteIdQuery)
	if err != nil {
		return 0, businessErr.NewBusinessError("Ошибка. Невалидный параметр noteId")
	}

	return noteId, nil
}
