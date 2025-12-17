package model

type CompanyList struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Company      string `json:"company"`
	Email        string `json:"email"`
	SocialCode   string `json:"socialCode"`
	LicenseUrl   string `json:"licenseUrl"`
	VerifyStatus string `json:"verifyStatus"`
}
