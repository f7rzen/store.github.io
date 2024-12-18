package handlers

import (
	"net/http"

	"lime-shop-backend/pkg/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetAllProducts(c *gin.Context) {
	var products []models.Product
	if err := h.DB.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить товары"})
		return
	}
	h.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"data": products})
}