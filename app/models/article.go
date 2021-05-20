package models

type Article struct {
	*Model
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Cover   string `json:"cover"`
	Content string `json:"content"`
	State   string `json:"state"`
}
