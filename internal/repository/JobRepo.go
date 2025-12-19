package repository

import (
	"CampusWorkGuardBackend/internal/dto"
	"CampusWorkGuardBackend/internal/initialize"
	"CampusWorkGuardBackend/internal/model"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

func CreateJobInfo(info *model.JobInfo) error {
	return initialize.DB.Create(info).Error
}

func GetJobByID(ID int) (model.JobInfo, error) {
	var job model.JobInfo
	err := initialize.DB.Where("id = ?", ID).Find(&job).Error
	return job, err
}

func GetJobsByCompanyID(
	companyID string,
	params dto.GetCompanyUserJobListParams,
) ([]model.JobInfo, int64, error) {

	var (
		jobs  []model.JobInfo
		total int64
	)

	// 基础查询
	db := initialize.DB.Model(&model.JobInfo{}).
		Where("company_id = ?", companyID)

	// ===== 条件查询（空字符串不参与）=====
	if params.Name != "" {
		db = db.Where("name LIKE ?", "%"+params.Name+"%")
	}

	if params.Type != "" {
		db = db.Where("type = ?", params.Type)
	}

	if params.Status != "" {
		db = db.Where("status = ?", params.Status)
	}

	// ===== 先查 total =====
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// ===== 分页参数 =====
	page := params.Page
	pageSize := params.PageSize

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	// ===== 查询列表 =====
	if err := db.
		Order("created_at DESC").
		Limit(pageSize).
		Offset(offset).
		Find(&jobs).Error; err != nil {
		return nil, 0, err
	}
	return jobs, total, nil
}

func UpdateJobInfo(info *model.JobInfo) error {
	return initialize.DB.Model(&model.JobInfo{}).Where("id = ?", info.ID).Updates(info).Error
}

func DeleteJobByID(id int64) error {
	return initialize.DB.Delete(&model.JobInfo{}, id).Error
}

func GetJobsForAdmin(params dto.GetAdminJobListParams) ([]model.AdminJobProfileInfo, int64, error) {
	var (
		list  []model.AdminJobProfileInfo
		total int64
	)

	db := initialize.DB.Table("job_infos AS j").
		Joins("LEFT JOIN company_users AS c ON j.company_id = c.social_code")

	// ===== 条件查询 =====
	if params.Status != "" {
		db = db.Where("j.status = ?", params.Status)
	}

	if params.Type != "" {
		db = db.Where("j.type = ?", params.Type)
	}

	// 公司名称关键字查询
	if params.Search != "" {
		db = db.Where("c.company LIKE ?", "%"+params.Search+"%")
	}

	// ===== total =====
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// ===== 分页 =====
	page := params.Page
	size := params.PageSize
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}

	offset := (page - 1) * size

	// ===== 查询字段 =====
	err := db.
		Select(`
			j.id,
			j.name,
			j.type,
			j.salary,
			j.status,
			j.salary_unit,
			j.created_at,
			c.company
		`).
		Order("j.created_at DESC").
		Limit(size).
		Offset(offset).
		Scan(&list).Error

	return list, total, err
}

func ReviewJob(ID int, status string, failInfo string) error {
	// 更新审核状态和失败原因
	return initialize.DB.Model(&model.JobInfo{}).
		Where("id = ?", ID).
		Updates(map[string]interface{}{
			"status":    status,
			"fail_info": failInfo,
		}).Error
}

func GetJobMatchesForStudentUser(order, search, Region, Major string, Page, PageSize int) ([]model.StudentUserJobMatchDetail, int, error) {
	var (
		total int64
		jobs  []model.StudentUserJobMatchDetail
	)

	db := initialize.DB.Table("job_infos AS j").
		Joins("LEFT JOIN company_users AS c ON j.company_id = c.social_code")

	// ===== 条件查询 =====
	if Region != "" {
		db = db.Where("j.region = ?", Region)
	}

	if Major != "" && Major != "ANY" {
		db = db.Where("j.major = ? OR j.major = 'ANY'", Major)
	}

	db = db.Where("j.status = ?", "approved")

	// 岗位名称关键字查询
	if search != "" {
		db = db.Where("j.name LIKE ?", "%"+search+"%")
	}
	// 查询hc > 0的职位
	db = db.Where("j.headcount > ?", 0)
	// ===== total =====
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// ===== 分页 =====
	if Page <= 0 {
		Page = 1
	}
	if PageSize <= 0 {
		PageSize = 10
	}

	offset := (Page - 1) * PageSize

	// ===== 查询字段 =====
	q := db.Select(`
		j.id,
		j.name,
		j.type,
		j.salary,
		j.status,
		j.salary_unit,
		j.region_name,
		c.company,
		j.major
	`)
	// 定义薪资排序逻辑，兼容 day/month/hour 三种单位
	var salarySortExpr string
	if order == "DESC" || order == "ASC" {
		salarySortExpr = fmt.Sprintf(
			"CASE salary_unit WHEN 'hour' THEN salary * 8 * 22 WHEN 'day' THEN salary * 22 ELSE salary END %s",
			order,
		)
		q = q.Order(salarySortExpr)
	} else {
		// 默认按创建时间降序
		q = q.Order("j.created_at DESC")
	}

	err := q.Limit(PageSize).Offset(offset).Scan(&jobs).Error

	return jobs, int(total), err
}

func HasStudentUserAppliedJob(studentID int, jobID int) (bool, error) {
	var count int64
	err := initialize.DB.Model(&model.JobApplication{}).
		Where("student_id = ? AND job_id = ?", studentID, jobID).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// CreateStudentUserJobApplication 新增求职申请（事务内扣减岗位招聘人数）
func CreateStudentUserJobApplication(studentID int, jobID int) error {
	// 开启GORM事务，封装两张表操作
	return initialize.DB.Transaction(func(tx *gorm.DB) error {
		// 步骤1：查询岗位信息并加悲观锁，防止并发扣减导致超招
		var jobInfo model.JobInfo
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}). // 加行锁，避免并发问题
										Where("id = ?", jobID).First(&jobInfo).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("岗位不存在")
			}
			return err // 其他查询错误，触发事务回滚
		}

		// 步骤2：校验招聘人数（不能为负）
		if jobInfo.Headcount <= 0 {
			return errors.New("该岗位招聘人数已用尽，无法申请")
		}

		// 步骤3：扣减岗位招聘人数（headcount--）
		if err := tx.Model(&jobInfo).Update("headcount", gorm.Expr("headcount - ?", 1)).Error; err != nil {
			return err // 扣减失败，触发回滚
		}

		// 步骤4：插入求职申请记录
		application := model.JobApplication{
			JobID:     jobID,
			StudentID: studentID,
			CreatedAt: time.Now(),
		}
		if err := tx.Create(&application).Error; err != nil {
			return err // 插入失败，触发回滚
		}

		// 所有操作成功，返回nil触发事务提交
		return nil
	})
}

func GetJobApplicationsByCompanySocialCode(socialCode string, params dto.GetJobApplicationListParams) ([]model.JobApplicationProfileInfo, int64, error) {
	var (
		list  []model.JobApplicationProfileInfo
		total int64
	)

	db := initialize.DB.Table("job_applications AS ja").
		Joins("LEFT JOIN job_infos AS j ON ja.job_id = j.id").
		Joins("LEFT JOIN student_users AS s ON ja.student_id = s.id").
		Joins("LEFT JOIN chsi_student_infos AS c ON s.email = c.email").
		Where("j.company_id = ?", socialCode)

	// ===== 条件查询 =====
	if params.Status != "" {
		db = db.Where("ja.status = ?", params.Status)
	}

	if params.Search != "" {
		db = db.Where("j.name LIKE ?", "%"+params.Search+"%")
	}
	// ===== total =====
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// ===== 分页 =====
	page := params.Page
	size := params.PageSize
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}
	offset := (page - 1) * size

	// ===== 查询字段 =====
	err := db.
		Select(`
			ja.id,
			j.name,
			c.name AS student_name,
			s.student_id,
			c.major AS student_major,
			j.major,
			j.salary,
			j.salary_unit,
			j.salary_period,
			ja.status
		`).
		Order("ja.created_at DESC").
		Limit(size).
		Offset(offset).
		Scan(&list).Error

	return list, total, err
}
