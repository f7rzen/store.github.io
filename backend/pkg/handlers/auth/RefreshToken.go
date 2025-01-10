package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"store.github.io/backend/pkg/models"
)

type RefreshResponse struct {
	AccessToken string `json:"accessToken" example:"<JWT_ACCESS_TOKEN>"`
}

type RefreshError struct {
	Error string `json:"error"`
}

// RefreshToken Обновляет access токен
// @Summary Обновление access токена
// @Description Обновляет access токен с использованием refresh токена, сохраненного в куки.
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} RefreshResponse
// @Failure 401 {object} RefreshError
// @Failure 500 {object} RefreshError
// @Router /refresh [post]
func RefreshToken(c *gin.Context) {
	// Получаем refresh токен из куки
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No refresh token provided"})
		return
	}

	// Разбираем и валидируем refresh токен
	token, err := jwt.ParseWithClaims(refreshToken, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return refreshTokenSecret, nil
	})
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	// Получаем claims из токена
	claims, ok := token.Claims.(*models.Claims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		return
	}

	// Генерация нового access токена
	accessClaims := &models.Claims{
		UserID:    claims.UserID,
		Email:     claims.Email,
		Role:      claims.Role,
		TokenType: "access",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Генерация нового access токена
	accessTokenString, err := GenerateToken(accessClaims, accessTokenSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate new access token"})
		return
	}

	// Отправляем новый access токен в ответе
	c.JSON(http.StatusOK, gin.H{
		"accessToken": accessTokenString,
	})
}
