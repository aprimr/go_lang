package models

type Todo struct {
	Id           string `json:"id"`
	Title        string `json:"title"`
	Is_completed bool   `json:"isCompleted"`
}
