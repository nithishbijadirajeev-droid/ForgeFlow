package repository

import (
	"github.com/google/uuid"

	"github.com/nithishbijadirajeev-droid/forgeflow/app/internal/database"
	"github.com/nithishbijadirajeev-droid/forgeflow/app/internal/models"
	"github.com/nithishbijadirajeev-droid/forgeflow/app/pkg/dto"
)

type ProjectRepositoryInterface interface {
	Create(project *models.Project) error
	GetAll(query dto.ProjectQuery) ([]models.Project, int64, error)
	GetByID(id uuid.UUID) (*models.Project, error)
	Update(project *models.Project) error
	Delete(id uuid.UUID) error
}

type ProjectRepository struct{}

func NewProjectRepository() *ProjectRepository {
	return &ProjectRepository{}
}

func (r *ProjectRepository) Create(project *models.Project) error {
	return database.DB.Create(project).Error
}

func (r *ProjectRepository) GetAll(query dto.ProjectQuery) ([]models.Project, int64, error) {

	var (
		projects []models.Project
		total    int64
	)

	db := buildProjectQuery(database.DB, query)

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	page := query.Page
	limit := query.Limit

	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 10
	}

	offset := (page - 1) * limit

	err := db.
		Limit(limit).
		Offset(offset).
		Find(&projects).Error

	if err != nil {
		return nil, 0, err
	}

	return projects, total, nil
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
