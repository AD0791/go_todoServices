package sqlhandlers

import (
	"strconv"

	sqlmodel "github.com/ad0791/todoServices/api/v1/models/sql_model"
	"github.com/ad0791/todoServices/api/v1/repository/sqlrepository"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func RegisterSQLTodoRoutes(router fiber.Router) {
	router.Get("/sql/todos", GetSQLTodos)
	router.Get("/sql/todos/:id", GetSQLTodoByID)
	router.Post("/sql/todos", CreateSQLTodo)
	router.Put("/sql/todos/:id", UpdateSQLTodo)
	router.Delete("/sql/todos/:id", DeleteSQLTodo)
}

// @Summary     Get all todos
// @Description Get all todos from database
// @Tags        sql
// @Accept      json
// @Produce     json
// @Success     200 {array}  sqlmodel.Todo
// @Router      /sql/todos [get]
func GetSQLTodos(c *fiber.Ctx) error {
	todos, err := sqlrepository.GetSQLTodos()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error fetching todos",
		})
	}
	return c.Status(fiber.StatusOK).JSON(todos)
}

// @Summary     Get todo by ID
// @Description Get single todo by ID
// @Tags        sql
// @Accept      json
// @Produce     json
// @Param       id path int true "Todo ID"
// @Success     200 {object} sqlmodel.Todo
// @Router      /sql/todos/{id} [get]
func GetSQLTodoByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	todo, err := sqlrepository.GetSQLTodoByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Todo not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(todo)
}

// @Summary     Create todo
// @Description Create new todo
// @Tags        sql
// @Accept      json
// @Produce     json
// @Param       todo body sqlmodel.Todo true "Todo object"
// @Success     201 {object} sqlmodel.Todo
// @Router      /sql/todos [post]
func CreateSQLTodo(c *fiber.Ctx) error {
	todo := new(sqlmodel.Todo)
	if err := c.BodyParser(&todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	if err := validate.Struct(todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Validation error",
		})
	}

	if err := sqlrepository.CreateSQLTodo(todo); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not create todo",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(todo)
}

// @Summary     Update todo
// @Description Update todo by ID
// @Tags        sql
// @Accept      json
// @Produce     json
// @Param       id path int true "Todo ID"
// @Param       todo body sqlmodel.Todo true "Todo object"
// @Success     200 {object} sqlmodel.Todo
// @Router      /sql/todos/{id} [put]
func UpdateSQLTodo(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	todo, err := sqlrepository.GetSQLTodoByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Todo not found",
		})
	}

	updatedTodo := new(sqlmodel.Todo)
	if err := c.BodyParser(updatedTodo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	todo.Title = updatedTodo.Title
	todo.Completed = updatedTodo.Completed

	if err := sqlrepository.UpdateSQLTodo(todo); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not update todo",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(todo)
}

// @Summary     Delete todo
// @Description Delete todo by ID
// @Tags        sql
// @Accept      json
// @Produce     json
// @Param       id path int true "Todo ID"
// @Success		202	{object}	fiber.Map
// @Failure		404	{string}	string	"Todo not found"
// @Router      /sql/todos/{id} [delete]
func DeleteSQLTodo(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	if err := sqlrepository.DeleteSQLTodo(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not delete todo",
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "Todo was deleted",
	})
}
