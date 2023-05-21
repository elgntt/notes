package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	response "github.com/elgntt/notes/internal/pkg/http"
)

func (h *Handler) GetAllNotes(c *gin.Context) {
	ctx := context.Background()

	notes, err := h.noteService.GetAllNotes(ctx)
	if err != nil {
		response.WriteErrorResponse(c, h.logger, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"notes": notes,
	})
}
