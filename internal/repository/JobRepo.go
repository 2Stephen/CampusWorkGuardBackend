package repository

import (
	"CampusWorkGuardBackend/internal/dto"
	"CampusWorkGuardBackend/internal/initialize"
	"CampusWorkGuardBackend/internal/model"
)

func CreateJobInfo(info *model.JobInfo) error {
	return initialize.DB.Create(info).Error
}

func GetJobByID(ID int) (model.JobInfo, error) {
	var job model.JobInfo
	err := initialize.DB.Where("id = ?", ID).Find(&job).Error
	return job, err
}

func GetJobsByCompanyID(companyID string, params dto.GetCompanyUserJobListParams) ([]model.JobInfo, error) {
	var jobs []model.JobInfo

	db := initialize.DB.Model(&model.JobInfo{}).
		Where("company_id = ?", companyID)

	// ===== 条件查询（为空不加）=====
	if params.Name != "" {
		db = db.Where("name LIKE ?", "%"+params.Name+"%")
	}

	if params.Type != "" {
		db = db.Where("type = ?", params.Type)
	}

	if params.Status != "" {
		db = db.Where("status = ?", params.Status)
	}

	// ===== 分页 =====
	page := params.Page
	pageSize := params.PageSize
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	err := db.
		Order("created_at DESC").
		Limit(pageSize).
		Offset(offset).
		Find(&jobs).Error

	return jobs, err
}
