package dto

import (
	"time"
)

type CreateTaskReq struct {
	Title       string     `json:"title" binding:"required"`
	Description *string    `json:"description"`
	DueDate     *Date      `json:"dueDate" binding:"required,date"`
	Status      TaskStatus `json:"status" binding:"required,status"`
}

type UpdateTaskReq struct {
	Title       *string     `json:"title"`
	Description *string     `json:"text"`
	DueDate     *Date       `json:"dueDate" binding:"date"`
	Status      *TaskStatus `json:"status" binding:"status"`
}

type TaskResp struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description *string    `json:"description,omitempty"`
	DueDate     *Date      `json:"dueDate,omitempty"`
	Status      TaskStatus `json:"status"`
}

type UpdateTask struct {
	ID          int
	Title       *string
	Description *string
	DueDate     *time.Time
	Status      *TaskStatus
}

type NewTask struct {
	Title       string
	Description *string
	DueDate     *Date
	Status      TaskStatus
}
