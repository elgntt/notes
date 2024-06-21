package tasks

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/samber/lo"

	apimodels "task-manager/internal/api/builders/models"
	"task-manager/internal/api/builders/response"
	"task-manager/internal/model/domain"
	"task-manager/internal/model/dto"
)

func (api *API) List(c *gin.Context) {
	categoryID, err := strconv.Atoi(c.Param("categoryId"))
	if err != nil {
		response.WithParameterInvalidErr(c, InvalidTaskIDErrMsg)
		return
	}

	tasks, err := api.service.Get(c, categoryID)
	if err != nil {
		response.WithInternalServerError(c)
		api.logger.Err(err)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"task": lo.Map(tasks, func(item domain.Task, _ int) apimodels.TaskResp {
			return apimodels.TaskResp{
				ID:          item.ID,
				Title:       item.Title,
				Description: item.Description,
				Status:      dto.TaskStatus(item.Status),
				CategoryID:  item.CategoryID,
				ProjectID:   item.ProjectID,
			}
		}),
	})
}
