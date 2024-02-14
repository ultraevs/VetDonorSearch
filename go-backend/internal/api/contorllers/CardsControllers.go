package controller

import (
	"app/internal/database"
	"app/internal/model"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetUserCard Карточки Юзеров.
// @Summary Первые 10 карточек юзеров
// @Description Возвращает 10 карточек нуждающихся.
// @Accept json
// @Produce json
// @Success 200 {object} model.CodeResponse "Карточки получены"
// @Failure 400 {object} model.ErrorResponse "Не удалось получить карточки"
// @Tags Cards
// @Router /v1/get_user_cards [get]
func GetUserCard(context *gin.Context) {
	rows, err := database.Db.Query("SELECT * FROM vetdonor_need LIMIT 10")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query database"})
		return
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
		}
	}(rows)
	var users []model.UserInfo
	for rows.Next() {
		var user model.UserInfo
		err := rows.Scan(&user.Name, &user.Email)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan rows"})
			return
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to iterate over rows"})
		return
	}
	context.JSON(http.StatusOK, users)
}

// GetClinicCard Карточки Клиник.
// @Summary Первые 10 карточек клиник
// @Description Возвращает 10 карточек клиник.
// @Accept json
// @Produce json
// @Success 200 {object} model.CodeResponse "Карточки получены"
// @Failure 400 {object} model.ErrorResponse "Не удалось получить карточки"
// @Tags Cards
// @Router /v1/get_clinic_cards [get]
func GetClinicCard(context *gin.Context) {
	rows, err := database.Db.Query("SELECT * FROM vetdonor_clinic LIMIT 10")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query database"})
		return
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
		}
	}(rows)
	var clinics []model.ClinicInfo
	for rows.Next() {
		var clinic model.ClinicInfo
		err := rows.Scan(&clinic.Name, &clinic.Email, &clinic.Address)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan rows"})
			return
		}
		clinics = append(clinics, clinic)
	}
	if err := rows.Err(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to iterate over rows"})
		return
	}
	context.JSON(http.StatusOK, clinics)
}
