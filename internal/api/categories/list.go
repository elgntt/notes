package categories

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/samber/lo"

	apimodels "task-manager/internal/api/builders/models"
	"task-manager/internal/api/builders/response"
	"task-manager/internal/model/domain"
)

func (api *API) List(c *gin.Context) {
	projectID, err := strconv.Atoi(c.Param("projectId"))
	if err != nil {
		response.WithParameterInvalidErr(c, InvalidProjectIDErrMsg)
		return
	}

	categories, err := api.service.FindByProjectID(c, projectID)
	if err != nil {
		response.WithInternalServerError(c)
		api.logger.Err(err)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"categories": lo.Map(categories, func(item domain.Category, _ int) apimodels.CategoryResp {
			return apimodels.CategoryResp(item)
		}),
	})
}
