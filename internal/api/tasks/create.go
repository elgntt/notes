package tasks

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	apimodels "task-manager/internal/api/builders/models"
	"task-manager/internal/api/builders/response"
	"task-manager/internal/model/dto"
	serv "task-manager/internal/service"
)

func (api *API) Create(c *gin.Context) {
	var req apimodels.CreateTaskReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.WithJSONError(c, err)
		return
	}

	if err := api.validator.Struct(req); err != nil {
		response.WithValidationError(c, err)
		return
	}

	task, err := api.service.Create(c, dto.NewTask{
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
		CategoryID:  req.CategoryID,
		ProjectID:   req.ProjectID,
	})
	if err != nil {
		if errors.Is(err, serv.ErrProjectNotFound) {
			response.WithNotFoundErr(c, err.Error())
			return
		}

		if errors.Is(err, serv.ErrCategoryNotFound) {
			response.WithNotFoundErr(c, err.Error())
			return
		}

		response.WithInternalServerError(c)
		api.logger.Err(err)

		return
	}

	c.JSON(http.StatusCreated, gin.H{
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
