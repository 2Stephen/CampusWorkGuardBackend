package service

import (
	"CampusWorkGuardBackend/internal/dto"
	"CampusWorkGuardBackend/internal/model"
	"CampusWorkGuardBackend/internal/repository"
	"errors"
	"log"
	"time"
)

type JobDetail struct {
	//	  id: int // 岗位id
	//    name: string // 岗位名称
	//    type: string // 岗位类型
	//    salary: number // 薪资标准
	//    salaryUnit: string //薪资单位
	//    salaryPeriod: string //薪资发放周期
	//    content: string // 工作内容
	//    headcount: number //招聘人数
	//    major: string // 专业要求
	//    region: string // 工作地点
	//    address: string // 详细地址
	//    shift: string // 工作时段
	//    experience: string // 经验要求
	//    pictureList: string //岗位图片列表
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Salary       int    `json:"salary"`
	SalaryUnit   string `json:"salaryUnit"`
	SalaryPeriod string `json:"salaryPeriod"`
	Content      string `json:"content"`
	Headcount    int    `json:"headcount"`
	Major        string `json:"major"`
	Region       string `json:"region"`
	Address      string `json:"address"`
	Shift        string `json:"shift"`
	Experience   string `json:"experience"`
	PictureList  string `json:"pictureList"`
}

type JobProfileInfo struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Type       string    `json:"type"`
	Salary     int       `json:"salary"`
	SalaryUnit string    `json:"salaryUnit"`
	CreatedAt  time.Time `json:"createdAt"`
	Status     string    `json:"status"`
}

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
		CompanyID:    user.SocialCode,
	}
	// 调用存储层存储职位信息
	return repository.CreateJobInfo(info)

}

func GetCompanyUserJobInfoService(ID int) (*JobDetail, error) {
	info, err := repository.GetJobByID(ID)
	if err != nil {
		log.Println("Error retrieving job info:", err)
		return nil, err
	}
	jobDetail := &JobDetail{
		Id:           info.ID,
		Name:         info.Name,
		Type:         info.Type,
		Salary:       info.Salary,
		SalaryUnit:   info.SalaryUnit,
		SalaryPeriod: info.SalaryPeriod,
		Content:      info.Content,
		Headcount:    info.Headcount,
		Major:        info.Major,
		Region:       info.Region,
		Address:      info.Address,
		Shift:        info.Shift,
		Experience:   info.Experience,
		PictureList:  info.PictureList,
	}
	return jobDetail, nil
}

func GetCompanyUserJobListService(userID int, email string, params dto.GetCompanyUserJobListParams) ([]JobProfileInfo, int64, error) {
	user, err := repository.GetCompanyUserById(userID)
	if err != nil {
		log.Println("Error retrieving company user:", err)
		return nil, 0, err
	}
	if user == nil {
		log.Println("Company user not found with ID:", userID)
		return nil, 0, errors.New("企业用户不存在")
	}
	if user.Email != email {
		log.Println("Email mismatch for user ID:", userID)
		return nil, 0, errors.New("用户邮箱与认证邮箱不匹配")
	}
	jobInfos, total, err := repository.GetJobsByCompanyID(user.SocialCode, params)
	if err != nil {
		log.Println("Error retrieving job list:", err)
		return nil, 0, err
	}
	var jobDetails []JobProfileInfo
	for _, info := range jobInfos {
		jobDetail := JobProfileInfo{
			Id:         info.ID,
			Name:       info.Name,
			Type:       info.Type,
			Salary:     info.Salary,
			SalaryUnit: info.SalaryUnit,
			CreatedAt:  info.CreatedAt,
			Status:     info.Status,
		}
		jobDetails = append(jobDetails, jobDetail)
	}
	return jobDetails, total, nil
}
