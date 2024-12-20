package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"store.github.io/pkg/models"
)

func (h handler) CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}

	if err := h.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать товар"})
		return
	}

	c.JSON(http.StatusCreated, product)
}
