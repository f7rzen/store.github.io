package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"store.github.io/pkg/db"
	"store.github.io/pkg/models"
)

func DeleteProduct(c *gin.Context) {
	// Получаем ID продукта из URL параметров
	id := c.Param("id")

	// Проверяем, существует ли продукт с таким ID
	var product models.Product
	if err := db.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Товар не найден"})
		return
	}

	// Пробуем удалить продукт
	if err := db.DB.Delete(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось удалить товар"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Товар успешно удален"})
}
