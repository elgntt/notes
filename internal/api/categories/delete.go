package categories

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"task-manager/internal/api/builders/response"
)

func (api *API) Delete(c *gin.Context) {
	categoryID, err := strconv.Atoi(c.Param("categoryId"))
	if err != nil {
		response.WithParameterInvalidErr(c, InvalidProjectIDErrMsg)
		return
	}

	if err := api.service.Delete(c, categoryID); err != nil {
		response.WithInternalServerError(c)
		api.logger.Err(err)

		return
	}

	c.Status(http.StatusOK)
}
