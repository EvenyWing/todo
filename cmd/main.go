package main

import (
	"evenywing/todo/internal/controller"
	"evenywing/todo/internal/handler"
	"evenywing/todo/internal/model"
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	db, err := gorm.Open(sqlite.Open("test.sqlite"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	migrate := flag.Bool("migrate db", false, "migrate db to todolist")
	flag.Parse()
	if *migrate == true {
		err := db.AutoMigrate(&model.ModelTodo{})

		if err != nil {
			fmt.Println(err)

			return
		}

		fmt.Println("DB as succefully migrate")
	}

	c := controller.NewTodoController(db)
	h := handler.NewTodoHandler(c)
	r := gin.Default()
	r.GET("/todos/:owner", h.db.GetTodo())
	r.GET("/todo/:id", h.db.GetTodoByID())
	r.POST("/todo", h.db.CreateTodo())
	r.PUT("/todo", h.db.UpdeateTodo())
	r.DELETE("/todo/:id", h.db.DeleteTodo())
}
