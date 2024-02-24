package model

import (
	"encoding/json"
)

type ClinicQuestionnaire struct {
	Email                 string   `json:"email"`
	Name                  string   `json:"name"`
	Address               string   `json:"address"`
	BloodTypesIncluded    []string `json:"BloodTypesIncluded"`
	BloodTypesNotIncluded []string `json:"bloodTypesNotIncluded"`
	WorkHours             string   `json:"workHours"`
	Contacts              string   `json:"contacts"`
}

type PetQuestionnaire struct {
	PetType      string `json:"petType"`
	Breed        string `json:"breed"`
	PetName      string `json:"petName"`
	BloodType    string `json:"bloodType"`
	Age          string `json:"age"`
	Weight       string `json:"weight"`
	Vaccinations string `json:"vaccinations"`
	PhotoPath    string `json:"photoPath"`
}

type OneOfUserClinicQuestionnaireResponse struct {
	ClinicQuestionnaire
	PetQuestionnaire
}

type CreateUserQuestionnaire struct {
	Email        string `json:"email" binding:"required"`
	PetType      string `json:"petType" binding:"required"`
	Breed        string `json:"breed" binding:"required"`
	PetName      string `json:"petName" binding:"required"`
	BloodType    string `json:"bloodType" binding:"required"`
	Age          string `json:"age" binding:"required"`
	Weight       string `json:"weight" binding:"required"`
	Vaccinations string `json:"vaccinations" binding:"required"`
	PhotoPath    string `json:"photoPath" binding:"required"`
}

type CreateClinicQuestionnaire struct {
	Email                 string   `json:"email" binding:"required"`
	BloodTypesIncluded    []string `json:"BloodTypesIncluded"`
	BloodTypesNotIncluded []string `json:"bloodTypesNotIncluded"`
	WorkHours             string   `json:"workHours" binding:"required"`
	Contacts              string   `json:"contacts" binding:"required"`
}

type CreateUserOtherInfo struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	City     string `json:"city"`
	Phone    string `json:"phone"`
	Telegram string `json:"telegram"`
	Path     string `json:"path"`
}

type RequestUserOtherInfo struct {
	Email string `json:"email"`
}

type RequestQuestionnaire struct {
	Email string `json:"Email"`
}

type ResponseUserOtherInfo struct {
	City     string `json:"city"`
	Phone    string `json:"phone"`
	Telegram string `json:"telegram"`
	Path     string `json:"path"`
}

type QuestionnaireRequest struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}
