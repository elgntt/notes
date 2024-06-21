package categories

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
	categoryID, err := strconv.Atoi(c.Param("categoryId"))
	if err != nil {
		response.WithParameterInvalidErr(c, InvalidCategoryIDErrMsg)
		return
	}

	category, err := api.service.GetByID(c, categoryID)
	if err != nil {
		if errors.Is(err, serv.ErrCategoryNotFound) {
			response.WithNotFoundErr(c, err.Error())
			return
		}

		response.WithInternalServerError(c)
		api.logger.Err(err)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"category": apimodels.CategoryResp(category),
	})
}
