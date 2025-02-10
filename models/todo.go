package models

type Todo struct {
	Title  string `json:"title"`
    Done   bool   `json:"done"`
    UserID uint   `json:"user_id"`
}