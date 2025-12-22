package service

import (
	"CampusWorkGuardBackend/internal/model"
	"CampusWorkGuardBackend/internal/repository"
	"log"
)

func GetHomeStaticInfo(userId int, role string, email string) (*model.HomeStaticInfo, error) {
	var (
		info *model.HomeStaticInfo
	)
	if role == "student" {
		studentInfo, err := repository.GetStudentUserInfoById(userId)
		if err != nil {
			return nil, err
		}

		info = &model.HomeStaticInfo{
			AvatarURL: studentInfo.AvatarURL,
		}
		chsiInfo, err := repository.GetCHSIStudentInfoByEmail(email)
		if err != nil {
			return nil, err
		}
		info.Name = chsiInfo.Name
		info.VerifyStatus = "verified"
		info.FailInfo = ""
	} else if role == "company" {
		companyInfo, err := repository.GetCompanyUserById(userId)
		if err != nil {
			return nil, err
		}
		info = &model.HomeStaticInfo{
			AvatarURL:    companyInfo.AvatarURL,
			Name:         companyInfo.Company,
			VerifyStatus: companyInfo.VerifyStatus,
			FailInfo:     companyInfo.FailInfo,
		}
	} else if role == "admin" {
		adminInfo, err := repository.GetAdminUserByEmail(email)
		if err != nil {
			return nil, err
		}
		info = &model.HomeStaticInfo{
			AvatarURL: adminInfo.AvatarURL,
			Name:      adminInfo.Name,
		}
	}
	info.Role = role
	return info, nil
}

func UploadAvatarService(filePath string, userId int, role string) error {
	if role == "student" {
		return repository.UpdateStudentUserAvatarURL(filePath, userId)
	} else if role == "company" {
		return repository.UpdateCompanyUserAvatarURL(filePath, userId)
	}
	return nil
}

func GetTop5MajorJobsService() ([]model.TopMajorJob, error) {
	majorJobs, err := repository.GetTop5MajorJobs()
	if err != nil {
		return nil, err
	}
	return majorJobs, nil
}

func GetJobTypesService() ([]model.JobType, error) {
	jobTypes, err := repository.GetJobTypes()
	if err != nil {
		return nil, err
	}
	return jobTypes, nil
}

func GetAverageSalariesByMajorService() ([]model.AverageSalaryByMajor, error) {
	avgSalaries, err := repository.GetAverageSalariesByMajor()
	if err != nil {
		log.Println("Error in service layer fetching average salaries by major:", err)
		return nil, err
	}
	return avgSalaries, nil
}
