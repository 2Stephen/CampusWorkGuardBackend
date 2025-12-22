package repository

import (
	"CampusWorkGuardBackend/internal/initialize"
	"CampusWorkGuardBackend/internal/model"
)

func GetTop5MajorJobs() ([]model.TopMajorJob, error) {
	// SELECT count(*) as value,job_infos.major FROM job_infos
	//	GROUP BY job_infos.major
	//	ORDER BY value DESC
	//	LIMIT 5
	var majorJobs []model.TopMajorJob
	err := initialize.DB.Raw("SELECT count(*) as value,job_infos.major FROM job_infos GROUP BY job_infos.major ORDER BY value DESC LIMIT 5").Scan(&majorJobs).Error
	if err != nil {
		return nil, err
	}
	return majorJobs, nil
}
