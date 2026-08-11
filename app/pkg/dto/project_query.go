package dto

type ProjectQuery struct {
	Page int `form:"page"`

	Limit int `form:"limit"`

	Search string `form:"search"`

	Language string `form:"language"`

	Sort string `form:"sort"`

	Order string `form:"order"`
}

type PaginationMeta struct {
	Page int `json:"page"`

	Limit int `json:"limit"`

	Total int64 `json:"total"`

	TotalPages int `json:"total_pages"`
}
