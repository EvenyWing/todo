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

	migrate := flag.Bool("migrate", false, "migrate db to todolist")
	flag.Parse()
	if *migrate == true {
		err := db.AutoMigrate(&model.ModelTodo{})

		if err != nil {
			fmt.Println(err)

			return
		}

		fmt.Println("DB as succefully migrate")
		return
	}

	c := controller.NewTodoController(db)
	h := handler.NewTodoHandler(c)
	r := gin.Default()
	r.GET("/todos/:owner", h.GetTodo())
	r.GET("/todo/:id", h.GetTodoByID())
	r.POST("/todo", h.CreateTodo())
	r.PUT("/todo", h.UpdeateTodo())
	r.DELETE("/todo/:id", h.DeleteTodo())

	r.Run()
}
