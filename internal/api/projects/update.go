package projects

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
	projectID, err := strconv.Atoi(c.Param("projectId"))
	if err != nil {
		response.WithParameterInvalidErr(c, InvalidProjectIDErrMsg)
		return
	}

	var req apimodels.UpdateProjectReq

	if err := c.ShouldBindJSON(&req); err != nil {
		response.WithJSONError(c, err)
		return
	}

	if err := api.validator.Struct(req); err != nil {
		response.WithValidationError(c, err)
		return
	}

	project, err := api.service.Update(c, dto.Project(req), projectID)
	if err != nil {
		if errors.Is(err, serv.ErrProjectNotFound) {
			response.WithNotFoundErr(c, err.Error())
			return
		}

		response.WithInternalServerError(c)
		api.logger.Err(err)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"project": apimodels.ProjectResp(project),
	})
}
