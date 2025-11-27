package repository

import (
	"CampusWorkGuardBackend/internal/initialize"
	"CampusWorkGuardBackend/internal/model"
)

func IsUserExistByEmail(email string) bool {
	// 查询数据库，检查是否存在该邮箱的用户
	query := initialize.DB.Model(&model.StudentUser{}).Where("email = ?", email)
	var count int64
	query.Count(&count)
	return count > 0
}
