package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"store.github.io/pkg/db"
	handlers "store.github.io/pkg/handlers/admin"
)

func main() {
	// Загрузка переменных окружения
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	DB := db.Init()
	h := handlers.New(DB)

	router := gin.Default()
	admin := router.Group("/admin")
	{
		admin.GET("/products", h.GetAllProducts)
		admin.POST("/products", h.CreateProduct)
		admin.PUT("/products/:id", h.UpdateProduct)
		admin.DELETE("/products/:id", h.DeleteProduct)
	}
	// Запуск сервера
	router.Run(":8080")
}
