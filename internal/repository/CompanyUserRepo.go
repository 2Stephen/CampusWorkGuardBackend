package repository

import (
	"CampusWorkGuardBackend/internal/initialize"
	"CampusWorkGuardBackend/internal/model"
	"log"
)

func GetCompanyUserById(userId int) (*model.CompanyUser, error) {
	var user model.CompanyUser
	err := initialize.DB.Where("id = ?", userId).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func SaveCompanyUserPassword(hashedPassword string, userId string) error {
	err := initialize.DB.Model(&model.CompanyUser{}).Where("id = ?", userId).Update("password", hashedPassword).Error
	if err != nil {
		log.Println("Error saving student user password to database:", err)
		return err
	}
	return nil
}
