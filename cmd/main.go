package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "store.github.io/docs"
	"store.github.io/pkg/db"
	"store.github.io/pkg/handlers/admin"
	"store.github.io/pkg/handlers/auth"
	"store.github.io/pkg/middleware"
)

func init() {
	// Загрузка переменных окружения
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}
	db.ConnectToDb()
	db.MigrationDatabase()
}

// @title     Gingo Bookstore API
func main() {

	router := gin.Default()

	router.POST("/signup", auth.Signup)
	router.POST("/login", auth.Login)

	// Группа маршрутов для админки
	dashboard := router.Group("/admin")
	dashboard.Use(middleware.AuthMiddleware) // Применяем middleware
	{
		dashboard.GET("/products", admin.GetAllProducts)
		dashboard.POST("/products", admin.CreateProduct)
		dashboard.PUT("/products/:id", admin.UpdateProduct)
		dashboard.DELETE("/products/:id", admin.DeleteProduct)
		dashboard.POST("/upload", admin.UploadAndSaveExcel)
	}

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Запуск сервера
	router.Run(":8080")
}
