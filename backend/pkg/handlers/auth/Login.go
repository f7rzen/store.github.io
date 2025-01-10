package auth

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"store.github.io/backend/pkg/db"
	"store.github.io/backend/pkg/models"
)

type LoginRequest struct {
	Email    string `json:"email" example:"user@example.com"`
	Password string `json:"password" example:"password123"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token" example:"<JWT_ACCESS_TOKEN>"`
}

type LoginErrorResponse struct {
	Error string `json:"error" example:"Invalid email or password"`
}

var (
	accessTokenSecret  = []byte(os.Getenv("JWT_ACCESS_SECRET"))
	refreshTokenSecret = []byte(os.Getenv("JWT_REFRESH_SECRET"))
)

// GenerateToken создаёт JWT токен с указанными claims и секретным ключом.
func GenerateToken(claims *models.Claims, secret []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

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

	// Проверяем пароль
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, LoginErrorResponse{Error: "Invalid email or password"})
		return
	}

	// Генерация access токена
	accessClaims := &models.Claims{
		UserID:    user.ID,
		Email:     user.Email,
		Role:      "admin",
		TokenType: "access",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	accessTokenString, err := GenerateToken(accessClaims, accessTokenSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, LoginErrorResponse{Error: "Failed to generate access token"})
		return
	}

	// Генерация refresh токена
	refreshClaims := &models.Claims{
		UserID:    user.ID,
		Email:     user.Email,
		Role:      "admin",
		TokenType: "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	refreshTokenString, err := GenerateToken(refreshClaims, refreshTokenSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, LoginErrorResponse{Error: "Failed to generate refresh token"})
		return
	}

	// Устанавливаем refresh токен в куки
	c.SetCookie("refresh_token", refreshTokenString, 7*24*60*60, "/", "", true, true)

	// Возвращаем access токен клиенту
	c.JSON(http.StatusOK, gin.H{
		"accessToken": accessTokenString,
	})
}
