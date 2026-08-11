package repository

import (
	"strings"

	"gorm.io/gorm"

	"github.com/nithishbijadirajeev-droid/forgeflow/app/internal/models"
	"github.com/nithishbijadirajeev-droid/forgeflow/app/pkg/dto"
)

func buildProjectQuery(db *gorm.DB, query dto.ProjectQuery) *gorm.DB {

	db = db.Model(&models.Project{})

	// Search by project name
	if query.Search != "" {
		db = db.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(query.Search)+"%")
	}

	// Filter by language
	if query.Language != "" {
		db = db.Where("language = ?", query.Language)
	}

	// Sorting
	sortField := "created_at"

	if query.Sort != "" {
		sortField = query.Sort
	}

	order := "DESC"

	if strings.ToUpper(query.Order) == "ASC" {
		order = "ASC"
	}

	db = db.Order(sortField + " " + order)

	return db
}
