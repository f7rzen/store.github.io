package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"store.github.io/pkg/db"
	"store.github.io/pkg/models"
)

func GetAllProducts(c *gin.Context) {
	var products []models.Product
	if err := db.DB.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить товары"})
		return
	}
	db.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"data": products})
}
