package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/elgntt/notes/internal/api/handlers/builders"
	"github.com/elgntt/notes/internal/model/dto"
	"github.com/elgntt/notes/internal/service"
)

func (h *Handler) GetTask(c *gin.Context) {
	taskId, err := parseTaskID(c.Param("taskId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": dto.Error{
				Message: err.Error(),
			},
		})

		return
	}

	task, err := h.taskService.GetByID(c, taskId)
	if err != nil {
		if errors.Is(err, service.ErrTaskNotExists) {
			c.AbortWithStatus(http.StatusNotFound)

			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": dto.Error{
				Message: "Something went wrong, try again later.",
			},
		})
		h.logger.Err(err)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"task": builders.BuildTask(task),
	})
}
