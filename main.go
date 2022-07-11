package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/vini-fda/todo-list-go-api/database"
	"github.com/vini-fda/todo-list-go-api/todo"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Database connection successfully opened")
	database.DBConn.AutoMigrate(&todo.Todo{})
	fmt.Println("Database migrated")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/todo", todo.GetTodos)
	app.Get("/api/v1/todo/:id", todo.GetTodo)
	app.Post("/api/v1/todo", todo.NewTodo)
	app.Delete("/api/v1/todo/:id", todo.DeleteTodo)
}

func main() {
	app := fiber.New()

	initDatabase()

	setupRoutes(app)
	app.Listen(":3000")
}
