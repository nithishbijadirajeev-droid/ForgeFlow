package dto

type CreateProjectRequest struct {
	Name          string `json:"name" binding:"required"`
	Description   string `json:"description"`
	RepositoryURL string `json:"repository_url"`
	Language      string `json:"language"`
}

type UpdateProjectRequest struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	RepositoryURL string `json:"repository_url"`
	Language      string `json:"language"`
}