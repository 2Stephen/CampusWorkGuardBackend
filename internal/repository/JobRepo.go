package repository

import (
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
