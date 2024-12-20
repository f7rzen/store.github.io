package handlers

import (
	"net/http"
	"lime-shop-backend/pkg/models"
	"github.com/gin-gonic/gin"
)

func (h handler) DeleteProduct(c *gin.Context) {
    // Получаем ID продукта из URL параметров
    id := c.Param("id")
    
    // Проверяем, существует ли продукт с таким ID
    var product models.Product
    if err := h.DB.First(&product, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Товар не найден"})
        return
    }

    // Пробуем удалить продукт
    if err := h.DB.Delete(&product).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось удалить товар"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Товар успешно удален"})
}
