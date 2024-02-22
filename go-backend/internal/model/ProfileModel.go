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

type UserQuestionnaire struct {
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
	Breed     string `json:"breed" binding:"required"`
	PetName   string `json:"petName" binding:"required"`
	BloodType string `json:"bloodType" binding:"required"`
	Age       string `json:"age" binding:"required"`
}

type CreateClinicQuestionnaire struct {
	Email                 string   `json:"email" binding:"required"`
	BloodTypesIncluded    []string `json:"BloodTypesIncluded"`
	BloodTypesNotIncluded []string `json:"bloodTypesNotIncluded"`
	WorkHours             string   `json:"workHours" binding:"required"`
	Contacts              string   `json:"contacts" binding:"required"`
}

type CreateUserOtherInfo struct {
	Email      string `json:"email"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Age        string `json:"age"`
	Gender     string `json:"gender"`
	About      string `json:"about"`
	City       string `json:"city"`
}

type RequestUserOtherInfo struct {
	Email string `json:"email"`
}

type RequestQuestionnaire struct {
	Email string `json:"Email"`
}

type ResponseUserOtherInfo struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Age        string `json:"age"`
	Gender     string `json:"gender"`
	About      string `json:"about"`
	City       string `json:"city"`
}

type QuestionnaireRequest struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}
