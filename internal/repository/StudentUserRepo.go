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
