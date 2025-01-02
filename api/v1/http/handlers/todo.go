package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterTodoRoutes(router fiber.Router) {
	router.Get("/todos", GetTodos)
	router.Get("/todos/:id", GetTodoByID)
	router.Post("/todos", CreateTodo)
	router.Put("/todos/:id", UpdateTodo)
	router.Delete("/todos/:id", DeleteTodo)
}

func GetTodos(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusAccepted)
}

func GetTodoByID(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusAccepted)
}

func CreateTodo(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusCreated)
}

func UpdateTodo(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

func DeleteTodo(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusContinue)
}
