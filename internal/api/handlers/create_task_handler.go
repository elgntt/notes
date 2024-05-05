package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/elgntt/notes/internal/model/dto"
)

func (h *Handler) CreateTask(c *gin.Context) {
	var req dto.CreateTaskReq
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypeBind)

		return
	}

	if err := h.taskService.Create(c, dto.NewTask{
		Title:       req.Title,
		Description: req.Description,
		DueDate:     req.DueDate,
		Status:      req.Status,
	}); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong, try again later.",
		})
		h.logger.Err(err)

		return
	}

	c.Status(http.StatusOK)
}
