package repository

import (
	"github.com/google/uuid"
	"github.com/nithishbijadirajeev-droid/forgeflow/app/internal/database"
	"github.com/nithishbijadirajeev-droid/forgeflow/app/internal/models"
)

type ProjectRepository struct{}

func NewProjectRepository() *ProjectRepository {
	return &ProjectRepository{}
}

func (r *ProjectRepository) Create(project *models.Project) error {
	return database.DB.Create(project).Error
}

func (r *ProjectRepository) GetAll() ([]models.Project, error) {
	var projects []models.Project
	err := database.DB.Find(&projects).Error
	return projects, err
}

func (r *ProjectRepository) GetByID(id uuid.UUID) (*models.Project, error) {
	var project models.Project

	err := database.DB.First(&project, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &project, nil
}

func (r *ProjectRepository) Update(project *models.Project) error {
	return database.DB.Save(project).Error
}

func (r *ProjectRepository) Delete(id uuid.UUID) error {
	return database.DB.Delete(&models.Project{}, "id = ?", id).Error
}