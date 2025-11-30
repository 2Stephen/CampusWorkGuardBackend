package service

import (
	"CampusWorkGuardBackend/internal/dto"
	"CampusWorkGuardBackend/internal/repository"
	"CampusWorkGuardBackend/internal/utils"
	"errors"
)

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
