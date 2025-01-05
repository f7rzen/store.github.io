package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"store.github.io/pkg/db"
	"store.github.io/pkg/models"
)

type GetResponse struct {
	Data []ProductResponse `json:"data"`
}

type GetError struct {
	Error string `json:"error"`
}

// GetAllProducts получает список всех товаров
// @Summary Получить все товары
// @Description Возвращает список всех товаров из базы данных
// @Tags admin
// @Produce json
// @Success 200 {object} GetResponse "Список продуктов"
// @Failure 500 {object} GetError "Не удалось получить товары"
// @Security BearerAuth
// @Router /admin/products [get]
func GetAllProducts(c *gin.Context) {
	var products []models.Product
	if err := db.DB.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить товары"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": products})
}
