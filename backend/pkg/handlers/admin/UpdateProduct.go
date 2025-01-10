package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"store.github.io/backend/pkg/db"
	"store.github.io/backend/pkg/models"
)

type UpdateError struct {
	Error string `json:"error"`
}

// UpdateProduct обновляет данные товара
// @Summary Обновить товар
// @Description Обновляет информацию о товаре по ID
// @Tags admin
// @Param id path int true "ID товара"
// @Param product body map[string]string{} true "Обновляемые данные товара"
// @Produce json
// @Success 200 {object} ProductResponse "Обновленный товар"
// @Failure 400 {object} UpdateError "Неверные данные"
// @Failure 404 {object} UpdateError "Товар не найден"
// @Failure 500 {object} UpdateError "Не удалось обновить товар"
// @Security BearerAuth
// @Router /admin/products/{id} [put]
func UpdateProduct(c *gin.Context) {
	// Получаем ID продукта из URL параметров
	id := c.Param("id")

	// Проверяем, существует ли продукт с таким ID
	var existingProduct models.Product
	if err := db.DB.First(&existingProduct, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Товар не найден"})
		return
	}

	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}

	// Обновляем только указанные поля
	if err := db.DB.Model(&existingProduct).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось обновить товар"})
		return
	}

	// Возвращаем обновленный продукт
	c.JSON(http.StatusOK, existingProduct)
}
