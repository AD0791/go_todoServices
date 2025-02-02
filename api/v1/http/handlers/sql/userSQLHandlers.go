package sqlhandlers

import (
	sqlmodel "github.com/ad0791/todoServices/api/v1/models/sql_model"
	"github.com/ad0791/todoServices/api/v1/repository/sqlrepository"
	"github.com/ad0791/todoServices/api/v1/schema"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

// validator define in the todo handler

//func hashPassword(password string) (string, error) {
//	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
//	return string(b), err
//}

func RegisterUserRoutes(router fiber.Router) {
	router.Post("/users", CreateUserHandler)
	//router.Get("/users", GetAllUsersHandler)
	//router.Get("/users/:id", GetUserByIDHandler)
	//router.Put("/users/:id", UpdateUserHandler)
	//router.Delete("/users/:id", DeleteUserHandler)
}

// @Summary     Create user
// @Description Create a new user
// @Tags        users
// @Accept      json
// @Produce     json
// @Param       user body schema.UserRequest true "User details"
// @Success     201 {object} schema.UserResponse
// @Failure     400 {object} schema.ErrorMessage
// @Failure     500 {object} schema.ErrorMessage
// @Router      /users [post]
func CreateUserHandler(c *fiber.Ctx) error {
	var req schema.UserRequest

	// Parse incoming JSON
	if err := c.BodyParser(&req); err != nil {
		log.Errorf("Failed to parse request body (create user): %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(schema.ErrorMessage{Error: "Cannot parse JSON"})
	}

	// Validate the request
	if err := validate.Struct(&req); err != nil {
		log.Errorf("Validation failed (create user): %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(schema.ErrorMessage{Error: "Validation error"})
	}

	// If we were hashing:
	// hashedPwd, err := hashPassword(req.Password)
	// if err != nil {
	// 	log.Errorf("Failed to hash password: %v", err)
	// 	return c.Status(fiber.StatusInternalServerError).JSON(schema.ErrorMessage{Error: "Password hashing failed"})
	// }

	user := &sqlmodel.User{
		FullName: req.FullName,
		Email:    req.Email,
		// Password: hashedPwd,
		Password: req.Password, // Temporarily storing raw password (for now)
	}

	if err := sqlrepository.CreateUser(user); err != nil {
		log.Errorf("Database error (create user): %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(schema.ErrorMessage{Error: "Database error"})
	}

	// Prepare response
	resp := &schema.UserResponse{
		ID:        user.ID,
		FullName:  user.FullName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format("2006-01-02"),
	}

	log.Infof("User with ID %d created successfully", user.ID)
	return c.Status(fiber.StatusCreated).JSON(resp)
}
