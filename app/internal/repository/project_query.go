package repository

import (
	"strings"

	"gorm.io/gorm"

	"github.com/nithishbijadirajeev-droid/forgeflow/app/internal/models"
	"github.com/nithishbijadirajeev-droid/forgeflow/app/pkg/dto"
)

func buildProjectQuery(db *gorm.DB, query dto.ProjectQuery) *gorm.DB {
	db = db.Model(&models.Project{})

	// Search by project name or description
	if strings.TrimSpace(query.Search) != "" {
		search := "%" + strings.ToLower(strings.TrimSpace(query.Search)) + "%"

		db = db.Where(
			"LOWER(name) LIKE ? OR LOWER(description) LIKE ?",
			search,
			search,
		)
	}

	// Filter by language
	if strings.TrimSpace(query.Language) != "" {
		language := strings.ToLower(strings.TrimSpace(query.Language))
		db = db.Where("LOWER(language) = ?", language)
	}

	// Allow only safe sort fields
	allowedSortFields := map[string]string{
		"name":       "name",
		"language":   "language",
		"created_at": "created_at",
	}

	sortField := "created_at"

	if field, ok := allowedSortFields[strings.ToLower(strings.TrimSpace(query.Sort))]; ok {
		sortField = field
	}

	order := "DESC"

	if strings.ToLower(strings.TrimSpace(query.Order)) == "asc" {
		order = "ASC"
	}

	db = db.Order(sortField + " " + order)

	return db
}
