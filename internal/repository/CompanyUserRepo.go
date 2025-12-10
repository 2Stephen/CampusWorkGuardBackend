package repository

import (
	"CampusWorkGuardBackend/internal/initialize"
	"CampusWorkGuardBackend/internal/model"
)

func GetCompanyUserById(userId int) (*model.CompanyUser, error) {
	var user model.CompanyUser
	err := initialize.DB.Where("id = ?", userId).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
