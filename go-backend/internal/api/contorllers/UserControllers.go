package controller

import (
	"app/internal/database"
	"app/internal/model"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"net/http"
	"strings"
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

			var otherInfo model.ResponseUserOtherInfo
			err = database.Db.QueryRow("SELECT city, phone, telegram, path FROM vetdonor_user_other_info WHERE email = $1", email).Scan(&otherInfo.City, &otherInfo.Phone, &otherInfo.Telegram, &otherInfo.Path)
			if err != nil && !errors.Is(err, sql.ErrNoRows) {
				context.JSON(http.StatusInternalServerError, gin.H{"error": "Can't find User"})
				return
			}

			context.JSON(http.StatusOK, gin.H{"type": "user", "name": user.Name, "email": user.Email, "city": otherInfo.City, "phone": otherInfo.Phone, "telegram": otherInfo.Telegram, "path": otherInfo.Path})
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
			context.JSON(http.StatusOK, gin.H{"type": "clinic", "name": clinic.Name, "email": clinic.Email, "address": clinic.Address})
			return
		}
	}

}

// CreateClinicQuestionnaire Создать анкету клиники.
// @Summary Создать анкету
// @Description Создает анкету клиники
// @Consumes application/json
// @Produce json
// @Param request body model.CreateClinicQuestionnaire true "Запрос на создание анкеты клиники"
// @Success 200 {object} model.CodeResponse "Успешное создание анкеты"
// @Failure 400 {object} model.ErrorResponse "Не удалось создать анкету"
// @Tags User
// @Router /v1/create_questionnaire [post]
func CreateClinicQuestionnaire(context *gin.Context) {
	var clinic model.CreateClinicQuestionnaire
	if err := context.ShouldBindJSON(&clinic); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Can't read the body"})
		return
	}

	var existingClinicID int
	err := database.Db.QueryRow("SELECT id FROM vetdonor_clinic WHERE email = $1", clinic.Email).Scan(&existingClinicID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query database"})
		return
	}

	if existingClinicID != 0 {
		_, err := database.Db.Exec("UPDATE vetdonor_clinic SET bloodtypesincluded = $2, workhours = $3, contacts = $4, bloodtypesnotincluded = $5 WHERE email = $1", clinic.Email, pq.Array(clinic.BloodTypesIncluded), clinic.WorkHours, clinic.Contacts, pq.Array(clinic.BloodTypesNotIncluded))
		if err != nil {
			fmt.Println(err)
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update clinic's questionnaire"})
			return
		}
		context.JSON(http.StatusOK, gin.H{"message": "Clinic's questionnaire updated successfully"})
		return
	}

	//_, err = database.Db.Exec("INSERT INTO vetdonor_clinic (email, bloodtypesincluded, workhours, contacts, bloodtypesnotincluded) VALUES ($1, $2, $3, $4, $5)", clinic.Email, pq.Array(clinic.BloodTypesIncluded), clinic.WorkHours, clinic.Contacts, pq.Array(clinic.BloodTypesNotIncluded))
	//if err != nil {
	//	fmt.Println(err)
	//	context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create clinic's questionnaire"})
	//	return
	//}
	context.JSON(http.StatusOK, gin.H{"message": "Clinic's questionnaire created successfully"})
}

// CreatePetQuestionnaire Создать анкету питомцв.
// @Summary Создать питомца
// @Description Создает анкету питомца
// @Consumes application/json
// @Produce json
// @Param request body model.CreateUserQuestionnaire true "Запрос на создание анкеты питомца"
// @Success 200 {object} model.CodeResponse "Успешное создание анкеты"
// @Failure 400 {object} model.ErrorResponse "Не удалось создать анкету"
// @Tags User
// @Router /v1/create_user_questionnaire [post]
func CreatePetQuestionnaire(context *gin.Context) {
	var pet model.CreateUserQuestionnaire
	if err := context.ShouldBindJSON(&pet); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Can't read the body"})
		return
	}

	var existingClinicID int
	err := database.Db.QueryRow("SELECT id FROM vetdonor_users WHERE email = $1", pet.Email).Scan(&existingClinicID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query database"})
		return
	}

	if existingClinicID != 0 {
		_, err := database.Db.Exec("INSERT INTO vetdonor_pets (email, breed, petname, pettype, bloodtype, age, weight, vaccinations, photopath) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)", pet.Email, pet.Breed, pet.PetName, pet.PetType, pet.BloodType, pet.Age, pet.Weight, pet.Vaccinations, pet.PhotoPath)
		if err != nil {
			fmt.Println(err)
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update pet's questionnaire"})
			return
		}
		context.JSON(http.StatusOK, gin.H{"message": "Pet's questionnaire updated successfully"})
		return
	}
}

// GetClinicQuestionnaire Получить анкету клиники.
// @Summary Получить анкету клиники
// @Description Получить анкету клиники
// @Consumes application/json
// @Produce json
// @Param request body model.RequestQuestionnaire true "Запрос на получение анкеты клиники"
// @Success 200 {object} model.CodeResponse "Успешно получена анкета"
// @Failure 400 {object} model.ErrorResponse "Не удалось получить анкету"
// @Tags User
// @Router /v1/get_questionnaire [post]
func GetClinicQuestionnaire(context *gin.Context) {
	var request model.RequestQuestionnaire
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Can't read the body"})
		return
	}

	var info model.ClinicQuestionnaire
	var bloodTypesIncludedStr, bloodTypesNotIncludedStr string
	err := database.Db.QueryRow("SELECT email, name, address, array_to_string(bloodtypesincluded, ','), array_to_string(bloodtypesnotincluded, ','), workhours, contacts FROM vetdonor_clinic WHERE email = $1", request.Email).Scan(&info.Email, &info.Name, &info.Address, &bloodTypesIncludedStr, &bloodTypesNotIncludedStr, &info.WorkHours, &info.Contacts)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get clinic's questionnaire"})
		return
	}

	info.BloodTypesIncluded = strings.Split(bloodTypesIncludedStr, ",")
	info.BloodTypesNotIncluded = strings.Split(bloodTypesNotIncludedStr, ",")

	context.JSON(http.StatusOK, gin.H{"BloodTypesIncluded": info.BloodTypesIncluded, "BloodTypesNotIncluded": info.BloodTypesNotIncluded, "WorkHours": info.WorkHours, "Contacts": info.Contacts})
	return
}

// GetPetsQuestionnaire Получить анкету питомцев.
// @Summary Получить анкету питомцев
// @Description Получить анкету питомцев
// @Consumes application/json
// @Produce json
// @Param request body model.RequestQuestionnaire true "Запрос на получение анкеты питомцев"
// @Success 200 {object} model.CodeResponse "Успешно получена анкета"
// @Failure 400 {object} model.ErrorResponse "Не удалось получить анкету"
// @Tags User
// @Router /v1/get_pets_questionnaire [post]
func GetPetsQuestionnaire(context *gin.Context) {
	var request model.RequestQuestionnaire
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Can't read the body"})
		return
	}

	rows, err := database.Db.Query("SELECT breed, petname, pettype, bloodtype, age, weight, vaccinations, photopath FROM vetdonor_pets WHERE email = $1", request.Email)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get pets' questionnaire"})
		return
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var pets []model.PetQuestionnaire
	for rows.Next() {
		var pet model.PetQuestionnaire
		if err := rows.Scan(&pet.Breed, &pet.PetName, &pet.PetType, &pet.BloodType, &pet.Age, &pet.Weight, &pet.Vaccinations, &pet.PhotoPath); err != nil {
			fmt.Println(err)
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan pets' questionnaire"})
			return
		}
		pets = append(pets, pet)
	}
	if err := rows.Err(); err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error iterating over pets' questionnaire"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"pets": pets})
	return
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
	var count int
	err := database.Db.QueryRow("SELECT COUNT(*) FROM vetdonor_user_other_info WHERE email = $1", request.Email).Scan(&count)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check user's other info"})
		return
	}

	if count > 0 {
		_, err := database.Db.Exec("UPDATE vetdonor_user_other_info SET name = $1, city = $2, phone = $3, telegram = $4, path = $5 WHERE email = $6",
			request.Name, request.City, request.Phone, request.Telegram, request.Path, request.Email)
		if err != nil {
			fmt.Println(err)
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user's other info"})
			return
		}
	} else {
		_, err := database.Db.Exec("INSERT INTO vetdonor_user_other_info (email, name, city, phone, telegram, path) VALUES ($1, $2, $3, $4, $5, $6)",
			request.Email, request.Name, request.City, request.Phone, request.Telegram, request.Path)
		if err != nil {
			fmt.Println(err)
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert user's other info"})
			return
		}
	}

	context.JSON(http.StatusOK, gin.H{"message": "User's other info has been updated/inserted successfully"})
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
	err := database.Db.QueryRow("SELECT city, phone, telegram, path FROM vetdonor_user_other_info WHERE email = $1", request.Email).Scan(&form.City, &form.Phone, &form.Telegram, &form.Path)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user's other info"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"name": form.Name, "surname": form.Surname, "patronymic": form.Patronymic, "age": form.Age, "gender": form.Gender, "about": form.About})
}
