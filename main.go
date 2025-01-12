package main

import (
	v1 "github.com/ad0791/todoServices/api/v1"
	"github.com/ad0791/todoServices/api/v1/database"
	"github.com/ad0791/todoServices/config"
	_ "github.com/ad0791/todoServices/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
)

// @title			Todo API Documentation
// @version		1.1
// @description	API documentation for Todo service with persistence
// @termsOfService	Developpement
// @contact.name	AD0791
// @contact.email	alexandrodisla@hotmail.com
// @host			localhost:8080
// @BasePath		/api/v1
// @schemes	       http
// @tag.name       file
// @tag.description File-based todo operations
// @tag.name       service
// @tag.description Service-based todo operations
// @tag.name       sql
// @tag.description SQL-based todo operations
// @tag.name       users
// @tag.description User SQL operations
func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	if err := database.InitDatabase(cfg); err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}

	//cfg.APP.ENABLEPRINTROUTES = true
	appConfig := fiber.Config{
		AppName:           cfg.APP.NAME,
		ServerHeader:      cfg.APP.SERVERHEADER,
		EnablePrintRoutes: cfg.APP.ENABLEPRINTROUTES,
	}

	app := fiber.New(appConfig)

	//cfg.APP.AllowCredentials = true
	app.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.APP.AllowOrigins,
		AllowMethods:     cfg.APP.AllowMethods,
		AllowHeaders:     cfg.APP.AllowHeaders,
		AllowCredentials: cfg.APP.AllowCredentials,
	}))

	app.Get("/swagger/*", swagger.HandlerDefault)

	v1.RegisterRoutes(app)

	log.Infof("Server start on %s", cfg.APP.Address)
	if err := app.Listen(cfg.APP.Address); err != nil {
		log.Fatalf("Error Server didn't start: %v", err)
	}

}
