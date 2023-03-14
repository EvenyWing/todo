package controller

import (
	"evenywing/todo/internal/model"

	"gorm.io/gorm"
)

type TodoController struct {
	DB *gorm.DB
}

func NewTodoController(db *gorm.DB) TodoController {
	return TodoController{
		DB: db,
	}
}

func (c *TodoController) GetTodo(*model.ModelTodo) (error, []model.ModelTodo) {

}

func (c *TodoController) GetTodoByID(*model.ModelTodo) (error, *model.ModelTodo) {

}

func (c *TodoController) CreateTodo(*model.ModelTodo) error {

}

func (c *TodoController) UpdeateTodo(*model.ModelTodo) error {

}

func (c *TodoController) DeleteTodo(*model.ModelTodo) error {

}
