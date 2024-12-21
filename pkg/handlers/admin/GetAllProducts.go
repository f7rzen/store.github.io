package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"store.github.io/pkg/models"
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
