package dto

type NewCategory struct {
	Title     string
	ProjectID int
}

type UpdateCategory struct {
	ID    int
	Title string
}
