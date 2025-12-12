package service

import (
	"CampusWorkGuardBackend/internal/dto"
	"CampusWorkGuardBackend/internal/repository"
	"CampusWorkGuardBackend/internal/utils"
	"errors"
	"log"
)

type ProfileInfo struct {
	AvatarURL     string `json:"avatar_url"`
	StudentID     string `json:"student_id"`
	Email         string `json:"email"`
	Name          string `json:"name"`
	Gender        string `json:"gender"`
	Birthday      string `json:"birthday"`
	Nation        string `json:"nation"`
	School        string `json:"school"`
	Level         string `json:"level"`
	Major         string `json:"major"`
	Duration      string `json:"duration"`
	College       string `json:"college"`
	Department    string `json:"department"`
	Entrance_date string `json:"entrance_date"`
	Status        string `json:"status"`
	ExpectedGrad  string `json:"expected_grad"`
}

func SetStudentUserPassword(params dto.SetStudentUserPasswordParams, userId string) error {
	// 校验密码长度、复杂度等
	if len(params.Password) < 8 {
		return errors.New("密码长度不足，至少需要8位")
	}
	if !containsNumber(params.Password) || !containsLetter(params.Password) {
		return errors.New("密码必须包含字母和数字")
	}
	if len(params.Password) > 64 {
		return errors.New("密码长度过长，不能超过64位")
	}
	hashedPassword, err := utils.HashPassword(params.Password)
	if err != nil {
		return errors.New("密码加密失败")
	}
	// 调用repository层保存密码逻辑
	return repository.SaveStudentUserPassword(hashedPassword, userId)
}

func containsNumber(s string) bool {
	for _, char := range s {
		if char >= '0' && char <= '9' {
			return true
		}
	}
	return false
}

func containsLetter(s string) bool {
	for _, char := range s {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			return true
		}
	}
	return false
}

func GetStudentUserProfileInfoService(userID int) (*ProfileInfo, error) {
	user := repository.GetStudentUserByID(int64(userID))
	if user == nil {
		return nil, errors.New("用户不存在")
	}
	userInfo, err := repository.GetCHSIStudentInfoByEmail(user.Email)
	if err != nil {
		log.Println("Error retrieving student info from CHSI:", err)
		return nil, err
	}
	profileInfo := &ProfileInfo{
		AvatarURL:     user.AvatarURL,
		StudentID:     user.StudentId,
		Email:         user.Email,
		Name:          userInfo.Name,
		Gender:        userInfo.Gender,
		Birthday:      userInfo.Birthday,
		Nation:        userInfo.Nation,
		School:        userInfo.School,
		Level:         userInfo.Level,
		Major:         userInfo.Major,
		Duration:      userInfo.Duration,
		College:       userInfo.College,
		Department:    userInfo.Department,
		Entrance_date: userInfo.EntranceDate,
		Status:        userInfo.Status,
		ExpectedGrad:  userInfo.ExpectedGrad,
	}
	return profileInfo, nil
}
