package dto

type CompanyRegisterRequest struct {
	Name       string `json:"name" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Company    string `json:"company" binding:"required"`
	LicenseURL string `json:"licenseUrl" binding:"required"`
	Code       string `json:"code" binding:"required"`
	SocialCode string `json:"socialCode" binding:"required"`
}
