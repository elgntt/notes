package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/elgntt/notes/internal/model/dto"
)

func (h *Handler) DeleteTask(c *gin.Context) {
	taskId, err := parseTaskID(c.Param("taskId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": dto.Error{
				Message: err.Error(),
			},
		})

		return
	}

	err = h.taskService.Delete(c, taskId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong, try again later.",
		})
		h.logger.Err(err)

		return
	}

	c.Status(http.StatusOK)
}
