package model

type ModelTodo struct {
	ID    uint   `json:"id" gorm:"primary_key"`
	Text  string `json:"text"`
	Owner uint   `json:"owner"`
}
