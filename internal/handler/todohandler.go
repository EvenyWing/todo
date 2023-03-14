package handler

import (
	"evenywing/todo/internal/controller"
	"evenywing/todo/internal/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	db controller.TodoController
}

func NewTodoHandler(db controller.TodoController) TodoHandler {
	return TodoHandler{
		db: db,
	}
}

func (h *TodoHandler) GetTodo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		owners := ctx.Param("owner")

		owneri, err := strconv.Atoi(owners)

		if err != nil {

		}

		h.db.GetTodo(&model.ModelTodo{Owner: uint(owneri)})
	}
}

func (h *TodoHandler) GetTodoByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func (h *TodoHandler) CreateTodo() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func (h *TodoHandler) UpdeateTodo() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func (h *TodoHandler) DeleteTodo() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
