package models

type NewProjectReq struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

type ProjectResp struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

type UpdateProjectReq struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
