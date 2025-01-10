package admin

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"

	"store.github.io/backend/pkg/db"
	"store.github.io/backend/pkg/models"
)

type UploadResponse struct {
	Message string `json:"message"`
}

type UploadError struct {
	Error string `json:"error"`
}

// UploadAndSaveExcel загружает Excel файл и сохраняет данные в базу
// @Summary Загрузить Excel файл и сохранить товары
// @Description Принимает Excel файл с товарами, парсит его содержимое и сохраняет продукты в базу данных
// @Tags admin
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Excel файл с данными о товарах"
// @Success 200 {object} UploadResponse "Товары успешно добавлены в базу данных"
// @Failure 400 {object} UploadError "Не удалось получить файл"
// @Failure 500 {object} UploadError "Ошибка обработки или сохранения данных"
// @Security BearerAuth
// @Router /admin/upload [post]
func UploadAndSaveExcel(c *gin.Context) {
	// Получаем файл из запроса
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Не удалось получить файл"})
		return
	}

	// Открываем файл
	f, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось открыть файл"})
		return
	}
	defer f.Close()

	// Парсим файл
	excel, err := excelize.OpenReader(f)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось прочитать Excel файл"})
		return
	}

	// Считываем строки из первого листа
	rows, err := excel.GetRows("Лист1")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить строки из Excel"})
		return
	}

	var products []models.Product

	// Преобразуем строки в структуры Product
	for i, row := range rows {
		// Пропускаем заголовок (первая строка)
		if i == 0 {
			continue
		}

		// Проверяем, что достаточно данных в строке
		if len(row) < 5 {
			continue
		}

		// Конвертируем данные
		price, err := strconv.ParseFloat(row[2], 64)
		if err != nil {
			continue
		}

		categoryID, err := strconv.ParseUint(row[4], 10, 32)
		if err != nil {
			continue
		}

		product := models.Product{
			Name:        row[0],
			Description: row[1],
			Price:       price,
			ImageURL:    row[3],
			CategoryID:  uint(categoryID),
		}

		// Добавляем продукт в срез
		products = append(products, product)
	}

	// Записываем продукты в базу данных
	for _, product := range products {
		if err := db.DB.Create(&product).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать товар"})
			return
		}
	}

	// Возвращаем успешный ответ
	c.JSON(http.StatusOK, gin.H{"message": "Продукты успешно добавлены в базу данных"})
}
