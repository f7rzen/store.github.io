package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"store.github.io/pkg/db"
	"store.github.io/pkg/models"
)

func UpdateProduct(c *gin.Context) {
	// Получаем ID продукта из URL параметров
	id := c.Param("id")

	// Проверяем, существует ли продукт с таким ID
	var existingProduct models.Product
	if err := db.DB.First(&existingProduct, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Товар не найден"})
		return
	}

	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}

	// Обновляем только указанные поля
	if err := db.DB.Model(&existingProduct).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось обновить товар"})
		return
	}

	// Возвращаем обновленный продукт
	c.JSON(http.StatusOK, existingProduct)
}
