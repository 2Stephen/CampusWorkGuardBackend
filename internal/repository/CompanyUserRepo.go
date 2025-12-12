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
func UpdateStudentUserAvatarURL(filePath string, userId int) error {
	err := initialize.DB.Model(&model.StudentUser{}).Where("id = ?", userId).Update("avatar_url", filePath).Error
	if err != nil {
		log.Println("Error updating company user avatar URL in database:", err)
		return err
	}
	return nil
}

func UpdateCompanyUserAvatarURL(filePath string, userId int) error {
	err := initialize.DB.Model(&model.CompanyUser{}).Where("id = ?", userId).Update("avatar_url", filePath).Error
	if err != nil {
		log.Println("Error updating company user avatar URL in database:", err)
		return err
	}
	return nil
}
