package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/elgntt/notes/internal/model"
	businessErr "github.com/elgntt/notes/internal/pkg/errors"
	response "github.com/elgntt/notes/internal/pkg/http"
)

type UpdateNoteData struct {
	Id   int    `json:"id"`
	Text string `json:"text"`
}

func (h *Handler) UpdateNote(c *gin.Context) {
	data := UpdateNoteData{}

	if err := c.BindJSON(&data); err != nil {
		response.WriteErrorResponse(c, h.logger, err)
		return
	}

	err := checkUpdateRequestData(data)
	if err != nil {
		response.WriteErrorResponse(c, h.logger, err)
		return
	}

	note := model.Note{
		Id:   data.Id,
		Text: data.Text,
	}

	ctx := context.Background()

	err = h.noteService.UpdateNote(ctx, note)
	if err != nil {
		response.WriteErrorResponse(c, h.logger, err)
		return
	}

	c.Status(http.StatusOK)
}

func checkUpdateRequestData(data UpdateNoteData) error {
	if data.Id == 0 {
		return businessErr.NewBusinessError(noteIsEmptyErr)
	}
	if data.Text == "" {
		return businessErr.NewBusinessError(noteIsEmptyErr)
	}

	return nil
}
