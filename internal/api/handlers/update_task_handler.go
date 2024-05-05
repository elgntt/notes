package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/elgntt/notes/internal/model/dto"
)

func (h *Handler) UpdateTask(c *gin.Context) {
	var req dto.UpdateTaskReq
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypeBind)

		return
	}

	taskID, err := parseTaskID(c.Param("taskId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": dto.Error{
				Message: err.Error(),
			},
		})

		return
	}

	err = h.taskService.Update(c, buildUpdateTask(taskID, req))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": dto.Error{
				Message: "Something went wrong, try again later.",
			},
		})
		h.logger.Err(err)

		return
	}

	c.Status(http.StatusOK)
}

func buildUpdateTask(taskID int, req dto.UpdateTaskReq) dto.UpdateTask {
	var dueDate *time.Time
	if req.DueDate != nil {
		dueDate = &req.DueDate.Time
	}

	return dto.UpdateTask{
		ID:          taskID,
		Title:       req.Title,
		Description: req.Description,
		DueDate:     dueDate,
		Status:      req.Status,
	}
}
