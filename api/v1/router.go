package v1

import (
	filehandlers "github.com/ad0791/todoServices/api/v1/http/handlers/file"
	servicehandlers "github.com/ad0791/todoServices/api/v1/http/handlers/service"
	sqlhandlers "github.com/ad0791/todoServices/api/v1/http/handlers/sql"
	"github.com/ad0791/todoServices/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func RegisterRoutes(app *fiber.App) {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Errorf("Config error from Route Registers: %v", err)
	}

	api := app.Group(cfg.APP.PREFIX)

	filehandlers.RegisterFileTodoRoutes(api)
	servicehandlers.RegisterServiceTodoRoutes(api)
	sqlhandlers.RegisterSQLTodoRoutes(api)
	sqlhandlers.RegisterUserRoutes(api)
}
