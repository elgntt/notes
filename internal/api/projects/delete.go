package projects

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"task-manager/internal/api/builders/response"
)

func (api *API) Delete(c *gin.Context) {
	projectID, err := strconv.Atoi(c.Param("projectId"))
	if err != nil {
		response.WithParameterInvalidErr(c, InvalidProjectIDErrMsg)
		return
	}

	if err := api.service.Delete(c, projectID); err != nil {
		response.WithInternalServerError(c)
		api.logger.Err(err)

		return
	}

	c.Status(http.StatusOK)
}
