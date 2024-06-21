package projects

import (
	"net/http"

	"github.com/gin-gonic/gin"

	apimodels "task-manager/internal/api/builders/models"
	"task-manager/internal/api/builders/response"
	"task-manager/internal/model/dto"
)

func (api *API) Create(c *gin.Context) {
	var req apimodels.NewProjectReq

	if err := c.ShouldBindJSON(&req); err != nil {
		response.WithJSONError(c, err)
		return
	}

	if err := api.validator.Struct(req); err != nil {
		response.WithValidationError(c, err)
		return
	}

	project, err := api.service.Create(c, dto.Project{
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		response.WithJSONError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"project": apimodels.ProjectResp(project),
	})
}
