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

func (api *API) Update(c *gin.Context) {
	taskID, err := strconv.Atoi(c.Param("taskId"))
	if err != nil {
		response.WithParameterInvalidErr(c, InvalidTaskIDErrMsg)
		return
	}

	var req apimodels.UpdateTaskReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.WithJSONError(c, err)
		return
	}

	if err := api.validator.Struct(req); err != nil {
		response.WithValidationError(c, err)
		return
	}

	err = api.service.Update(c, dto.UpdateTask{
		ID:          taskID,
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
	})
	if err != nil {
		if errors.Is(err, serv.ErrTaskNotExists) {
			response.WithNotFoundErr(c, err.Error())
			return
		}

		response.WithInternalServerError(c)
		api.logger.Err(err)

		return
	}

	c.Status(http.StatusOK)
}
