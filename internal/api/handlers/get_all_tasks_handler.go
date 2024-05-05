package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samber/lo"

	"github.com/elgntt/notes/internal/api/handlers/builders"
	"github.com/elgntt/notes/internal/model/domain"
	"github.com/elgntt/notes/internal/model/dto"
)

func (h *Handler) GetAllTasks(c *gin.Context) {
	tasks, err := h.taskService.GetAll(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong, try again later.",
		})
		h.logger.Err(err)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tasks": lo.Map(tasks, func(item domain.Task, _ int) dto.TaskResp {
			return builders.BuildTask(item)
		}),
	})
}
