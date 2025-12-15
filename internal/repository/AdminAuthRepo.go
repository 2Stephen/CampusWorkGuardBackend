package repository

import (
	"CampusWorkGuardBackend/internal/initialize"
	"CampusWorkGuardBackend/internal/model"
	"log"
)

func GetAdminUserByName(name string) (*model.AdminUser, error) {
	var user model.AdminUser
	err := initialize.DB.Where("name = ?", name).First(&user).Error
	if err != nil {
		log.Println("GetAdminUserByName error:", err)
		return nil, err
	}
	return &user, nil
}

func GetAdminUserByEmail(email string) (*model.AdminUser, error) {
	var user model.AdminUser
	err := initialize.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		log.Println("GetAdminUserByEmail error:", err)
		return nil, err
	}
	return &user, nil
}

func UpdateAdminUserPasswordByID(id int64, hashedPassword string) error {
	return initialize.DB.Model(&model.AdminUser{}).Where("id = ?", id).Update("password", hashedPassword).Error
}
