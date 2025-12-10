package dto

type SetCompanyUserPasswordParams struct {
	Password string `json:"password" binding:"required"`
}
