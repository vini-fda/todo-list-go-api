package todo

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/vini-fda/todo-list-go-api/database"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    uint8  `json:"priority"`
	Checked     bool   `json:"checked"`
}

func GetTodos(c *fiber.Ctx) error {
	db := database.DBConn
	var todos []Todo
	db.Find(&todos)
	return c.JSON(todos)
}

func GetTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var todo Todo
	db.Find(&todo, id)
	return c.JSON(todo)
}

func NewTodo(c *fiber.Ctx) error {
	db := database.DBConn
	todo := new(Todo)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&todo)
	return c.JSON(todo)
}

func UpdateTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var todo Todo
	db.First(&todo, id)
	if todo.Title == "" {
		return c.Status(500).SendString("No 'Todo' found with given ID")
	}
	if err := c.BodyParser(&todo); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	db.Save(&todo)
	return c.JSON(todo)
}

func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var todo Todo
	db.First(&todo, id)
	if todo.Title == "" {
		return c.Status(500).SendString("No 'Todo' found with given ID")
	}
	db.Delete(&todo)
	return c.SendString("'Todo' successfully deleted!")
}
