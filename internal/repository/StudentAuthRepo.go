package repository

import (
	"CampusWorkGuardBackend/internal/initialize"
	"CampusWorkGuardBackend/internal/model"
	"log"
)

func CreateCHSIStudentInfo(info *model.CHSIStudentInfo) error {
	err := initialize.DB.Create(info).Error
	if err != nil {
		log.Println("Error saving student info to database:", err)
	}
	return err
}
