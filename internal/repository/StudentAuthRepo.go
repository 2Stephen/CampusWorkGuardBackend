package repository

import (
	"CampusWorkGuardBackend/internal/initialize"
	"CampusWorkGuardBackend/internal/model"
	"errors"
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

func GetSchoolId(name string) (int, error) {
	var res int
	err := initialize.DB.Model(&model.School{}).Select("id").Where("name = ?", name).Scan(&res).Error
	if err != nil {
		log.Println("Error retrieving school ID from database:", err)
		return 0, err
	}
	return res, nil
}

func CreateStudentUser(user model.StudentUser) error {
	// 检查是否已存在相同学生ID或邮箱的用户
	var existingUser model.StudentUser
	err := initialize.DB.Where("email = ?", user.Email).First(&existingUser).Error
	if err == nil {
		log.Println("Student user with the same Email already exists")
		return errors.New("该邮箱已经被注册")
	}
	err = initialize.DB.Where("school_id = ? and student_id = ?", user.SchoolId, user.StudentId).First(&existingUser).Error
	if err == nil {
		log.Println("Student user with the same Email already exists")
		return errors.New("该学校+学号已被注册，请勿重复注册")
	}
	err = initialize.DB.Create(&user).Error
	if err != nil {
		log.Println("Error saving student user to database:", err)
	}
	return err
}

func CreateCHSIStudentInfo(info *model.CHSIStudentInfo) error {
	var existingCHSIInfo model.CHSIStudentInfo
	err := initialize.DB.Where("email = ?", info.Email).First(&existingCHSIInfo).Error
	if err == nil {
		log.Println("Student user with the same Email already exists")
		return errors.New("该邮箱已经被注册")
	}
	err = initialize.DB.Where("school = ? and student_id = ?", info.School, info.StudentID).First(&existingCHSIInfo).Error
	if err == nil {
		log.Println("Student user with the same Email already exists")
		return errors.New("该学校+学号已被注册，请勿重复注册")
	}
	err = initialize.DB.Create(info).Error
	if err != nil {
		log.Println("Error saving student info to database:", err)
	}
	return err
}
