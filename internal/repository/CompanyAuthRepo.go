package repository

import (
	"CampusWorkGuardBackend/internal/initialize"
	"CampusWorkGuardBackend/internal/model"
)

func CreateCompanyUser(name, email, company, licenseURL, socialCode string) (int64, error) {
	user := model.CompanyUser{
		Name:         name,
		Email:        email,
		Company:      company,
		LicenseURL:   licenseURL,
		SocialCode:   socialCode,
		VerifyStatus: "验证中",
	}

	if err := initialize.DB.Create(&user).Error; err != nil {
		return 0, err
	}

	// Create 成功后，ID 会自动回填到 user.ID
	return user.ID, nil
}
