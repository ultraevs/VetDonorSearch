package model

import (
	"encoding/json"
)

type ClinicQuestionnaire struct {
	BloodTypes map[string]interface{} `json:"bloodTypes" binding:"required"`
	WorkHours  string                 `json:"workHours" binding:"required"`
	Contacts   string                 `json:"contacts" binding:"required"`
}

type UserQuestionnaire struct {
	City      string `json:"city" binding:"required"`
	Surname   string `json:"surname" binding:"required"`
	Breed     string `json:"breed" binding:"required"`
	PetName   string `json:"petName" binding:"required"`
	BloodType string `json:"bloodType" binding:"required"`
	Age       string `json:"age" binding:"required"`
}

type OneOfUserClinicQuestionnaireResponse struct {
	ClinicQuestionnaire
	UserQuestionnaire
}

type CreateUserQuestionnaire struct {
	Email     string `json:"email" binding:"required"`
	City      string `json:"city" binding:"required"`
	Surname   string `json:"surname" binding:"required"`
	Breed     string `json:"breed" binding:"required"`
	PetName   string `json:"petName" binding:"required"`
	BloodType string `json:"bloodType" binding:"required"`
	Age       string `json:"age" binding:"required"`
}

type CreateClinicQuestionnaire struct {
	Email      string                 `json:"email" binding:"required"`
	BloodTypes map[string]interface{} `json:"bloodTypes" binding:"required"`
	WorkHours  string                 `json:"workHours" binding:"required"`
	Contacts   string                 `json:"contacts" binding:"required"`
}

type CreateUserOtherInfo struct {
	Email      string `json:"email"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Age        string `json:"age"`
	Gender     string `json:"gender"`
	About      string `json:"about"`
}

type RequestUserOtherInfo struct {
	Email string `json:"email"`
}

type ResponseUserOtherInfo struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Age        string `json:"age"`
	Gender     string `json:"gender"`
	About      string `json:"about"`
}

type QuestionnaireRequest struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}
