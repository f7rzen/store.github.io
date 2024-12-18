package handlers

import (
	"net/http"

	"lime-shop-backend/pkg/models"
	"github.com/gin-gonic/gin"
)

func (h handler) UpdateProduct(c *gin.Context) {
    // Получаем ID продукта из URL параметров
    id := c.Param("id")

    var product models.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
        return
    }

    // Находим продукт по ID и обновляем его
    if err := h.DB.Model(&models.Product{}).Where("id = ?", id).Updates(product).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось обновить товар"})
        return
    }

    // Возвращаем обновленный продукт
    c.JSON(http.StatusOK, product)
}
