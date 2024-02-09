package controller

import (
	"app/internal/database"
	"app/internal/model"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

// UserCreate создает нового пользователя.
// @Summary Создать нового пользователя
// @Description Создает нового пользователя с предоставленным email, паролем и именем.
// @Accept json
// @Produce json
// @Param request body model.UserCreateRequest true "Запрос на создание пользователя"
// @Success 200 {object} model.CodeResponse "Пользователь успешно создан"
// @Failure 400 {object} model.ErrorResponse "Не удалось создать пользователя"
// @Tags Auth
// @Router /v1/user_create [post]
func UserCreate(context *gin.Context) {
	var body struct {
		Email    string
		Password string
		Name     string
	}
	if context.Bind(&body) != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}
	hashPass, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
	}
	_, err = database.Db.Exec("INSERT INTO vetdonor_users (email, password, name) VALUES ($1, $2, $3)", body.Email, string(hashPass), body.Name)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

// ClinicCreate создает новый аккаунт клиники.
// @Summary Создать новоый аккаунт клиники
// @Description Создает новый аккаунт клиники с предоставленным email, паролем, именем и адрессом.
// @Accept json
// @Produce json
// @Param request body model.ClinicCreateRequest true "Запрос на создание пользователя"
// @Success 200 {object} model.CodeResponse "Аккаунт клиники успешно создан"
// @Failure 400 {object} model.ErrorResponse "Не удалось создать аккаунт клиники"
// @Tags Auth
// @Router /v1/clinic_create [post]
func ClinicCreate(context *gin.Context) {
	var body struct {
		Email    string
		Address  string
		Password string
		Name     string
	}
	if context.Bind(&body) != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}
	hashPass, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
	}
	_, err = database.Db.Exec("INSERT INTO vetdonor_clinic (email, password, name, address) VALUES ($1, $2, $3, $4)", body.Email, string(hashPass), body.Name, body.Address)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Clinic"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Clinic created successfully"})
}

// Login Вход в аккаунт.
// @Summary Логин
// @Description Авторизует пользователя с предоставленным email и паролем.
// @Accept json
// @Produce json
// @Param request body model.LoginRequest true "Запрос на авторизацию пользователя"
// @Success 200 {object} model.CodeResponse "Пользователь авторизован"
// @Failure 400 {object} model.ErrorResponse "Не удалось авторизовать пользователя"
// @Tags Auth
// @Router /v1/login [post]
func Login(context *gin.Context) {
	var body struct {
		Email    string
		Password string
	}
	if context.Bind(&body) != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}

	var form model.LoginRequest
	err := database.Db.QueryRow("SELECT email, password FROM vetdonor_users WHERE email = $1", body.Email).Scan(&form.Email, &form.Password)
	if err == nil {
		context.JSON(http.StatusOK, gin.H{"user": form})
		return
	}
	if !errors.Is(err, sql.ErrNoRows) {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	err = database.Db.QueryRow("SELECT email, password FROM vetdonor_clinic WHERE email = $1", body.Email).Scan(&form.Email, &form.Password)
	if err == nil {
		context.JSON(http.StatusOK, gin.H{"clinic": form})
		return
	}
	if errors.Is(err, sql.ErrNoRows) {
		context.JSON(http.StatusNotFound, gin.H{"error": "User is not found"})
		return
	}
	context.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})

	err = bcrypt.CompareHashAndPassword([]byte(form.Password), []byte(body.Password))
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong Password"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": form.Email,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		return
	}

	context.SetSameSite(http.SameSiteLaxMode)
	context.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	context.JSON(http.StatusOK, gin.H{"response": "success"})
}
