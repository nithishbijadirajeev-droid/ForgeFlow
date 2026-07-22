package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/nithishbijadirajeev-droid/forgeflow/app/internal/service"
	"github.com/nithishbijadirajeev-droid/forgeflow/app/pkg/dto"
	"github.com/nithishbijadirajeev-droid/forgeflow/app/pkg/response"
	"github.com/nithishbijadirajeev-droid/forgeflow/app/pkg/validation"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		service: service.NewAuthService(nil),
	}
}

// Register godoc
//
//	@Summary		Register a new user
//	@Description	Create a new user account
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.RegisterRequest	true	"Register Request"
//	@Success		201		{object}	response.APIResponse
//	@Failure		400		{object}	response.APIResponse
//	@Router			/api/v1/register [post]
func (h *AuthHandler) Register(c *gin.Context) {

	var req dto.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(
			c,
			http.StatusBadRequest,
			"Validation failed",
			validation.FormatValidationErrors(err),
		)
		return
	}

	if err := h.service.Register(req); err != nil {
		response.Error(
			c,
			http.StatusBadRequest,
			"Registration failed",
			err.Error(),
		)
		return
	}

	response.Success(
		c,
		http.StatusCreated,
		"User registered successfully",
		nil,
	)
}

// Login godoc
//
//	@Summary		Authenticate user
//	@Description	Login with email and password
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.LoginRequest	true	"Login Request"
//	@Success		200		{object}	response.APIResponse
//	@Failure		400		{object}	response.APIResponse
//	@Failure		401		{object}	response.APIResponse
//	@Router			/api/v1/login [post]
func (h *AuthHandler) Login(c *gin.Context) {

	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(
			c,
			http.StatusBadRequest,
			"Validation failed",
			validation.FormatValidationErrors(err),
		)
		return
	}

	token, err := h.service.Login(req)

	if err != nil {
		response.Error(
			c,
			http.StatusUnauthorized,
			"Authentication failed",
			err.Error(),
		)
		return
	}

	response.Success(
		c,
		http.StatusOK,
		"Login successful",
		dto.AuthResponse{
			Token: token,
		},
	)
}
