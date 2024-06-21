package models

type CreateCategoryReq struct {
	Title     string `json:"title" validate:"required"`
	ProjectID int    `json:"projectId" validate:"required"`
}

type CategoryResp struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	ProjectID int    `json:"projectId"`
}

type UpdateCategoryReq struct {
	Title string `json:"title"`
}
