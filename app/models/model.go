package models

type Model struct {
	ID uint32 `json:"id"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
	DeleteTime string `json:"delete_time"`
}
