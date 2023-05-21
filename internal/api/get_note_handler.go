package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	response "github.com/elgntt/notes/internal/pkg/http"
)

func (h *Handler) GetNote(c *gin.Context) {
	ctx := context.Background()

	noteId, err := parseNoteId(c.Query("noteId"))
	if err != nil {
		response.WriteErrorResponse(c, h.logger, err)
		return
	}

	note, err := h.noteService.GetNote(ctx, noteId)
	if err != nil {
		response.WriteErrorResponse(c, h.logger, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"note": note,
	})
}
