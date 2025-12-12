package service

import (
	"CampusWorkGuardBackend/internal/dto"
	"CampusWorkGuardBackend/internal/repository"
	"CampusWorkGuardBackend/internal/utils"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

var allowedTypes = map[string]bool{
	".jpg":  true,
	".png":  true,
	".jpeg": true,
}

func SaveImage(file *multipart.FileHeader) (string, error) {
	if file.Size > 5*1024*1024 {
		return "", errors.New("文件大小超过5MB限制")
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !allowedTypes[ext] {
		return "", errors.New("不支持的文件类型")
	}

	uuidStr := uuid.New().String()
	subDir := filepath.Join("uploads", uuidStr[:2])
	err := os.MkdirAll(subDir, os.ModePerm)
	if err != nil {
		return "", err
	}

	savePath := filepath.Join(subDir, uuidStr+ext)
	if err := SaveUploadedFile(file, savePath); err != nil {
		return "", fmt.Errorf("保存文件失败: %v", err)
	}
	return filepath.Join(uuidStr[:2], uuidStr+ext), nil
}

func SaveUploadedFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer func(src multipart.File) {
		err := src.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(src)

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(out)

	_, err = io.Copy(out, src)
	return err
}

// RegisterCompanyService 公司注册服务，负责处理公司注册逻辑
func RegisterCompanyService(req *dto.CompanyRegisterRequest) (string, error) {
	// 验证邮箱验证码
	realCode, err := utils.RedisGet("register_code:" + req.Email)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", errors.New("邮箱验证码已过期，请重新获取")
		}
		return "", errors.New("获取邮箱验证码失败")
	}
	if realCode != req.Code {
		return "", errors.New("邮箱验证码有误")
	} else {
		// 清除redis缓存验证码
		err := utils.RedisDel("register_code:" + req.Email)
		if err != nil {
			log.Println("删除邮箱验证码缓存失败", err)
		}
	}
	// 保存公司信息到数据库
	id, err := repository.CreateCompanyUser(req.Name, req.Email, req.Company, req.LicenseURL, req.SocialCode)
	if err != nil {
		log.Println("Error saving company user to database:", err)
		return "", err
	}
	token, err := utils.GenerateJWTToken(int(id), req.Email, "company")
	if err != nil {
		return "", errors.New("生成登录令牌失败")
	}
	return token, nil
}

func CompanyLoginService(req *dto.CompanyLoginRequest) (string, error) {
	user := repository.GetCompanyUserByEmail(req.Email)
	if user == nil {
		return "", errors.New("用户登录失败，检查邮箱或密码是否正确")
	}
	if user.Password == "" {
		return "", errors.New("用户未设置密码，请使用邮箱验证登录后设置密码")
	}
	// 哈希验证
	ok, err := utils.VerifyPassword(req.Password, user.Password)
	if err != nil {
		return "", errors.New("用户登录失败" + err.Error())
	}
	if !ok {
		return "", errors.New("用户登录失败，检查邮箱或密码是否正确")
	}
	// 生成JWT token
	token, err := utils.GenerateJWTToken(int(user.ID), user.Email, "company")
	if err != nil {
		return "", errors.New("生成登录令牌失败")
	}
	return token, nil
}

func CompanyEmailLogin(params dto.CompanyEmailLoginRequest) (string, error) {
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
	user := repository.GetCompanyUserByEmail(params.Email)
	if user == nil {
		return "", errors.New("用户登录失败，检查邮箱或验证码是否正确")
	}
	// 生成JWT token
	token, err := utils.GenerateJWTToken(int(user.ID), user.Email, "company")
	if err != nil {
		return "", errors.New("生成登录令牌失败")
	}
	return token, nil
}
