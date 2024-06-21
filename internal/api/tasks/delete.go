package tasks

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"task-manager/internal/api/builders/response"
)

func (api *API) Delete(c *gin.Context) {
	taskId, err := strconv.Atoi(c.Param("taskId"))
	if err != nil {
		response.WithParameterInvalidErr(c, InvalidTaskIDErrMsg)
		return
	}

	err = api.service.Delete(c, taskId)
	if err != nil {
		response.WithInternalServerError(c)
		api.logger.Err(err)

		return
	}

	c.Status(http.StatusOK)
}
