package v1

import (
	"github.com/ad0791/todoServices/api/v1/http/handlers"
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
	handlers.RegisterTodoRoutes(api)
}
