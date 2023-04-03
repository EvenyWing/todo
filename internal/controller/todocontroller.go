package controller

import (
	"evenywing/todo/internal/model"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type TodoController struct {
	db *gorm.DB
}

func NewTodoController(db *gorm.DB) TodoController {
	return TodoController{
		db: db,
	}
}

func (c *TodoController) GetTodo(todo *model.ModelTodo) (error, []model.ModelTodo) {

	var todos []model.ModelTodo
	err := c.db.Model(&model.ModelTodo{}).Where("owner = ?", todo.Owner).Find(&todos).Error

	if err != nil {
		return errors.Wrap(err, "Failed to retrivied all todo"), nil
	}
	return nil, todos
}

func (c *TodoController) GetTodoByID(todo *model.ModelTodo) (error, *model.ModelTodo) {

	err := c.db.Model(&model.ModelTodo{}).Find(&todo).Error

	if err != nil {
		return errors.Wrap(err, "Failed to get todo by id"), nil
	}

	return nil, todo
}

func (c *TodoController) CreateTodo(todo *model.ModelTodo) error {

	err := c.db.Model(&model.ModelTodo{}).Create(&todo).Error

	if err != nil {
		return errors.Wrap(err, "Failed to create todo in db")
	}

	return nil
}

func (c *TodoController) UpdeateTodo(todo *model.ModelTodo) error {

	err := c.db.Model(&model.ModelTodo{}).Where("id = ?", todo.ID).Update("text", todo.Text).Error

	if err != nil {
		return errors.Wrap(err, "Failed to update todo")
	}

	return nil
}

func (c *TodoController) DeleteTodo(todo *model.ModelTodo) error {

	err := c.db.Model(&model.ModelTodo{}).Where("id = ?", todo.ID).Delete(&todo).Error

	if err != nil {
		errors.Wrap(err, "Failed to delete todo")
	}

	return nil
}
