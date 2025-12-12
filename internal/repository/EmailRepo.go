package repository

import (
	"CampusWorkGuardBackend/internal/initialize"
	"CampusWorkGuardBackend/internal/model"
)

func IsUserExistByEmail(email string) bool {
	// 查询数据库，检查是否存在该邮箱的用户
	query_student := initialize.DB.Model(&model.StudentUser{}).Where("email = ?", email)
	var studentCount int64
	query_student.Count(&studentCount)
	query_company := initialize.DB.Model(&model.CompanyUser{}).Where("email = ?", email)
	var companyCount int64
	query_company.Count(&companyCount)
	total := studentCount + companyCount
	return total > 0
}
