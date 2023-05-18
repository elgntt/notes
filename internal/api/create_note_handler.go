package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/elgntt/notes/internal/model"
	businessErr "github.com/elgntt/notes/internal/pkg/errors"
	response "github.com/elgntt/notes/internal/pkg/http"
)

type NoteData struct {
	Text string `json:"text"`
}

func (h *Handler) CreateNote(c *gin.Context) {
	data := NoteData{}

	if err := c.BindJSON(&data); err != nil {
		response.WriteErrorResponse(c, h.logger, err)
		return
	}

	err := checkRequest(data)
	if err != nil {
		response.WriteErrorResponse(c, h.logger, err)
		return
	}

	note := model.Note{
		Text: data.Text,
	}

	ctx := context.Background()

	err = h.noteService.CreateNote(ctx, note)
	if err != nil {
		response.WriteErrorResponse(c, h.logger, err)
		return
	}

	c.Status(http.StatusOK)
}

func checkRequest(data NoteData) error {
	if data.Text == "" {
		return businessErr.NewBusinessError("Ошибка. Заметка пустая!")
	}

	return nil
}
