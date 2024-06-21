package domain

import "task-manager/internal/model/dto"

type Project struct {
	ID          int
	Name        string
	Description string
}

func (p *Project) Update(proj dto.Project) {
	if proj.Name != "" {
		p.Name = proj.Name
	}
	if proj.Description != "" {
		p.Description = proj.Description
	}
}
