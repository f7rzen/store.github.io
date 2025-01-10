package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"store.github.io/backend/docs"
	"store.github.io/backend/pkg/db"
	"store.github.io/backend/pkg/handlers/admin"
	"store.github.io/backend/pkg/handlers/auth"
	"store.github.io/backend/pkg/middleware"
)

// @title Store Admin API
// @description API server for Admin Panel
// @host localhost:8080
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description <b>Enter the token with the: `Bearer ` prefix, e.g. "Bearer eeyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVJ9..."</b>

func init() {
	// Загрузка переменных окружения
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}
	db.ConnectToDb()
	db.MigrationDatabase()
}

func main() {
	docs.SwaggerInfo.BasePath = ""

	router := gin.Default()

	router.Use(middleware.CORSMiddleware())

	router.POST("/signup", auth.Signup)
	router.POST("/login", auth.Login)
	router.POST("/refresh", auth.RefreshToken)
	// Маршрут для обновления токенов
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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

	// Запуск сервера
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
