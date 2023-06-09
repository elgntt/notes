package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

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
