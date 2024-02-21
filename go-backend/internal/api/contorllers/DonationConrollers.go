package controller

import (
	"app/internal/database"
	"app/internal/model"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetUserStat Статистика Юзера.
// @Summary Статистика конкретного юзера
// @Description Статистика донаций пользователя.
// @Accept json
// @Param key path string true "user's email key"
// @Produce json
// @Success 200 {object} model.CodeResponse "Статистика получена"
// @Failure 400 {object} model.ErrorResponse "Не удалось получить статистику"
// @Tags Stats
// @Router /v1/get_user_stats/{key} [get]
func GetUserStat(context *gin.Context) {
	Email := context.Param("key")

	var stat model.ResponseStat
	err := database.Db.QueryRow("SELECT blood, plasma, platelets FROM vetdonor_donation_stat WHERE email = $1", Email).Scan(&stat.Blood, &stat.Plasma, &stat.Platelets)
	if err == nil {
		context.JSON(http.StatusOK, gin.H{"blood": stat.Blood, "plasma": stat.Plasma, "platelets": stat.Platelets})
		return
	}
	context.JSON(http.StatusInternalServerError, gin.H{"error": "Can't find stats"})
}

// UpdateUserStat Обновить Статистику.
// @Summary Обновить статистику конкретного юзера
// @Description Статистика донаций пользователя.
// @Accept json
// @Produce json
// @Param request body model.RequestUpdateStat true "Запрос на получение статистики юзера"
// @Success 200 {object} model.CodeResponse "Статистика получена"
// @Failure 400 {object} model.ErrorResponse "Не удалось получить статистику"
// @Tags Stats
// @Router /v1/update_user_stats [put]
func UpdateUserStat(context *gin.Context) {
	var request model.RequestUpdateStat
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Can't read the body"})
		return
	}

	var existingUser model.RequestUpdateStat
	err := database.Db.QueryRow("SELECT email, blood, plasma, platelets FROM vetdonor_donation_stat WHERE email = $1", request.Email).Scan(&existingUser.Email, &existingUser.Blood, &existingUser.Plasma, &existingUser.Platelets)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		_, insertErr := database.Db.Exec("INSERT INTO vetdonor_donation_stat (email, blood, plasma, platelets) VALUES ($1, $2, $3, $4)", request.Email, 0, 0, 0)
		if insertErr != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating new user"})
			return
		}
	case err != nil:
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking existing user"})
		return
	}

	query := "UPDATE vetdonor_donation_stat SET blood = blood + $1, plasma = plasma + $2, platelets = platelets + $3 WHERE email = $4"
	_, updateErr := database.Db.Exec(query, request.Blood, request.Plasma, request.Platelets, request.Email)
	if updateErr != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating data"})
		return
	}

	var currentStats model.ResponseStat
	query = "SELECT blood, plasma, platelets FROM vetdonor_donation_stat WHERE email = $1"
	updateErr = database.Db.QueryRow(query, request.Email).Scan(&currentStats.Blood, &currentStats.Plasma, &currentStats.Platelets)
	if updateErr != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch updated statistics"})
		return
	}
	context.JSON(http.StatusOK, currentStats)
}
