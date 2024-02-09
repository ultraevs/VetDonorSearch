package model

type UserCreateRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name"  binding:"required"`
}

type ClinicCreateRequest struct {
	Email    string `json:"email" binding:"required"`
	Address  string `json:"address" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"Name"  binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"pass"`
}

type ClinicResponse struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Name     string `json:"name"`
	Password string `json:"pass"`
}
