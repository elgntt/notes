package models

import (
	"task-manager/internal/model/dto"
)

type CreateTaskReq struct {
	Title       string         `json:"title" validate:"required"`
	Description *string        `json:"description"`
	Status      dto.TaskStatus `json:"status" validate:"required,status"`
	CategoryID  int            `json:"categoryId" validate:"required"`
	ProjectID   int            `json:"projectId" validate:"required"`
}

type UpdateTaskReq struct {
	Title       *string         `json:"title"`
	Description *string         `json:"text"`
	Status      *dto.TaskStatus `json:"status" validate:"status"`
}

type TaskResp struct {
	ID          int            `json:"id"`
	Title       string         `json:"title"`
	Description *string        `json:"description,omitempty"`
	Status      dto.TaskStatus `json:"status"`
	CategoryID  int            `json:"categoryId"`
	ProjectID   int            `json:"projectId"`
}
