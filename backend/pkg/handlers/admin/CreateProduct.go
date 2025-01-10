package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"store.github.io/backend/pkg/db"
	"store.github.io/backend/pkg/models"
)

type CreateError struct {
	Error string `json:"error"`
}
type CreateRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"gte=0"`
	ImageURL    string  `json:"image_url"`
	CategoryID  uint    `json:"category_id"`
}

type ProductResponse struct {
	ID          uint    `json:"ID"`
	CreatedAt   string  `json:"Created_at"`
	UpdatedAt   string  `json:"Updated_at"`
	DeletedAt   string  `json:"DeletedAt"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageURL    string  `json:"image_url"`
	CategoryID  uint    `json:"category_id"`
}

// CreateProduct создает новый товар
// @Summary Создание товара
// @Description Добавить новый товар в базу данных
// @Tags admin
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param product body CreateRequest true "Параметры нового товара"
// @Success 201 {object} ProductResponse
// @Failure 400 {object} CreateError
// @Failure 500 {object} CreateError
// @Router /admin/products [post]
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
