package main

import (
	"lime-shop-backend/pkg/db"
	"lime-shop-backend/pkg/handlers"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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
	router.GET("/products",h.GetAllProducts)
	router.POST("/products",h.CreateProduct)

	router.DELETE("/products/:id", h.DeleteProduct)
    router.PUT("/products/:id", h.UpdateProduct)

	// Запуск сервера
	router.Run(":8080")
}
