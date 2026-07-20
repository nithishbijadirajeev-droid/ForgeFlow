package service

import (
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/nithishbijadirajeev-droid/forgeflow/app/internal/models"
	"github.com/nithishbijadirajeev-droid/forgeflow/app/internal/repository"
	"github.com/nithishbijadirajeev-droid/forgeflow/app/pkg/dto"
	jwtutil "github.com/nithishbijadirajeev-droid/forgeflow/app/pkg/jwt"
)

type AuthService struct {
	repository *repository.UserRepository
}

func NewAuthService() *AuthService {
	return &AuthService{
		repository: repository.NewUserRepository(),
	}
}

func (s *AuthService) Register(req dto.RegisterRequest) error {

	_, err := s.repository.GetByEmail(req.Email)
	if err == nil {
		return errors.New("email already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &models.User{
		ID:       uuid.New(),
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hash),
	}

	return s.repository.Create(user)
}

func (s *AuthService) Login(req dto.LoginRequest) (string, error) {

	user, err := s.repository.GetByEmail(req.Email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	)

	if err != nil {
		return "", errors.New("invalid credentials")
	}

	return jwtutil.GenerateToken(user.ID.String())
}