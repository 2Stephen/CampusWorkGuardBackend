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

func GetJobsByCompanyID(
	companyID string,
	params dto.GetCompanyUserJobListParams,
) ([]model.JobInfo, int64, error) {

	var (
		jobs  []model.JobInfo
		total int64
	)

	// 基础查询
	db := initialize.DB.Model(&model.JobInfo{}).
		Where("company_id = ?", companyID)

	// ===== 条件查询（空字符串不参与）=====
	if params.Name != "" {
		db = db.Where("name LIKE ?", "%"+params.Name+"%")
	}

	if params.Type != "" {
		db = db.Where("type = ?", params.Type)
	}

	if params.Status != "" {
		db = db.Where("status = ?", params.Status)
	}

	// ===== 先查 total =====
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// ===== 分页参数 =====
	page := params.Page
	pageSize := params.PageSize

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	// ===== 查询列表 =====
	if err := db.
		Order("created_at DESC").
		Limit(pageSize).
		Offset(offset).
		Find(&jobs).Error; err != nil {
		return nil, 0, err
	}
	return jobs, total, nil
}

func UpdateJobInfo(info *model.JobInfo) error {
	return initialize.DB.Model(&model.JobInfo{}).Where("id = ?", info.ID).Updates(info).Error
}

func DeleteJobByID(id int64) error {
	return initialize.DB.Delete(&model.JobInfo{}, id).Error
}
