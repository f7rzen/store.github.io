package admin

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"

	"store.github.io/pkg/db"
	"store.github.io/pkg/models"
)

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

	// Парсим файл с помощью excelize
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