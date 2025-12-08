package repository

import (
	"CampusWorkGuardBackend/internal/initialize"
	"CampusWorkGuardBackend/internal/model"
	"errors"
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
	err := initialize.DB.
		Where("email = ? OR social_code = ?", email, socialCode).
		First(&model.CompanyUser{}).Error
	if err == nil {
		return 0, errors.New("重复注册") // 已存在相同邮箱或社会信用代码的用户
	}
	if err := initialize.DB.Create(&user).Error; err != nil {
		return 0, err
	}

	// Create 成功后，ID 会自动回填到 user.ID
	return user.ID, nil
}

func GetCompanyUserByEmail(email string) *model.CompanyUser {
	var user model.CompanyUser
	err := initialize.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil
	}
	return &user
}
