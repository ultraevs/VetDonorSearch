package controller

import (
	"app/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Message Сообщение ассистенту.
// @Summary Сообщение ассистенту
// @Description Сообщение ассистенту.
// @Accept json
// @Produce json
// @Param request body model.MessageBody true "Сообщение ассистенту"
// @Success 200 {object} model.CodeResponse "Сообщене получено"
// @Failure 400 {object} model.ErrorResponse "Не удалось получить сообщение"
// @Tags Assistant
// @Router /v1/message [post]
func Message(context *gin.Context) {
	var mes model.MessageBody
	// Работа скрипта
	context.JSON(http.StatusOK, gin.H{"callback": "TESTETST", "message": mes})
}
