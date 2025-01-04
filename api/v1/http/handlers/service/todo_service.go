package servicehandlers

import (
	"github.com/ad0791/todoServices/api/v1/http/services"
	"github.com/ad0791/todoServices/api/v1/schema"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

var validate = validator.New()

func RegisterServiceTodoRoutes(router fiber.Router) {
	router.Get("service/todos", GetTodos)
	router.Get("service/todos/:id", GetTodoByID)
	router.Post("service/todos", CreateTodo)
	router.Put("service/todos/:id", UpdateTodo)
	router.Delete("service/todos/:id", DeleteTodo)
}

// @Summary		Get all todos
// @Description	Retrieve all todos
// @Tags			service
// @Accept			json
// @Produce		json
// @Success		200	{array}	schema.TodoResponse
// @Router			/service/todos [get]
func GetTodos(c *fiber.Ctx) error {
	todos, err := services.FetchTodos()
	if err != nil {
		log.Errorf("Server error from fetching todos: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(todos)
}

// @Summary		Get todo by ID
// @Description	Retrieve a single todo by its ID
// @Tags			service
// @Param			id	path	int	true	"Todo ID"
// @Accept			json
// @Produce		json
// @Success		200	{object}	schema.TodoResponse
// @Failure		404	{string}	string	"Todo not found"
// @Router			/service/todos/{id} [get]
func GetTodoByID(c *fiber.Ctx) error {
	id := c.Params("id")
	todo, err := services.FetchTodoByID(id)
	if err != nil {
		log.Errorf("Server error from fetching todo by %s: %v", id, err)
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(todo)

}

// @Summary		Create a new todo
// @Description	Add a new todo to the list
// @Tags			service
// @Accept			json
// @Produce		json
// @Param			todo	body		schema.TodoRequest	true	"Todo to create"
// @Success		201		{object}	schema.TodoResponse
// @Failure		400		{string}	string	"Validation error"
// @Router			/service/todos [post]
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

// @Summary		Update a todo
// @Description	Modify an existing todo
// @Tags			service
// @Accept			json
// @Produce		json
// @Param			id		path		int					true	"Todo ID"
// @Param			todo	body		schema.TodoRequest	true	"Updated todo data"
// @Success		200		{object}	schema.TodoResponse
// @Failure		404		{string}	string	"Todo not found"
// @Failure		400		{string}	string	"Validation error"
// @Router			/service/todos/{id} [put]
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

// @Summary		Delete a todo
// @Description	Remove a todo from the list
// @Tags			service
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"Todo ID"
// @Success		202	{object}	schema.MessageResponse
// @Failure		404	{string}	string	"Todo not found"
// @Router			/service/todos/{id} [delete]
func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	message, err := services.DeleteTodoByID(id)
	if err != nil {
		log.Errorf("Server error from delete todo by %s: %v", id, err)
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusAccepted).JSON(message)
}
