package controller

import (
	"app/internal/database"
	"app/internal/model"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"
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
	var email string
	err := database.Db.QueryRow("SELECT email FROM vetdonor_donation_stat WHERE email = $1", request.Email).Scan(&email)
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

	query := fmt.Sprintf("UPDATE vetdonor_donation_stat SET %s = %s + 1 WHERE email = $1", request.DonationType, request.DonationType)
	_, updateErr := database.Db.Exec(query, request.Email)
	if updateErr != nil {
		fmt.Println(updateErr)
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

// CheckDonation Загрузить справку.
// @Summary Загрузить справку конкретного юзера
// @Description Загрузка справки пользователя.
// @Accept json
// @Produce json
// @Param email formData string true "Адрес электронной почты пользователя"
// @Param type formData string true "Тип донации пользователя"
// @Param photo formData file true "Фотография пользователя"
// @Success 200 {object} model.CodeResponse "Фотография успешно загружена"
// @Failure 400 {object} model.ErrorResponse "Не удалось загрузить фотографию"
// @Tags Stats
// @Router /v1/check_donation [post]
func CheckDonation(context *gin.Context) {
	file, err := context.FormFile("photo")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	email := context.PostForm("email")
	donationType := context.PostForm("type")

	photo, err := file.Open()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to open photo"})
		return
	}
	defer func(photo multipart.File) {
		err := photo.Close()
		if err != nil {

		}
	}(photo)

	filename := strconv.FormatInt(time.Now().Unix(), 10) + filepath.Ext(file.Filename)

	if err := context.SaveUploadedFile(file, "uploads/"+filename); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}
	filenameWithoutExtension := strings.TrimSuffix(filename, ".png")
	filePath := "https://vetdonor.shmyaks.ru/image/" + filenameWithoutExtension

	err = SaveFilePathInDatabase(email, filePath, donationType)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file path in database"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Photo uploaded successfully"})
}

func SaveFilePathInDatabase(email string, filepath string, dontaionType string) error {
	_, err := database.Db.Exec("INSERT INTO vetdonor_donation (email, photo_path, donationType) VALUES ($1, $2, $3)", email, filepath, dontaionType)
	return err
}

// GetAllNewDonations Получить все спровки из БД.
// @Summary Получить все справки о донациях юзеров
// @Description Получает все справки о новых донациях юзеров.
// @Accept json
// @Produce json
// @Success 200 {object} model.CodeResponse "Заявки успешно получены"
// @Failure 400 {object} model.ErrorResponse "Не удалось получить заявки"
// @Tags Stats
// @Router /v1/get_all_donations [get]
func GetAllNewDonations(context *gin.Context) {
	type Donation struct {
		Id           int    `json:"id"`
		Email        string `json:"email"`
		PhotoPath    string `json:"photoPath"`
		DonationType string `json:"donationType"`
	}

	var donations []Donation

	rows, err := database.Db.Query("SELECT id, email, photo_path, donationType FROM vetdonor_donation")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch updated statistics"})
		return
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	for rows.Next() {
		var donation Donation
		if err := rows.Scan(&donation.Id, &donation.Email, &donation.PhotoPath, &donation.DonationType); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch updated statistics"})
			return
		}
		donations = append(donations, donation)
	}
	if err := rows.Err(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch updated statistics"})
		return
	}

	context.JSON(http.StatusOK, donations)
}

// DeleteDonationApplication Удалить выполненную заявку из БД.
// @Summary Удалить выполненную заявку из БД
// @Description Удаляет запись по указанному айди.
// @Accept json
// @Produce json
// @Param request body model.RequestBody true "Запрос на удаление заявки"
// @Success 200 {object} model.CodeResponse "Заявка успешно удалена"
// @Failure 400 {object} model.ErrorResponse "Не удалось удалить заявку"
// @Tags Stats
// @Router /v1/delete_application [post]
func DeleteDonationApplication(context *gin.Context) {
	var requestBody model.RequestBody
	if err := context.BindJSON(&requestBody); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	query := "DELETE FROM vetdonor_donation WHERE id = $1"
	_, err := database.Db.Exec(query, requestBody.ID)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete donation application"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Donation application deleted successfully"})
}

// GetImage Фото справки.
// @Summary Фото справки
// @Description Возвращает фото справки пользователя.
// @Accept json
// @Produce json
// @Param key path string true "image id"
// @Success 200 {object} model.CodeResponse "Фото получено"
// @Failure 400 {object} model.ErrorResponse "Не удалось получить фото"
// @Tags Photo
// @Router /v1/image/{key} [get]
func GetImage(context *gin.Context) {
	key := context.Param("key")
	imagePath := "uploads/" + key + ".png"

	imageData, err := ioutil.ReadFile(imagePath)
	if err != nil {
		context.String(http.StatusInternalServerError, "Ошибка при чтении изображения")
		return
	}

	context.Header("Content-Type", http.DetectContentType(imageData))
	context.Data(http.StatusOK, "image", imageData)
}
