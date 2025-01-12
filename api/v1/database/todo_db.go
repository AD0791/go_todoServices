package database

import (
	"fmt"

	sqlmodel "github.com/ad0791/todoServices/api/v1/models/sql_model"
	"github.com/ad0791/todoServices/config"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDatabase(cfg *config.Config) error {
	dsn := fmt.Sprintf(
		"host=database user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=UTC",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DBName,
	)

	log.Infof("We should have this connection string: %s", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Info("failed to connect to the database, review dsn")
		log.Errorf("Failed to connect to the db: %v", err)
		return err
	}

	err = db.AutoMigrate(&sqlmodel.User{}, &sqlmodel.Todo{})
	if err != nil {
		log.Errorf("Failed to migrate the model(s) to the db: %v", err)
		return err
	}

	nameInfo := db.Migrator().CurrentDatabase()

	DB = db

	log.Infof("The current db name: %s", nameInfo)
	return nil
}
