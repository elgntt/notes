package builders

import (
	"github.com/elgntt/notes/internal/model/domain"
	"github.com/elgntt/notes/internal/model/dto"
)

func BuildTask(task domain.Task) dto.TaskResp {
	var dueDate *dto.Date
	if task.DueDate != nil {
		dueDate = &dto.Date{
			Time: *task.DueDate,
		}
	}

	return dto.TaskResp{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		DueDate:     dueDate,
		Status:      dto.TaskStatus(task.Status),
	}
}
