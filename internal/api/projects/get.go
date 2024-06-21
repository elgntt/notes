package projects

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	apimodels "task-manager/internal/api/builders/models"
	"task-manager/internal/api/builders/response"
	serv "task-manager/internal/service"
)

func (api *API) Get(c *gin.Context) {
	projectID, err := strconv.Atoi(c.Param("projectId"))
	if err != nil {
		response.WithParameterInvalidErr(c, InvalidProjectIDErrMsg)
		return
	}

	project, err := api.service.GetByID(c, projectID)
	if err != nil {
		if errors.Is(err, serv.ErrProjectNotFound) {
			response.WithNotFoundErr(c, err.Error())
			return
		}

		api.logger.Err(err)
		response.WithInternalServerError(c)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"project": apimodels.ProjectResp(project),
	})
}
