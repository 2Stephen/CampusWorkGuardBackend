package repository

import (
	"CampusWorkGuardBackend/internal/initialize"
	"CampusWorkGuardBackend/internal/model"
)

func CreateJobInfo(info *model.JobInfo) error {
	return initialize.DB.Create(info).Error
}
