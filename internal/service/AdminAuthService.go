package service

import (
	"CampusWorkGuardBackend/internal/dto"
	"CampusWorkGuardBackend/internal/repository"
	"CampusWorkGuardBackend/internal/utils"
	"errors"
	"github.com/redis/go-redis/v9"
	"log"
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
	token, err := utils.GenerateJWTToken(int(user.ID), user.Email, "admin")
	if err != nil {
		return "", errors.New("生成登录令牌失败")
	}
	return token, nil
}

func AdminEmailLoginService(params *dto.AdminEmailLoginRequest) (string, error) {
	// 验证邮箱验证码
	realCode, err := utils.RedisGet("login_code:" + params.Email)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", errors.New("邮箱验证码已过期，请重新获取")
		}
		return "", errors.New("获取邮箱验证码失败")
	}
	if realCode != params.Code {
		return "", errors.New("邮箱验证码有误")
	}
	// 清除redis缓存验证码
	err = utils.RedisDel("login_code:" + params.Email)
	if err != nil {
		log.Println("删除邮箱验证码缓存失败", err)
	}
	user, err := repository.GetAdminUserByEmail(params.Email)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("用户登录失败，检查邮箱或验证码是否正确")
	}
	// 生成JWT token
	token, err := utils.GenerateJWTToken(user.ID, user.Email, "admin")
	if err != nil {
		return "", errors.New("生成登录令牌失败")
	}
	return token, nil
}

func SetAdminPasswordService(params *dto.SetAdminPasswordRequest, userID int) error {
	hashedPassword, err := utils.HashPassword(params.Password)
	if err != nil {
		return errors.New("密码加密失败")
	}
	err = repository.UpdateAdminUserPasswordByID(int64(userID), hashedPassword)
	if err != nil {
		return errors.New("设置密码失败")
	}
	return nil
}
