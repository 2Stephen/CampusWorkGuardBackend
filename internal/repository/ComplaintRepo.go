package repository

import (
	"CampusWorkGuardBackend/internal/initialize"
	"CampusWorkGuardBackend/internal/model"
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
