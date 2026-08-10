package mocks

import (
	"github.com/nithishbijadirajeev-droid/forgeflow/app/internal/models"
)

type MockUserRepository struct {
	CreateFunc     func(*models.User) error
	GetByEmailFunc func(string) (*models.User, error)
	GetByIDFunc    func(string) (*models.User, error)
}

func (m *MockUserRepository) Create(user *models.User) error {
	return m.CreateFunc(user)
}

func (m *MockUserRepository) GetByEmail(email string) (*models.User, error) {
	return m.GetByEmailFunc(email)
}

func (m *MockUserRepository) GetByID(id string) (*models.User, error) {
	return m.GetByIDFunc(id)
}
