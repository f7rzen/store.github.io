package handlers

import (
	"net/http"
	"strings"
	"lime-shop-backend/pkg/models"
	"github.com/gin-gonic/gin"
)

func (h handler) CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}

	if err := h.DB.Create(&product).Error; err != nil {
		// Проверяем на ошибку дублирования
		if strings.Contains(err.Error(), "duplicate key value") {
			c.JSON(http.StatusConflict, gin.H{"error": "Товар с таким названием уже существует"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать товар"})
		return
	}

	c.JSON(http.StatusCreated, product)
}