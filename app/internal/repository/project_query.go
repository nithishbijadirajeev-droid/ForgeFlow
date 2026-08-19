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
	// Search by project name or description
	if query.Search != "" {

		search := "%" + strings.ToLower(query.Search) + "%"

		db = db.Where(
			"LOWER(name) LIKE ? OR LOWER(description) LIKE ?",
			search,
			search,
		)
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
