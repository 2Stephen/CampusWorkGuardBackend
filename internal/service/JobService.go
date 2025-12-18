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
	Status       string `json:"status"`
	FailInfo     string `json:"failInfo"`
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
		RegionName:   params.RegionName,
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

func UpdateJobService(params dto.UpdateJobParams, userID int, email string) error {
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
	// 检查是否存在当前职位
	existingJob, err := repository.GetJobByID(params.Id)
	if err != nil {
		log.Println("Error retrieving job info:", err)
		return err
	}
	if existingJob.ID == 0 {
		log.Println("Job not found with ID:", params.Id)
		return errors.New("职位不存在")
	}
	if existingJob.CompanyID != user.SocialCode {
		log.Println("Unauthorized update attempt for job ID:", params.Id)
		return errors.New("无权限修改该职位信息")
	}
	info := &model.JobInfo{
		ID:           params.Id,
		Name:         params.Name,
		Type:         params.Type,
		Salary:       params.Salary,
		SalaryUnit:   params.SalaryUnit,
		SalaryPeriod: params.SalaryPeriod,
		Content:      params.Content,
		Headcount:    params.Headcount,
		Major:        params.Major,
		Region:       params.Region,
		RegionName:   params.RegionName,
		Address:      params.Address,
		Shift:        params.Shift,
		Experience:   params.Experience,
		PictureList:  params.PictureList,
		CreatedAt:    time.Now(),
		Status:       "pending",
		CompanyID:    user.SocialCode,
	}
	// 调用存储层存储职位信息
	return repository.UpdateJobInfo(info)
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
		FailInfo:     info.FailInfo,
		Status:       info.Status,
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

func DeleteJobService(ID int, userID int, email string) error {
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
	if user.Email != email {
		log.Println("Email mismatch for user ID:", userID)
		return errors.New("用户邮箱与认证邮箱不匹配")
	}
	// 检查是否存在当前职位
	existingJob, err := repository.GetJobByID(ID)
	if err != nil {
		log.Println("Error retrieving job info:", err)
		return err
	}
	if existingJob.ID == 0 {
		log.Println("Job not found with ID:", ID)
		return errors.New("职位不存在")
	}
	if existingJob.CompanyID != user.SocialCode {
		log.Println("Unauthorized delete attempt for job ID:", ID)
		return errors.New("无权限删除该职位信息")
	}
	// 调用存储层删除职位信息
	return repository.DeleteJobByID(int64(ID))
}

func GetAdminJobListService(params dto.GetAdminJobListParams) ([]model.AdminJobProfileInfo, int64, error) {
	jobInfos, total, err := repository.GetJobsForAdmin(params)
	if err != nil {
		log.Println("Error retrieving admin job list:", err)
		return nil, 0, err
	}
	return jobInfos, total, nil
}

func ReviewJobService(params dto.ReviewJobParams) error {
	// 检查是否存在当前职位
	existingJob, err := repository.GetJobByID(params.Id)
	if err != nil {
		log.Println("Error retrieving job info:", err)
		return err
	}
	if existingJob.ID == 0 {
		log.Println("Job not found with ID:", params.Id)
		return errors.New("职位不存在")
	}
	// 调用存储层审核职位信息
	return repository.ReviewJob(params.Id, params.Status, params.FailInfo)
}

func StudentUserJobMatchListService(params dto.StudentUserJobMatchListParams) ([]model.StudentUserJobMatchDetail, int, error) {
	jobInfo, total, err := repository.GetJobMatchesForStudentUser(params.SalaryOrder, params.Search, params.Region, params.Major, params.Page, params.PageSize)
	if err != nil {
		log.Println("Error retrieving student user job match list:", err)
		return nil, 0, err
	}

	return jobInfo, total, nil
}
