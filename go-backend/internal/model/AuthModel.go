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

type ForgotRequest struct {
	Email string `json:"email" binding:"required"`
}

type NewPassword struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserInfo struct {
	Email string
	Name  string
}

type ClinicInfo struct {
	Email   string
	Name    string
	Address string
}
