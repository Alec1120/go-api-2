package main

import (
	"github.com/gin-gonic/gin"
	"go-api/src/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("./database/todos.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&handlers.Todo{})
	return db
}

var todos = []handlers.Todo{
	{ID: 1, Task: "Learn Go", Done: false},
	{ID: 2, Task: "Build a REST API", Done: false},
}

func main() {
	db := initDB()
	
	r := gin.Default()

	r.GET("/todos", handlers.GetTodos(db))
	r.POST("/todos", handlers.CreateTodo(db))
	r.PUT("/todos/:id", handlers.UpdateTodo(db))
	r.DELETE("/todos/:id", handlers.DeleteTodo(db))

	r.Run(":8080")
}
