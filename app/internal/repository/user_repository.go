package repository

import (
	"github.com/nithishbijadirajeev-droid/forgeflow/app/internal/database"
	"github.com/nithishbijadirajeev-droid/forgeflow/app/internal/models"
)

type UserRepositoryInterface interface {
	Create(user *models.User) error
	GetByEmail(email string) (*models.User, error)
	GetByID(id string) (*models.User, error)
}

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) Create(user *models.User) error {
	return database.DB.Create(user).Error
}

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User

	err := database.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetByID(id string) (*models.User, error) {
	var user models.User

	err := database.DB.First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
