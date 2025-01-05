package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"store.github.io/pkg/db"
	"store.github.io/pkg/models"
)

type DeleteError struct {
	Error string `json:"error"`
}

type DeleteResponse struct {
	Message string `json:"message"`
}

// DeleteProduct удаление товара по ID
// @Summary Удаление товара
// @Description Удалить товар из базы данных по ID
// @Tags admin
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "ID товара"
// @Success 200 {object} DeleteError "Товар успешно удален"
// @Failure 404 {object} DeleteError "Товар не найден"
// @Failure 500 {object} DeleteError "Не удалось удалить товар"
// @Router /admin/products/{id} [delete]
func DeleteProduct(c *gin.Context) {
	// Получаем ID продукта из URL параметров
	id := c.Param("id")

	// Проверяем, существует ли продукт с таким ID
	var product models.Product
	if err := db.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Товар не найден"})
		return
	}

	// Пробуем удалить продукт
	if err := db.DB.Delete(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось удалить товар"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Товар успешно удален"})
}
