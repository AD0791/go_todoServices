package sqlhandlers

import (
	"github.com/ad0791/todoServices/api/v1/schema"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New()

func hashPassword(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(b), err
}

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
	if err := c.BodyParser(&req)
}
