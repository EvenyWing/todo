package handler

import (
	"evenywing/todo/internal/controller"
	"evenywing/todo/internal/model"
	"fmt"
	"net/http"
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
			fmt.Println(err)
		}

		var todos []model.ModelTodo
		err, todos = h.db.GetTodo(&model.ModelTodo{Owner: uint(owneri)})

		if err != nil {
			fmt.Println(err)
		}

		ctx.JSON(http.StatusAccepted, todos)
	}
}

func (h *TodoHandler) GetTodoByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ids := ctx.Param("id")

		id, err := strconv.Atoi(ids)

		if err != nil {
			fmt.Println(err)
		}

		todo := &model.ModelTodo{}

		err, todo = h.db.GetTodoByID(&model.ModelTodo{ID: uint(id)})

		if err != nil {
			fmt.Println(err)
		}

		ctx.JSON(http.StatusAccepted, todo)
	}
}

func (h *TodoHandler) CreateTodo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		todo := &model.ModelTodo{}

		err := ctx.BindJSON(todo)

		if err != nil {
			fmt.Println(err)
		}

		err = h.db.CreateTodo(todo)

		if err != nil {
			fmt.Println(err)
		}
	}
}

func (h *TodoHandler) UpdeateTodo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		todo := &model.ModelTodo{}

		err := ctx.BindJSON(todo)

		if err != nil {
			fmt.Println(err)
		}

		err = h.db.UpdeateTodo(todo)

		if err != nil {
			fmt.Println(err)
		}
	}
}

func (h *TodoHandler) DeleteTodo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ids := ctx.Param("id")

		id, err := strconv.Atoi(ids)

		if err != nil {
			fmt.Println(err)
		}

		err = h.db.DeleteTodo(&model.ModelTodo{ID: uint(id)})

		if err != nil {
			fmt.Println(err)
		}
	}
}
