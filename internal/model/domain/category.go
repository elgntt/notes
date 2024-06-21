package domain

import "task-manager/internal/model/dto"

type Category struct {
	ID        int
	Title     string
	ProjectID int
}

func (c *Category) Update(draft dto.UpdateCategory) {
	if draft.Title != "" {
		c.Title = draft.Title
	}
}
