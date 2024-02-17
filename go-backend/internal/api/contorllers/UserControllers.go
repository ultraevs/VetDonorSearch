package controller

import (
	"app/internal/database"
	"app/internal/model"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Profile Профиль пользователя.
// @Summary Профиль
// @Description Возвращает профиль пользователя.
// @Accept json
// @Produce json
// @Success 200 {object} model.CodeResponse "Профиль получен"
// @Failure 400 {object} model.ErrorResponse "Не удалось получить профиль"
// @Tags User
// @Router /v1/profile [get]
func Profile(context *gin.Context) {
	email := context.MustGet("Email").(string)
	Type := context.MustGet("Type").(string)
	fmt.Println(email, Type)
	switch Type {
	case "user":
		{
			var user model.UserInfo
			err := database.Db.QueryRow("SELECT name, email FROM vetdonor_users WHERE email = $1", email).Scan(&user.Name, &user.Email)
			if err != nil && !errors.Is(err, sql.ErrNoRows) {
				context.JSON(http.StatusInternalServerError, gin.H{"error": "Can't find User"})
				return
			}
			context.JSON(http.StatusOK, gin.H{"name": user.Name, "email": user.Email})
			return
		}
	case "clinic":
		{
			var clinic model.ClinicInfo
			err := database.Db.QueryRow("SELECT name, email, address FROM vetdonor_clinic WHERE email = $1", email).Scan(&clinic.Name, &clinic.Email, &clinic.Address)
			if err != nil && !errors.Is(err, sql.ErrNoRows) {
				context.JSON(http.StatusInternalServerError, gin.H{"error": "Can't find Clinic"})
				return
			}
			context.JSON(http.StatusOK, gin.H{"name": clinic.Name, "email": clinic.Email, "address": clinic.Address})
			return
		}
	}

}

// UserQuestionnaire Анкета пользователя.
// @Summary Анкета
// @Description Возвращает анкету пользователя для старнички "profile".
// @Accept json
// @Produce json
// @Success 200 {object} model.OneOfUserClinicQuestionnaireResponse "Анкета получена"
// @Failure 400 {object} model.ErrorResponse "Не удалось получить анкету"
// @Tags User
// @Router /v1/questionnaire [get]
func UserQuestionnaire(context *gin.Context) {
	email := context.MustGet("Email").(string)
	var clinic model.ClinicQuestionnaire
	isFind := [2]bool{true, true}
	var bloodTypesJSON []byte
	err := database.Db.QueryRow("SELECT bloodtypes, workhours, contacts FROM vetdonor_clinic_questionnaire WHERE email = $1", email).Scan(&bloodTypesJSON, &clinic.WorkHours, &clinic.Contacts)
	err = json.Unmarshal(bloodTypesJSON, &clinic.BloodTypes)
	if err == nil {
		context.JSON(http.StatusOK, gin.H{
			"BloodTypes": clinic.BloodTypes,
			"WorkHours":  clinic.WorkHours,
			"Contacts":   clinic.Contacts,
		})
		return
	} else if errors.Is(err, sql.ErrNoRows) {
		isFind[0] = false
	}
	var user model.UserQuestionnaire
	err = database.Db.QueryRow("SELECT city, surname, breed, petname, bloodtype, age FROM vetdonor_user_questionnaire WHERE email = $1", email).Scan(&user.City, &user.Surname, &user.Breed, &user.PetName, &user.BloodType, &user.Age)
	if err == nil {
		context.JSON(http.StatusOK, gin.H{
			"City":      user.City,
			"Surname":   user.Surname,
			"Breed":     user.Breed,
			"PetName":   user.PetName,
			"BloodType": user.BloodType,
			"Age":       user.Age,
		})
		return
	} else if errors.Is(err, sql.ErrNoRows) {
		isFind[1] = false
	}
	if isFind[0] == false && isFind[1] == false {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Questionnaire isn't founded"})
		return
	}
}

// CreateQuestionnaire Создать анкету.
// @Summary Создать анкету
// @Description Создает анкету пользователя или клиники
// @Consumes application/json
// @Produce json
// @Param request body model.QuestionnaireRequest true "Запрос на создание анкеты пользователя или клиники"
// @Success 200 {object} model.CodeResponse "Успешное создание анкеты"
// @Failure 400 {object} model.ErrorResponse "Не удалось создать анкету"
// @Tags User
// @Router /v1/create_questionnaire [post]
func CreateQuestionnaire(context *gin.Context) {
	var request model.QuestionnaireRequest
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Can't read the body"})
		return
	}

	switch request.Type {
	case "user":
		var user model.CreateUserQuestionnaire
		if err := json.Unmarshal(request.Data, &user); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Can't parse the user data"})
			return
		}

		var existingUserID int
		err := database.Db.QueryRow("SELECT id FROM vetdonor_user_questionnaire WHERE email = $1", user.Email).Scan(&existingUserID)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query database"})
			return
		}

		if existingUserID != 0 {
			_, err := database.Db.Exec("UPDATE vetdonor_user_questionnaire SET city = $2, surname = $3, breed = $4, petname = $5, bloodtype = $6, age = $7 WHERE email = $1", user.Email, user.City, user.Surname, user.Breed, user.PetName, user.BloodType, user.Age)
			if err != nil {
				context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user's questionnaire"})
				return
			}
			context.JSON(http.StatusOK, gin.H{"message": "User's questionnaire updated successfully"})
			return
		}

		_, err = database.Db.Exec("INSERT INTO vetdonor_user_questionnaire (email, city, surname, breed, petname, bloodtype, age) VALUES ($1, $2, $3, $4, $5, $6, $7)", user.Email, user.City, user.Surname, user.Breed, user.PetName, user.BloodType, user.Age)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user's questionnaire"})
			return
		}
		context.JSON(http.StatusOK, gin.H{"message": "User's questionnaire created successfully"})
	case "clinic":
		var clinic model.CreateClinicQuestionnaire
		if err := json.Unmarshal(request.Data, &clinic); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Can't parse the clinic data"})
			return
		}

		var existingClinicID int
		err := database.Db.QueryRow("SELECT id FROM vetdonor_clinic_questionnaire WHERE email = $1", clinic.Email).Scan(&existingClinicID)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query database"})
			return
		}

		if existingClinicID != 0 {
			bloodTypesJSON, err := json.Marshal(clinic.BloodTypes)
			if err != nil {
				fmt.Println(err)
				context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create clinic's questionnaire"})
				return
			}

			_, err = database.Db.Exec("UPDATE vetdonor_clinic_questionnaire SET bloodtypes = $2, workhours = $3, contacts = $4 WHERE email = $1", clinic.Email, string(bloodTypesJSON), clinic.WorkHours, clinic.Contacts)
			if err != nil {
				fmt.Println(err)
				context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update clinic's questionnaire"})
				return
			}
			context.JSON(http.StatusOK, gin.H{"message": "Clinic's questionnaire updated successfully"})
			return
		}

		bloodTypesJSON, err := json.Marshal(clinic.BloodTypes)
		if err != nil {
			fmt.Println(err)
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create clinic's questionnaire"})
			return
		}
		_, err = database.Db.Exec("INSERT INTO vetdonor_clinic_questionnaire (email, bloodtypes, workhours, contacts) VALUES ($1, $2, $3, $4)", clinic.Email, string(bloodTypesJSON), clinic.WorkHours, clinic.Contacts)
		if err != nil {
			fmt.Println(err)
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create clinic's questionnaire"})
			return
		}
		context.JSON(http.StatusOK, gin.H{"message": "Clinic's questionnaire created successfully"})
	default:
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid questionnaire type"})
	}
}

// GetProfile Профиль другого пользователя.
// @Summary Чужой Профиль
// @Description Возвращает профиль другого пользователя.
// @Accept json
// @Produce json
// @Param key path string true "user's email key"
// @Success 200 {object} model.CodeResponse "Профиль получен"
// @Failure 400 {object} model.ErrorResponse "Не удалось получить профиль"
// @Tags User
// @Router /v1/profile/{key} [get]
func GetProfile(context *gin.Context) {
	Email := context.Param("key")
	var clinic model.ClinicInfo
	err := database.Db.QueryRow("SELECT email, name, address FROM vetdonor_clinic WHERE email = $1", Email).Scan(&clinic.Email, &clinic.Name, &clinic.Address)
	if err == nil {
		context.JSON(http.StatusOK, gin.H{"email": clinic.Email, "name": clinic.Name, "address": clinic.Address})
		return
	}

	var user model.UserInfo
	err = database.Db.QueryRow("SELECT email, name FROM vetdonor_users WHERE email = $1", Email).Scan(&user.Email, &user.Name)
	if err == nil {
		context.JSON(http.StatusOK, gin.H{"email": user.Email, "name": user.Name})
		return
	}

	context.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
}

// MarkAsNeed Пометить анкету как нуждающегося.
// @Summary Пометить анкету как нуждающегося
// @Description Помещает анкету в отдельную базу нуждающихся
// @Consumes application/json
// @Produce json
// @Param request body model.UserInfo true "Запрос на перемещение анкеты пользователя в раздел нуждающихся"
// @Success 200 {object} model.CodeResponse "Успешное перемещение анкеты"
// @Failure 400 {object} model.ErrorResponse "Не удалось переместить анкету"
// @Tags User
// @Router /v1/mark [post]
func MarkAsNeed(context *gin.Context) {
	var request model.UserInfo
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Can't read the body"})
		return
	}

	var count int
	err := database.Db.QueryRow("SELECT COUNT(*) FROM vetdonor_need WHERE email = $1", request.Email).Scan(&count)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check if email exists"})
		return
	}

	if count > 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "User with this email already marked as need"})
		return
	}

	_, err = database.Db.Exec("INSERT INTO vetdonor_need(email, name) VALUES ($1, $2)", request.Email, request.Name)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to move user's questionnaire"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "User's questionnaire moved successfully"})
}

// CreateUserOtherInfo Создать анкету юзера.
// @Summary Создать анкету юзера
// @Description Создает анкету юзера для данных в профиле
// @Consumes application/json
// @Produce json
// @Param request body model.CreateUserOtherInfo true "Запрос на создание анкеты пользователя"
// @Success 200 {object} model.CodeResponse "Успешное создание анкеты"
// @Failure 400 {object} model.ErrorResponse "Не удалось создать анкету"
// @Tags User
// @Router /v1/create_other_info [post]
func CreateUserOtherInfo(context *gin.Context) {
	var request model.CreateUserOtherInfo
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Can't read the body"})
		return
	}

	_, err := database.Db.Exec("INSERT INTO vetdonor_user_other_info(email, name, surname, patronymic, age, gender, about) VALUES ($1, $2, $3, $4, $5, $6, $7)", request.Email, request.Name, request.Surname, request.Patronymic, request.Age, request.Gender, request.About)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user's other info"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "User's other info created successfully"})
}

// GetUserOtherInfo Получить анкету юзера.
// @Summary Получить анкету юзера
// @Description Получает анкету пользователя для профиля
// @Consumes application/json
// @Produce json
// @Param request body model.RequestUserOtherInfo true "Запрос на получение анкеты пользователя"
// @Success 200 {object} model.CodeResponse "Успешно получена анкета"
// @Failure 400 {object} model.ErrorResponse "Не удалось получить анкету"
// @Tags User
// @Router /v1/other_info [get]
func GetUserOtherInfo(context *gin.Context) {
	var request model.RequestUserOtherInfo
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Can't read the body"})
		return
	}
	var form model.ResponseUserOtherInfo
	err := database.Db.QueryRow("SELECT name, surname, patronymic, age, gender, about FROM vetdonor_user_other_info WHERE email = $1", request.Email).Scan(&form.Name, &form.Surname, &form.Patronymic, &form.Age, &form.Gender, &form.About)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user's other info"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"name": form.Name, "surname": form.Surname, "patronymic": form.Patronymic, "age": form.Age, "gender": form.Gender, "about": form.About})
}
