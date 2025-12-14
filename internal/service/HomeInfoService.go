package service

import (
	"CampusWorkGuardBackend/internal/model"
	"CampusWorkGuardBackend/internal/repository"
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
