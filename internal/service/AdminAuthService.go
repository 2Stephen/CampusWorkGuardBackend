package service

import (
	"CampusWorkGuardBackend/internal/dto"
	"CampusWorkGuardBackend/internal/repository"
	"CampusWorkGuardBackend/internal/utils"
	"errors"
)

func AdminLoginService(params *dto.AdminLoginRequest) (string, error) {
	user, err := repository.GetAdminUserByName(params.Name)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("用户登录失败，检查用户名或密码是否正确")
	}
	// 哈希验证
	ok, err := utils.VerifyPassword(params.Password, user.Password)
	if err != nil {
		return "", errors.New("用户登录失败" + err.Error())
	}
	if !ok {
		return "", errors.New("用户登录失败，检查用户名或密码是否正确")
	}
	// 生成JWT token
	token, err := utils.GenerateJWTToken(int(user.ID), user.Name, "admin")
	if err != nil {
		return "", errors.New("生成登录令牌失败")
	}
	return token, nil
}
