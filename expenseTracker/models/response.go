package models

type SuccessResponse struct {
	Success bool   `json:"suucess"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ErrorResponse struct {
	Success bool   `json:"suucess"`
	Message string `json:"message"`
}
