package service

import (
	"github.com/google/uuid"

	"github.com/nithishbijadirajeev-droid/forgeflow/app/internal/models"
	"github.com/nithishbijadirajeev-droid/forgeflow/app/internal/repository"
	"github.com/nithishbijadirajeev-droid/forgeflow/app/pkg/dto"
)

type ProjectService struct {
	repository *repository.ProjectRepository
}

func NewProjectService() *ProjectService {
	return &ProjectService{
		repository: repository.NewProjectRepository(),
	}
}

func (s *ProjectService) Create(req dto.CreateProjectRequest) (*models.Project, error) {

	project := &models.Project{
		Name:          req.Name,
		Description:   req.Description,
		RepositoryURL: req.RepositoryURL,
		Language:      req.Language,
	}

	err := s.repository.Create(project)
	if err != nil {
		return nil, err
	}

	return project, nil
}

func (s *ProjectService) GetAll() ([]models.Project, error) {
	return s.repository.GetAll()
}

func (s *ProjectService) GetByID(id uuid.UUID) (*models.Project, error) {
	return s.repository.GetByID(id)
}

func (s *ProjectService) Update(id uuid.UUID, req dto.UpdateProjectRequest) (*models.Project, error) {

	project, err := s.repository.GetByID(id)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		project.Name = req.Name
	}

	if req.Description != "" {
		project.Description = req.Description
	}

	if req.RepositoryURL != "" {
		project.RepositoryURL = req.RepositoryURL
	}

	if req.Language != "" {
		project.Language = req.Language
	}

	err = s.repository.Update(project)
	if err != nil {
		return nil, err
	}

	return project, nil
}

func (s *ProjectService) Delete(id uuid.UUID) error {
	return s.repository.Delete(id)
}
