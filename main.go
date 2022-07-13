package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/vini-fda/todo-list-go-api/database"
	"github.com/vini-fda/todo-list-go-api/todo"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/todo", todo.GetTodos)
	app.Get("/api/v1/todo/:id", todo.GetTodo)
	app.Post("/api/v1/todo", todo.NewTodo)
	app.Delete("/api/v1/todo/:id", todo.DeleteTodo)
}

func main() {
	app := fiber.New()
	if err := database.DBConn.AutoMigrate(&todo.Todo{}); err != nil {
		panic("Could not properly migrate schema.")
	}
	fmt.Println("Database migration completed.")

	setupRoutes(app)
	if err := app.Listen(os.Getenv("API_PORT")); err != nil {
		panic("Could not establish a listening connection.")
	}
}
