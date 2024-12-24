package db

import (
	"log"

	"store.github.io/pkg/models"
)

func MigrationDatabase() {
	if err := DB.AutoMigrate(&models.Product{}); err != nil {
		log.Fatalf("ошибка миграции таблицы Product: %w", err)
	}
	if err := DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("ошибка миграции таблицы User: %w", err)
	}
}