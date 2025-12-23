package repository

import (
	"CampusWorkGuardBackend/internal/initialize"
	"CampusWorkGuardBackend/internal/model"
	"log"
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

func GetJobTypes() ([]model.JobType, error) {
	var jobTypes []model.JobType
	err := initialize.DB.Raw("SELECT count(*) as value,job_infos.type FROM job_infos GROUP BY job_infos.type").Scan(&jobTypes).Error
	if err != nil {
		log.Println("Error fetching job types:", err)
		return nil, err
	}
	return jobTypes, nil
}

func GetComplaintTypes() ([]model.ComplaintType, error) {
	var complaintTypes []model.ComplaintType
	err := initialize.DB.Raw("SELECT count(*) as value,complaint_records.complaint_type as type FROM complaint_records GROUP BY type").Scan(&complaintTypes).Error
	if err != nil {
		log.Println("Error fetching complaint types:", err)
		return nil, err
	}
	return complaintTypes, nil
}

func GetAverageSalariesByMajor() ([]model.AverageSalaryByMajor, error) {
	// SELECT job_infos.major, AVG(
	//	CASE job_infos.salary_unit
	//		WHEN 'hour' THEN salary * 8 * 22
	//		WHEN 'day' THEN salary * 22
	//		WHEN 'month' THEN salary
	//		ELSE 0
	//	END
	//) AS value FROM job_infos
	//GROUP BY job_infos.major
	var avgSalaries []model.AverageSalaryByMajor
	err := initialize.DB.Raw(`
		SELECT job_infos.major, AVG(
			CASE job_infos.salary_unit
				WHEN 'hour' THEN salary * 8 * 22
				WHEN 'day' THEN salary * 22
				WHEN 'month' THEN salary
				ELSE 0
			END
		) AS value FROM job_infos
		GROUP BY job_infos.major
		ORDER BY value DESC`).Scan(&avgSalaries).Error
	if err != nil {
		log.Println("Error fetching average salaries by major:", err)
		return nil, err
	}
	return avgSalaries, nil
}
