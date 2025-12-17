package dto

type AdminLoginRequest struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AdminEmailLoginRequest struct {
	Email string `json:"email" binding:"required,email"`
	Code  string `json:"code" binding:"required"`
}

type SetAdminPasswordRequest struct {
	Password string `json:"password" binding:"required"`
}
