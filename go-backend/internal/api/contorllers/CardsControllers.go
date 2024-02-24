package controller

import (
	"app/internal/database"
	"app/internal/model"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
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
	rows, err := database.Db.Query("SELECT name, email FROM vetdonor_need LIMIT 10")
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
			fmt.Println(err)
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
	rows, err := database.Db.Query("SELECT name, email, address FROM vetdonor_clinic LIMIT 10")
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

// GetAllUserCard Все Карточки Юзеров.
// @Summary Все карточки юзеров
// @Description Возвращает все карточки нуждающихся.
// @Accept json
// @Produce json
// @Success 200 {object} model.CodeResponse "Карточки получены"
// @Failure 400 {object} model.ErrorResponse "Не удалось получить карточки"
// @Tags Cards
// @Router /v1/get_all_user_cards [get]
func GetAllUserCard(context *gin.Context) {
	rows, err := database.Db.Query("SELECT name, email FROM vetdonor_need")
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
			fmt.Println(err)
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

// GetAllClinicCard Все Карточки Клиник.
// @Summary Все карточки клиник
// @Description Возвращает все карточеки клиник.
// @Accept json
// @Produce json
// @Success 200 {object} model.CodeResponse "Карточки получены"
// @Failure 400 {object} model.ErrorResponse "Не удалось получить карточки"
// @Tags Cards
// @Router /v1/get_all_clinic_cards [get]
func GetAllClinicCard(context *gin.Context) {
	rows, err := database.Db.Query("SELECT name, email, address, bloodtypesincluded, bloodtypesnotincluded, workhours, contacts FROM vetdonor_clinic")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query database"})
		return
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
		}
	}(rows)
	var clinics []model.ClinicQuestionnaire
	for rows.Next() {
		var clinic model.ClinicQuestionnaire
		err := rows.Scan(&clinic.Name, &clinic.Email, &clinic.Address, pq.Array(&clinic.BloodTypesIncluded), pq.Array(&clinic.BloodTypesNotIncluded), &clinic.WorkHours, &clinic.Contacts)
		if err != nil {
			fmt.Println(err)
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

// GetAllDonors Все Карточки Доноров.
// @Summary Все карточки доноров
// @Description Возвращает все карточеки доноров.
// @Accept json
// @Produce json
// @Success 200 {object} model.CodeResponse "Карточки получены"
// @Failure 400 {object} model.ErrorResponse "Не удалось получить карточки"
// @Tags Cards
// @Router /v1/get_all_donors_cards [get]
func GetAllDonors(context *gin.Context) {
	rows, err := database.Db.Query("SELECT email, name, petname, blood, breed, contacts FROM vetdonor_pet_donors")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query database"})
		return
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
		}
	}(rows)
	var donors []model.SetDonor
	for rows.Next() {
		var donor model.SetDonor
		err := rows.Scan(&donor.Email, &donor.Name, &donor.PetName, &donor.Blood, &donor.Breed, &donor.Contacts)
		if err != nil {
			fmt.Println(err)
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan rows"})
			return
		}
		donors = append(donors, donor)
	}
	if err := rows.Err(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to iterate over rows"})
		return
	}
	context.JSON(http.StatusOK, donors)
}
