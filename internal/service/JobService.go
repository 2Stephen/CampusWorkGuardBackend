package service

import (
	"CampusWorkGuardBackend/internal/dto"
	"CampusWorkGuardBackend/internal/model"
	"CampusWorkGuardBackend/internal/repository"
	"errors"
	"log"
	"time"
)

func PostJobService(params dto.PostJobParams, userID int, email string) error {
	// 检查数据库是否存在当前公司用户
	user, err := repository.GetCompanyUserById(userID)
	if err != nil {
		log.Println("Error retrieving company user:", err)
		return err
	}
	if user == nil {
		log.Println("Company user not found with ID:", userID)
		return errors.New("企业用户不存在")
	}
	if user.VerifyStatus != "verified" {
		return errors.New("企业用户未通过认证，无法发布职位")
	}
	if user.Email != email {
		log.Println("Email mismatch for user ID:", userID)
		return errors.New("用户邮箱与认证邮箱不匹配")
	}
	info := &model.JobInfo{
		Name:         params.Name,
		Type:         params.Type,
		Salary:       params.Salary,
		SalaryUnit:   params.SalaryUnit,
		SalaryPeriod: params.SalaryPeriod,
		Content:      params.Content,
		Headcount:    params.Headcount,
		Major:        params.Major,
		Region:       params.Region,
		Address:      params.Address,
		Shift:        params.Shift,
		Experience:   params.Experience,
		PictureList:  params.PictureList,
		CreatedAt:    time.Now(),
		Status:       "pending",
		CompanyID:    userID,
	}
	// 调用存储层存储职位信息
	return repository.CreateJobInfo(info)

}
