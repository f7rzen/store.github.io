package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"store.github.io/pkg/db"
	"store.github.io/pkg/models"
)

func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}

	if err := db.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать товар"})
		return
	}

	c.JSON(http.StatusCreated, product)
}
