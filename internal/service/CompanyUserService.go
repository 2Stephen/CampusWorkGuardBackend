package service

import (
	"CampusWorkGuardBackend/internal/dto"
	"CampusWorkGuardBackend/internal/repository"
	"CampusWorkGuardBackend/internal/utils"
	"errors"
)

func SetCompanyUserPassword(params dto.SetCompanyUserPasswordParams, userID string) error {
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
	return repository.SaveCompanyUserPassword(hashedPassword, userID)
}

func DeleteCompanyUserService(id int) error {
	return repository.DeleteCompanyUserByID(int64(id))
}
