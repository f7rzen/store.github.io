package handlers

import (
	"net/http"

	"lime-shop-backend/pkg/models"
	"github.com/gin-gonic/gin"
)

func (h handler) DeleteProduct(c *gin.Context) {
    // Получаем ID продукта из URL параметров
    id := c.Param("id")

    // Пробуем удалить продукт из базы данных
    if err := h.DB.Unscoped().Delete(&models.Product{}, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось удалить товар"})
        return
    }

    // Возвращаем успешный ответ
    c.JSON(http.StatusOK, gin.H{"message": "Товар успешно удален"})
}
