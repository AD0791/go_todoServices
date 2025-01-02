package main

import (
	v1 "github.com/ad0791/todoServices/api/v1"
	"github.com/ad0791/todoServices/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	cfg.APP.ENABLEPRINTROUTES = true
	appConfig := fiber.Config{
		AppName:           cfg.APP.NAME,
		ServerHeader:      cfg.APP.SERVERHEADER,
		EnablePrintRoutes: cfg.APP.ENABLEPRINTROUTES,
	}

	app := fiber.New(appConfig)

	v1.RegisterRoutes(app)

	log.Infof("Server start on %s", cfg.APP.Address)
	if err := app.Listen(cfg.APP.Address); err != nil {
		log.Fatalf("Error Server didn't start: %v", err)
	}

}
