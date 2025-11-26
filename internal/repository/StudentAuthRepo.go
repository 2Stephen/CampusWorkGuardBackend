package repository

import (
	"CampusWorkGuardBackend/internal/initialize"
	"CampusWorkGuardBackend/internal/model"
	"log"
)

func GetSchoolList(search string) ([]model.School, error) {
	var schools []model.School
	query := initialize.DB.Model(&model.School{})
	if search != "" {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}
	query = query.Limit(15)
	err := query.Find(&schools).Error
	if err != nil {
		log.Println("Error retrieving school list from database:", err)
		return nil, err
	}
	return schools, nil
}

func CreateCHSIStudentInfo(info *model.CHSIStudentInfo) error {
	err := initialize.DB.Create(info).Error
	if err != nil {
		log.Println("Error saving student info to database:", err)
	}
	return err
}
