package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/nithishbijadirajeev-droid/forgeflow/app/internal/service"
	"github.com/nithishbijadirajeev-droid/forgeflow/app/pkg/dto"
	"github.com/nithishbijadirajeev-droid/forgeflow/app/pkg/response"
)

type ProjectHandler struct {
	service *service.ProjectService
}

func NewProjectHandler() *ProjectHandler {
	return &ProjectHandler{
		service: service.NewProjectService(nil),
	}
}

// Create godoc
//
//	@Summary		Create a new project
//	@Description	Create a project
//	@Tags			Projects
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.CreateProjectRequest	true	"Create Project"
//	@Success		201		{object}	response.APIResponse
//	@Failure		400		{object}	response.APIResponse
//	@Failure		500		{object}	response.APIResponse
//	@Router			/api/v1/projects [post]
func (h *ProjectHandler) Create(c *gin.Context) {

	var req dto.CreateProjectRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Validation failed", err.Error())
		return
	}

	project, err := h.service.Create(req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Unable to create project", err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "Project created successfully", project)
}

// GetAll godoc
//
//	@Summary		Get all projects
//	@Description	Retrieve all projects
//	@Tags			Projects
//	@Security		BearerAuth
//	@Produce		json
//	@Success		200	{object}	response.APIResponse
//	@Failure		500	{object}	response.APIResponse
//	@Router			/api/v1/projects [get]
func (h *ProjectHandler) GetAll(c *gin.Context) {

	projects, err := h.service.GetAll()

	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Unable to retrieve projects", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Projects retrieved successfully", projects)
}

// GetByID godoc
//
//	@Summary		Get project by ID
//	@Description	Retrieve a project
//	@Tags			Projects
//	@Security		BearerAuth
//	@Produce		json
//	@Param			id	path		string	true	"Project ID"
//	@Success		200	{object}	response.APIResponse
//	@Failure		400	{object}	response.APIResponse
//	@Failure		404	{object}	response.APIResponse
//	@Router			/api/v1/projects/{id} [get]
func (h *ProjectHandler) GetByID(c *gin.Context) {

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid project ID", "invalid uuid")
		return
	}

	project, err := h.service.GetByID(id)
	if err != nil {
		response.Error(c, http.StatusNotFound, "Project not found", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Project retrieved successfully", project)
}

// Update godoc
//
//	@Summary		Update project
//	@Description	Update a project
//	@Tags			Projects
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string						true	"Project ID"
//	@Param			request	body		dto.UpdateProjectRequest	true	"Update Project"
//	@Success		200		{object}	response.APIResponse
//	@Failure		400		{object}	response.APIResponse
//	@Failure		404		{object}	response.APIResponse
//	@Router			/api/v1/projects/{id} [put]
func (h *ProjectHandler) Update(c *gin.Context) {

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid project ID", "invalid uuid")
		return
	}

	var req dto.UpdateProjectRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Validation failed", err.Error())
		return
	}

	project, err := h.service.Update(id, req)
	if err != nil {
		response.Error(c, http.StatusNotFound, "Project not found", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Project updated successfully", project)
}

// Delete godoc
//
//	@Summary		Delete project
//	@Description	Delete a project
//	@Tags			Projects
//	@Security		BearerAuth
//	@Produce		json
//	@Param			id	path		string	true	"Project ID"
//	@Success		200	{object}	response.APIResponse
//	@Failure		400	{object}	response.APIResponse
//	@Failure		500	{object}	response.APIResponse
//	@Router			/api/v1/projects/{id} [delete]
func (h *ProjectHandler) Delete(c *gin.Context) {

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid project ID", "invalid uuid")
		return
	}

	if err := h.service.Delete(id); err != nil {
		response.Error(c, http.StatusInternalServerError, "Unable to delete project", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Project deleted successfully", nil)
}
