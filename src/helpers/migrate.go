package helpers

import (
	"go-fiber-docker-api/src/config"
	"go-fiber-docker-api/src/models"
	"log"
)

func Migration() {
	err := config.DB.AutoMigrate(
		&models.Product{},
	)

	if err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}
}
