package dto

import "time"

type CreateProjectRequest struct {
	Name string `json:"name" binding:"required,min=2,max=100"`

	Description string `json:"description" binding:"max=500"`

	RepositoryURL string `json:"repository_url" binding:"omitempty,url"`

	Language string `json:"language" binding:"required,min=2,max=50"`
}

type UpdateProjectRequest struct {
	Name string `json:"name" binding:"omitempty,min=2,max=100"`

	Description string `json:"description" binding:"omitempty,max=500"`

	RepositoryURL string `json:"repository_url" binding:"omitempty,url"`

	Language string `json:"language" binding:"omitempty,min=2,max=50"`
}

type ProjectResponse struct {
	ID string `json:"id"`

	Name string `json:"name"`

	Description string `json:"description"`

	RepositoryURL string `json:"repository_url"`

	Language string `json:"language"`

	CreatedAt time.Time `json:"created_at"`

	UpdatedAt time.Time `json:"updated_at"`
}
