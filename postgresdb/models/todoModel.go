package models

type Todo struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Is_completed bool   `json:"isCompleted"`
}
