package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/nithishbijadirajeev-droid/forgeflow/app/internal/service"
	"github.com/nithishbijadirajeev-droid/forgeflow/app/pkg/dto"
)

type ProjectHandler struct {
	service *service.ProjectService
}

func NewProjectHandler() *ProjectHandler {
	return &ProjectHandler{
		service: service.NewProjectService(),
	}
}

func (h *ProjectHandler) Create(c *gin.Context) {

	var req dto.CreateProjectRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	project, err := h.service.Create(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, project)
}

func (h *ProjectHandler) GetAll(c *gin.Context) {

	projects, err := h.service.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, projects)
}

func (h *ProjectHandler) GetByID(c *gin.Context) {

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid uuid",
		})
		return
	}

	project, err := h.service.GetByID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "project not found",
		})
		return
	}

	c.JSON(http.StatusOK, project)
}

func (h *ProjectHandler) Update(c *gin.Context) {

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid uuid",
		})
		return
	}

	var req dto.UpdateProjectRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	project, err := h.service.Update(id, req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "project not found",
		})
		return
	}

	c.JSON(http.StatusOK, project)
}

func (h *ProjectHandler) Delete(c *gin.Context) {

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid uuid",
		})
		return
	}

	err = h.service.Delete(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "project deleted",
	})
}