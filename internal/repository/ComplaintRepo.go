package repository

import (
	"CampusWorkGuardBackend/internal/dto"
	"CampusWorkGuardBackend/internal/initialize"
	"CampusWorkGuardBackend/internal/model"
	"log"
)

func SaveComplaintRecord(complaint *model.ComplaintRecord) error {
	return initialize.DB.Create(complaint).Error
}

func GetComplaintRecordByID(id int) (*model.ComplaintRecord, error) {
	var complaint model.ComplaintRecord
	err := initialize.DB.Where("id = ?", id).First(&complaint).Error
	if err != nil {
		return nil, err
	}
	return &complaint, nil
}

func DeleteComplaintRecord(complaintID int) error {
	return initialize.DB.Delete(&model.ComplaintRecord{}, complaintID).Error
}

func UpdateComplaintRecordCompanyDefense(complaintID int, defense string) error {
	return initialize.DB.Model(&model.ComplaintRecord{}).Where("id = ?", complaintID).Update("company_defense", defense).Error
}

func GetComplaintRecordsByStudentID(params dto.GetComplaintListParams, studentID int) ([]model.ComplaintRecord, int, error) {
	var complaints []model.ComplaintRecord
	var total int64

	query := initialize.DB.Where("student_id = ?", studentID)

	if params.Search != "" {
		query = query.Where("Title LIKE ?", "%"+params.Search+"%")
	}
	err := query.Model(&model.ComplaintRecord{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = query.Offset((params.Page - 1) * params.PageSize).Limit(params.PageSize).Find(&complaints).Error
	if err != nil {
		return nil, 0, err
	}
	return complaints, int(total), nil
}

func GetComplaintRecordsByCompanyID(params dto.GetComplaintListParams, companyID int) ([]model.ComplaintRecord, int, error) {
	var complaints []model.ComplaintRecord
	var total int64

	query := initialize.DB.Where("company_id = ?", companyID)
	if params.Search != "" {
		query = query.Where("Title LIKE ?", "%"+params.Search+"%")
	}
	err := query.Model(&model.ComplaintRecord{}).Count(&total).Error
	if err != nil {
		log.Println("Error counting complaint records for company:", err)
		return nil, 0, err
	}
	err = query.Offset((params.Page - 1) * params.PageSize).Limit(params.PageSize).Find(&complaints).Error
	if err != nil {
		log.Println("Error retrieving complaint records for company:", err)
		return nil, 0, err
	}
	return complaints, int(total), nil
}

func GetAllComplaintRecords(params dto.GetComplaintListParams) ([]model.ComplaintRecord, int, error) {
	var complaints []model.ComplaintRecord
	var total int64

	query := initialize.DB.Model(&model.ComplaintRecord{})
	if params.Search != "" {
		query = query.Where("Title LIKE ?", "%"+params.Search+"%")
	}
	err := query.Count(&total).Error
	if err != nil {
		log.Println("Error counting all complaint records:", err)
		return nil, 0, err
	}
	err = query.Offset((params.Page - 1) * params.PageSize).Limit(params.PageSize).Find(&complaints).Error
	if err != nil {
		log.Println("Error retrieving all complaint records:", err)
		return nil, 0, err
	}
	return complaints, int(total), nil
}

func UpdateComplaintRecordResultInfo(complaintID int, resultInfo string) error {
	return initialize.DB.Model(&model.ComplaintRecord{}).Where("id = ?", complaintID).Update("result_info", resultInfo).Error
}
