package controller

import (
	"app/internal/database"
	"app/internal/model"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserCard(context *gin.Context) {

}

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
