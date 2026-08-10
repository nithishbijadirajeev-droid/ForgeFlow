package service

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/nithishbijadirajeev-droid/forgeflow/app/internal/models"
	"github.com/nithishbijadirajeev-droid/forgeflow/app/pkg/dto"
	"github.com/nithishbijadirajeev-droid/forgeflow/app/test/mocks"
)

func TestRegisterSuccess(t *testing.T) {

	mockRepo := &mocks.MockUserRepository{}

	mockRepo.GetByEmailFunc = func(email string) (*models.User, error) {
		return nil, errors.New("not found")
	}

	mockRepo.CreateFunc = func(user *models.User) error {

		assert.Equal(t, "Nithish", user.Name)
		assert.Equal(t, "nithish@test.com", user.Email)

		// Password should NOT be stored as plain text
		assert.NotEqual(t, "Password123", user.Password)

		return nil
	}

	mockRepo.GetByIDFunc = func(id string) (*models.User, error) {
		return &models.User{
			ID: uuid.New(),
		}, nil
	}

	service := NewAuthService(mockRepo)

	err := service.Register(dto.RegisterRequest{
		Name:     "Nithish",
		Email:    "nithish@test.com",
		Password: "Password123",
	})

	assert.NoError(t, err)
}

func TestRegisterDuplicateEmail(t *testing.T) {

	mockRepo := &mocks.MockUserRepository{}

	mockRepo.GetByEmailFunc = func(email string) (*models.User, error) {
		return &models.User{
			ID:    uuid.New(),
			Email: email,
		}, nil
	}

	mockRepo.CreateFunc = func(user *models.User) error {
		t.Fatal("Create() should not be called when email already exists")
		return nil
	}

	mockRepo.GetByIDFunc = func(id string) (*models.User, error) {
		return nil, nil
	}

	service := NewAuthService(mockRepo)

	err := service.Register(dto.RegisterRequest{
		Name:     "Nithish",
		Email:    "nithish@test.com",
		Password: "Password123",
	})

	assert.Error(t, err)
	assert.Equal(t, "email already exists", err.Error())
}
