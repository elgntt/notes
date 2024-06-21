package tasks

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	apimodels "task-manager/internal/api/builders/models"
	"task-manager/internal/api/builders/response"
	"task-manager/internal/model/dto"
	serv "task-manager/internal/service"
)

func (api *API) Get(c *gin.Context) {
	taskId, err := strconv.Atoi(c.Param("taskId"))
	if err != nil {
		response.WithParameterInvalidErr(c, InvalidTaskIDErrMsg)
		return
	}

	task, err := api.service.GetByID(c, taskId)
	if err != nil {
		if errors.Is(err, serv.ErrTaskNotExists) {
			response.WithNotFoundErr(c, err.Error())
			return
		}

		response.WithInternalServerError(c)
		api.logger.Err(err)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"task": apimodels.TaskResp{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Status:      dto.TaskStatus(task.Status),
			CategoryID:  task.CategoryID,
			ProjectID:   task.ProjectID,
		},
	})
}
