package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"store.github.io/pkg/db"
	"store.github.io/pkg/models"
)

// SignupRequest представляет параметры запроса для регистрации.
type SignupRequest struct {
	Email    string `json:"email" example:"user@example.com"`
	Password string `json:"password" example:"password123"`
}

// SignupResponse представляет успешный ответ на регистрацию.
type SignupResponse struct {
	Message string `json:"message" example:"User registered successfully"`
}

// ErrorResponse представляет структуру ошибки.
type SignupErrorResponse struct {
	Error string `json:"error" example:"Invalid JSON"`
}

// Signup регистрирует пользователя.
// @Summary Регистрация пользователя
// @Description Регистрирует нового пользователя в системе
// @Tags auth
// @Accept json
// @Produce json
// @Param input body SignupRequest true "Параметры регистрации"
// @Success 200 {object} SignupResponse
// @Failure 400 {object} SignupErrorResponse
// @Router /signup [post]
func Signup(c *gin.Context) {
	var body SignupRequest

	// Парсинг JSON из тела запроса
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, SignupErrorResponse{Error: "Invalid JSON"})
		return
	}

	// Хеширование пароля
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, SignupErrorResponse{Error: "Failed to hash password"})
		return
	}

	// Создание нового пользователя
	user := models.User{
		Email:    body.Email,
		Password: string(hash),
	}
	if err := db.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, SignupErrorResponse{Error: "Failed to create user"})
		return
	}

	// Успешный ответ
	c.JSON(http.StatusOK, SignupResponse{Message: "User registered successfully"})
}
