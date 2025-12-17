package repository

import (
	"CampusWorkGuardBackend/internal/initialize"
	"CampusWorkGuardBackend/internal/model"
)

func GetAdminCompanyList(page int, pageSize int, search string, status string) ([]model.CompanyUser, int64, error) {
	var (
		companyUsers []model.CompanyUser
		total        int64
	)
	db := initialize.DB.Model(&model.CompanyUser{})

	// ===== 条件查询（空字符串不参与）=====
	if search != "" {
		db = db.Where("company LIKE ?", "%"+search+"%")
	}
	if status != "" {
		db = db.Where("verify_status = ?", status)
	}
	// ===== 先查 total =====
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	// ===== 分页参数 =====
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize
	// ===== 查询数据 =====
	if err := db.Limit(pageSize).Offset(offset).Find(&companyUsers).Error; err != nil {
		return nil, 0, err
	}
	return companyUsers, total, nil
}

func ReviewCompany(id int, status string, failInfo string) error {
	return initialize.DB.Model(&model.CompanyUser{}).Where("id = ?", id).Updates(map[string]interface{}{
		"verify_status": status,
		"fail_info":     failInfo,
	}).Error
}
