package dto

import "time"

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

type ProjectResponse struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	RepositoryURL string    `json:"repository_url"`
	Language      string    `json:"language"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
