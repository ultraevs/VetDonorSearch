package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserProfile Профиль пользователя.
// @Summary Профиль
// @Description Возвращает данные пользователя для старнички "profile".
// @Accept json
// @Produce json
// @Success 200 {object} model.CodeResponse "Профиль получен"
// @Failure 400 {object} model.ErrorResponse "Не удалось получить профиль"
// @Tags User
// @Router /v1/profile [get]
func UserProfile(context *gin.Context) {
	email := context.MustGet("Email").(string)
	context.JSON(http.StatusOK, gin.H{"email": email})
}
