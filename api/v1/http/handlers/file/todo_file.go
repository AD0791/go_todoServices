package filehandlers

import (
	filemodel "github.com/ad0791/todoServices/api/v1/models/file_model"
	filerepository "github.com/ad0791/todoServices/api/v1/repository/file"
	"github.com/ad0791/todoServices/api/v1/schema"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func RegisterFileTodoRoutes(router fiber.Router) {
	router.Get("/file/todos", FileGetTodos)
	router.Get("/file/todos/:id", FileGetTodosByID)
	router.Post("/file/todos", FileCreateTodo)
	router.Put("/file/todos/:id", FileUpdateTodo)
	router.Delete("/file/todos/:id", FileDeleteTodo)
}

// @Summary		Get all todos
// @Description	Retrieve all todos
// @Tags			file
// @Accept			json
// @Produce		json
// @Success		200	{array}	filemodel.Todo
// @Router			/file/todos [get]
func FileGetTodos(c *fiber.Ctx) error {
	todos, err := filerepository.GetTodos()
	if err != nil {
		log.Errorf("Major issue: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch todos")
	}
	return c.Status(fiber.StatusOK).JSON(todos)
}

// @Summary		Get todo by ID
// @Description	Retrieve a single todo by its ID
// @Tags			file
// @Param			id	path	int	true	"Todo ID"
// @Accept			json
// @Produce		json
// @Success		200	{object}	filemodel.Todo
// @Failure		404	{string}	string	"Todo not found"
// @Router			/file/todos/{id} [get]
func FileGetTodosByID(c *fiber.Ctx) error {
	id := c.Params("id")
	todo, err := filerepository.GetTodoByID(id)
	if err != nil {
		log.Errorf("%s is invalid, no todo found: %v", id, err)
		return fiber.NewError(fiber.StatusNotFound, "Todo not found")
	}

	return c.Status(fiber.StatusOK).JSON(todo)
}

// @Summary		Create a new todo
// @Description	Add a new todo to the list
// @Tags			file
// @Accept			json
// @Produce		json
// @Param			todo	body		schema.TodoRequest	true	"Todo to create"
// @Success		201		{object}	filemodel.Todo
// @Failure		400		{string}	string	"Validation error"
// @Router			/file/todos [post]
func FileCreateTodo(c *fiber.Ctx) error {
	var req *schema.TodoRequest
	if err := c.BodyParser(&req); err != nil {
		log.Errorf("Check the create Todo request: %v", err)
		return fiber.NewError(fiber.StatusBadRequest, "Invalid Request body")
	}

	todo := filemodel.Todo{
		Title:     req.Title,
		Completed: req.Completed,
	}
	newTodo, err := filerepository.CreateTodo(&todo)
	if err != nil {
		log.Errorf("failed to create todo: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create todo")
	}

	return c.Status(fiber.StatusCreated).JSON(newTodo)
}

// @Summary		Update a todo
// @Description	Modify an existing todo
// @Tags			file
// @Accept			json
// @Produce		json
// @Param			id		path		int					true	"Todo ID"
// @Param			todo	body		schema.TodoRequest	true	"Updated todo data"
// @Success		200		{object}	schema.TodoResponse
// @Failure		404		{string}	string	"Todo not found"
// @Failure		400		{string}	string	"Validation error"
// @Router			/file/todos/{id} [put]
func FileUpdateTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	var req *schema.TodoRequest
	if err := c.BodyParser(&req); err != nil {
		log.Errorf("Check the Update Todo request: %v", err)
		return fiber.NewError(fiber.StatusBadRequest, "Invalid Request body")
	}

	todo := schema.TodoResponse{
		Title:     req.Title,
		Completed: req.Completed,
	}

	updatedTodo, err := filerepository.UpdateTodoByID(id, &todo)
	if err != nil {
		log.Errorf("update todo failed: %v", err)
		return fiber.NewError(fiber.StatusNotFound, "Todo aint here for update")
	}
	return c.Status(fiber.StatusCreated).JSON(updatedTodo)
}

// @Summary		Delete a todo
// @Description	Remove a todo from the json list
// @Tags			file
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"Todo ID"
// @Success		202	{object}	schema.MessageResponse
// @Failure		404	{string}	string	"Todo not found"
// @Router			/file/todos/{id} [delete]
func FileDeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	message, err := filerepository.DeleteTodoByID(id)
	if err != nil {
		log.Errorf("delete todo failed: %v", err)
		return fiber.NewError(fiber.StatusNotFound, "Todo aint here for deletion")
	}

	return c.Status(fiber.StatusAccepted).JSON(message)
}
