package handlers

import (
	"net/http"

	"lime-shop-backend/pkg/models"
	"github.com/gin-gonic/gin"
)

func (h handler) UpdateProduct(c *gin.Context) {
    // Получаем ID продукта из URL параметров
    id := c.Param("id")

    // Проверяем, существует ли продукт с таким ID
    var existingProduct models.Product
    if err := h.DB.First(&existingProduct, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Товар не найден"})
        return
    }

    var updates map[string]interface{}
    if err := c.ShouldBindJSON(&updates); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
        return
    }

    // Обновляем только указанные поля
    if err := h.DB.Model(&existingProduct).Updates(updates).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось обновить товар"})
        return
    }

    // Возвращаем обновленный продукт
    c.JSON(http.StatusOK, existingProduct)
}
