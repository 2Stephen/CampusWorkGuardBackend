package repository

import (
	"CampusWorkGuardBackend/internal/initialize"
	"CampusWorkGuardBackend/internal/model"
)

func SaveComplaintRecord(complaint *model.ComplaintRecord) error {
	return initialize.DB.Create(complaint).Error
}
