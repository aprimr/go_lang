package models

type Todo struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Is_completed bool   `json:"isCompleted"`
}

type PaginatedTodos struct {
	Data       []Todo `json:"data"`
	Page       int    `json:"page"`
	Limit      int    `json:"limit"`
	TotalCount int    `json:"total_count"`
	TotalPages int    `json:"total_pages"`
}
