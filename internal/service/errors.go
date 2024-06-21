package service

import "errors"

var (
	ErrTaskNotExists    = errors.New("задача не найдена")
	ErrProjectNotFound  = errors.New("project not found")
	ErrCategoryNotFound = errors.New("category not found")
)
