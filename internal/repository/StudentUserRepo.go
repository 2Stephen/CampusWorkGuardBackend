package repository

import (
	"CampusWorkGuardBackend/internal/initialize"
	"CampusWorkGuardBackend/internal/model"
	"log"
)

func SaveStudentUserPassword(hashedPassword string, userId string) error {
	err := initialize.DB.Model(&model.StudentUser{}).Where("id = ?", userId).Update("password", hashedPassword).Error
	if err != nil {
		log.Println("Error saving student user password to database:", err)
		return err
	}
	return nil
}

func GetStudentUserInfoById(userId int) (*model.StudentUser, error) {
	var user model.StudentUser
	err := initialize.DB.Where("id = ?", userId).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetStudentUserByID(id int64) *model.StudentUser {
	var user model.StudentUser
	err := initialize.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil
	}
	return &user
}
