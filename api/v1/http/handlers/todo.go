package handlers

import (
	"github.com/ad0791/todoServices/api/v1/http/services"
	"github.com/ad0791/todoServices/api/v1/schema"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

var validate = validator.New()

func RegisterTodoRoutes(router fiber.Router) {
	router.Get("/todos", GetTodos)
	router.Get("/todos/:id", GetTodoByID)
	router.Post("/todos", CreateTodo)
	router.Put("/todos/:id", UpdateTodo)
	router.Delete("/todos/:id", DeleteTodo)
}

func GetTodos(c *fiber.Ctx) error {
	todos, err := services.FetchTodos()
	if err != nil {
		log.Errorf("Server error from fetching todos: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(todos)
}

func GetTodoByID(c *fiber.Ctx) error {
	id := c.Params("id")
	todo, err := services.FetchTodoByID(id)
	if err != nil {
		log.Errorf("Server error from fetching todo by %s: %v", id, err)
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(todo)

}

func CreateTodo(c *fiber.Ctx) error {
	var req schema.TodoRequest
	if err := c.BodyParser(&req); err != nil {
		log.Errorf("Server error from create todo bodyParser: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if err := validate.Struct(&req); err != nil {
		log.Errorf("Server error from create todo req body unvalide: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	todo, err := services.CreateTodo(&req)
	if err != nil {
		log.Errorf("Server error from create todo service : %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(todo)
}

func UpdateTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	var req schema.TodoRequest
	if err := c.BodyParser(&req); err != nil {
		log.Errorf("Server error from update todo bodyParser: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if err := validate.Struct(&req); err != nil {
		log.Errorf("Server error from update todo req body unvalide: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	updatedTodo, err := services.UpdateTodoByID(id, &req)
	if err != nil {
		log.Errorf("Server error from update todo service : %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(updatedTodo)
}

func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	message, err := services.DeleteTodoByID(id)
	if err != nil {
		log.Errorf("Server error from delete todo by %s: %v", id, err)
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusAccepted).JSON(message)
}
