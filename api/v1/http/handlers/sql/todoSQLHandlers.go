package sqlhandlers

import (
	"strconv"

	sqlmodel "github.com/ad0791/todoServices/api/v1/models/sql_model"
	"github.com/ad0791/todoServices/api/v1/repository/sqlrepository"
	"github.com/ad0791/todoServices/api/v1/schema"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
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
// @Description Retrieve all todos from the database
// @Tags        sql
// @Accept      json
// @Produce     json
// @Success     200 {array} schema.TodoSQLResponse
// @Failure     500 {string} string "Database error"
// @Router      /sql/todos [get]
func GetSQLTodos(c *fiber.Ctx) error {
	todos, err := sqlrepository.GetSQLTodos()
	if err != nil {
		log.Errorf("Failed to fetch todos from DB: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
	}

	var resp []*schema.TodoSQLResponse

	for _, todo := range todos {
		resp = append(resp, &schema.TodoSQLResponse{
			ID:        todo.ID,
			Title:     todo.Title,
			Completed: todo.Completed,
			CreatedAt: todo.CreatedAt.Format("2006-01-02"),
			UpdatedAt: todo.UpdatedAt.Format("2006-01-02"),
		})
	}
	log.Infof("Successfully fetched %d todos", len(resp))
	return c.Status(fiber.StatusOK).JSON(resp)
}

// @Summary     Get todo by ID
// @Description Retrieve a single todo by its ID
// @Tags        sql
// @Accept      json
// @Produce     json
// @Param       id path int true "Todo ID"
// @Success     200 {object} schema.TodoSQLResponse
// @Failure     404 {string} string "Todo not found"
// @Failure     400 {string} string "Invalid ID"
// @Router      /sql/todos/{id} [get]
func GetSQLTodoByID(c *fiber.Ctx) error {
	//id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Errorf("ID not valid to get todo: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	todo, err := sqlrepository.GetSQLTodoByID(uint(id))
	if err != nil {
		log.Errorf("Todo with ID %d not found: %v", id, err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Todo not found",
		})
	}

	resp := &schema.TodoSQLResponse{
		ID:        todo.ID,
		Title:     todo.Title,
		Completed: todo.Completed,
		CreatedAt: todo.CreatedAt.Format("2006-01-02"),
	}
	log.Infof("Successfully fetched todo with ID %d", id)
	return c.Status(fiber.StatusOK).JSON(resp)
}

// @Summary     Create todo
// @Description Create a new todo
// @Tags        sql
// @Accept      json
// @Produce     json
// @Param       todo body schema.TodoRequest true "Todo request body"
// @Success     201 {object} schema.TodoSQLResponse
// @Failure     400 {string} string "Invalid input"
// @Failure     500 {string} string "Database error"
// @Router      /sql/todos [post]
func CreateSQLTodo(c *fiber.Ctx) error {
	var reqTodo schema.TodoRequest

	if err := c.BodyParser(&reqTodo); err != nil {
		log.Errorf("Failed to parse request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	if err := validate.Struct(&reqTodo); err != nil {
		log.Errorf("Validation failed: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Validation error",
		})
	}

	todo := sqlmodel.Todo{
		Title:     reqTodo.Title,
		Completed: reqTodo.Completed,
	}
	if err := sqlrepository.CreateSQLTodo(&todo); err != nil {
		log.Errorf("Failed to create todo in DB: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
	}

	resp := &schema.TodoSQLResponse{
		ID:        todo.ID,
		Title:     todo.Title,
		Completed: todo.Completed,
		CreatedAt: todo.CreatedAt.Format("2006-01-02"),
	}
	log.Infof("Successfully created todo with ID %d", todo.ID)
	return c.Status(fiber.StatusCreated).JSON(resp)
}

// @Summary     Update todo
// @Description Update an existing todo
// @Tags        sql
// @Accept      json
// @Produce     json
// @Param       id path int true "Todo ID"
// @Param       todo body schema.TodoRequest true "Updated todo data"
// @Success     200 {object} schema.TodoSQLResponse
// @Failure     404 {string} string "Todo not found"
// @Failure     400 {string} string "Invalid ID"
// @Failure     500 {string} string "Database error"
// @Router      /sql/todos/{id} [put]
func UpdateSQLTodo(c *fiber.Ctx) error {
	//id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Errorf("invalid id format to update a todo: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	todo, err := sqlrepository.GetSQLTodoByID(uint(id))
	if err != nil {
		log.Errorf("Todo with ID %d not found: %v", id, err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Todo not found",
		})
	}

	var reqTodo schema.TodoRequest
	if err := c.BodyParser(&reqTodo); err != nil {
		log.Errorf("body parser issue to update todo: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	// Validate the request body
	if err := validate.Struct(&reqTodo); err != nil {
		log.Errorf("Validation issue to update todo: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Validation failed",
		})
	}

	todo.Title = reqTodo.Title
	todo.Completed = reqTodo.Completed
	if err := sqlrepository.UpdateSQLTodo(todo); err != nil {
		log.Errorf("Failed to update todo with ID %d: %v", todo.ID, err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not update todo",
		})
	}

	response := &schema.TodoSQLResponse{
		ID:        todo.ID,
		Title:     todo.Title,
		Completed: todo.Completed,
		CreatedAt: todo.CreatedAt.Format("2006-01-02"),
		UpdatedAt: todo.UpdatedAt.Format("2006-01-02"),
	}
	log.Infof("Successfully updated todo with ID %d", todo.ID)
	return c.Status(fiber.StatusOK).JSON(response)
}

// @Summary     Delete todo
// @Description Soft delete a todo by ID
// @Tags        sql
// @Accept      json
// @Produce     json
// @Param       id path int true "Todo ID"
// @Success     202 {object} schema.MessageSQLResponse
// @Failure     404 {string} string "Todo not found"
// @Failure     400 {string} string "Invalid ID"
// @Failure     500 {string} string "Database error"
// @Router      /sql/todos/{id} [delete]
func DeleteSQLTodo(c *fiber.Ctx) error {
	//id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Errorf("Invalid ID for delete request: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	// Delete the todo and retrieve the soft-deleted record
	deletedTodo, err := sqlrepository.DeleteAndGetSQLTodoByID(uint(id))
	if err != nil {
		log.Errorf("Failed to delete todo with ID %d: %v", id, err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Todo not found or could not be deleted",
		})
	}

	// Prepare the response
	response := &schema.MessageSQLResponse{
		MessageResponse: schema.MessageResponse{
			ID:      int(deletedTodo.ID),
			Message: "Todo was successfully deleted",
		},
		DeletedAt: deletedTodo.DeletedAt.Time.Format("2006-01-02"),
	}

	log.Infof("Todo with ID %d successfully soft-deleted at %s", id, response.DeletedAt)
	return c.Status(fiber.StatusAccepted).JSON(response)
}
