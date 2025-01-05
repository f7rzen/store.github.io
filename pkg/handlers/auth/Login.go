package auth

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"store.github.io/pkg/db"
	"store.github.io/pkg/models"
)

// LoginRequest представляет параметры запроса авторизации.
type LoginRequest struct {
	Email    string `json:"email" example:"user@example.com"`
	Password string `json:"password" example:"password123"`
}

// LoginResponse представляет успешный ответ авторизации.
type LoginResponse struct {
	Token   string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	Message string `json:"message" example:"Login successful"`
}

type LoginErrorResponse struct {
	Error string `json:"error" example:"Invalid email or password"`
}

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

// Login авторизует пользователя и возвращает JWT-токен.
// @Summary Авторизация пользователя
// @Description Авторизует пользователя, возвращая JWT-токен
// @Tags auth
// @Accept json
// @Produce json
// @Param request body LoginRequest true "Параметры авторизации"
// @Success 200 {object} LoginResponse
// @Failure 400 {object} LoginErrorResponse
// @Failure 401 {object} LoginErrorResponse
// @Failure 500 {object} LoginErrorResponse
// @Router /login [post]
func Login(c *gin.Context) {
	var body LoginRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, LoginErrorResponse{Error: "Invalid JSON"})
		return
	}

	var user models.User
	if err := db.DB.Where("email = ?", body.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, LoginErrorResponse{Error: "Invalid email or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, LoginErrorResponse{Error: "Invalid email or password"})
		return
	}

	// Создаём токен
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &models.Claims{
		UserID: user.ID,
		Email:  user.Email,
		Role:   "admin",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, LoginErrorResponse{Error: "Failed to generate token"})
		return
	}

	// Возвращаем токен клиенту
	c.JSON(http.StatusOK, LoginResponse{
		Token:   tokenString,
		Message: "Login successful. You can now access the admin dashboard.",
	})
}
