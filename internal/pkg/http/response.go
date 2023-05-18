package http

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	businessErr "github.com/elgntt/notes/internal/pkg/errors"
)

type ErrorResponse struct {
	Error `json:"error"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Logger interface {
	Info(text ...any)
	Warn(text ...any)
	Err(text ...any)
}

func WriteErrorResponse(c *gin.Context, logger Logger, err error) {
	var bErr businessErr.BusinessError

	if errors.As(err, &bErr) {
		errorResponse := ErrorResponse{
			Error: Error{
				Code:    bErr.Code(),
				Message: bErr.Error(),
			},
		}

		c.JSON(http.StatusBadRequest, errorResponse)
	} else {
		errorResponse := ErrorResponse{
			Error: Error{
				Code:    "InternalServerError",
				Message: "Что-то пошло не так, попробуйте еще раз",
			},
		}
		logger.Err(err)

		c.JSON(http.StatusInternalServerError, errorResponse)
	}
}
