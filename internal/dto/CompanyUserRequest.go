package dto

type SetCompanyUserPasswordParams struct {
	Password string `json:"password" binding:"required"`
}

type CompanyInfo struct {
	ID   int    `json:"companyId"`
	Name string `json:"company"`
}
