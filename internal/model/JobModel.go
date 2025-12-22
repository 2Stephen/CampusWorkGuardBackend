package model

import "time"

type JobInfo struct {
	ID           int       `gorm:"primaryKey;autoIncrement;comment:id（主键）" json:"id"`
	Name         string    `gorm:"type:varchar(100);not null;comment:岗位名称" json:"name"`
	Type         string    `gorm:"type:varchar(20);not null;comment:岗位类型" json:"type"`
	Salary       int       `gorm:"not null;comment:薪资标准" json:"salary"`
	SalaryUnit   string    `gorm:"type:varchar(20);not null;comment:薪资单位" json:"salary_unit"`
	SalaryPeriod string    `gorm:"type:varchar(20);not null;comment:薪资发放周期" json:"salary_period"`
	Content      string    `gorm:"type:text;comment:工作内容" json:"content"`
	Headcount    int       `gorm:"comment:招聘人数" json:"headcount"`
	Major        string    `gorm:"type:varchar(100);comment:专业要求" json:"major"`
	Region       string    `gorm:"type:varchar(100);comment:工作地点" json:"region"`
	RegionName   string    `gorm:"type:varchar(100);comment:工作地点名称" json:"region_name"`
	Address      string    `gorm:"type:varchar(255);comment:详细地址" json:"address"`
	Shift        string    `gorm:"type:varchar(20);comment:工作时段" json:"shift"`
	Experience   string    `gorm:"type:varchar(20);comment:经验要求" json:"experience"`
	PictureList  string    `gorm:"type:text;comment:岗位相关图片" json:"picture_list"`
	CreatedAt    time.Time `gorm:"autoCreateTime;comment:岗位发布时间" json:"created_at"`
	Status       string    `gorm:"type:varchar(20);default:pending;comment:审核状态" json:"status"`
	CompanyID    string    `gorm:"not null;comment:发布公司id" json:"company_id"`
	FailInfo     string    `gorm:"type:varchar(255);comment:审核失败原因" json:"fail_info"`
}

type JobApplication struct {
	ID        int       `gorm:"primaryKey;autoIncrement;comment:id（主键）" json:"id"`
	JobID     int       `gorm:"not null;comment:职位ID" json:"job_id"`
	StudentID int       `gorm:"not null;comment:学生用户ID" json:"student_id"`
	CreatedAt time.Time `gorm:"autoCreateTime;comment:申请时间" json:"created_at"`
	Status    string    `gorm:"type:varchar(20);default:pending;comment:申请状态" json:"status"`
	Payment   int       `gorm:"comment:薪资待遇" json:"payment"`
}

type AdminJobProfileInfo struct {
	Id         int       `json:"id"`
	Company    string    `json:"company"`
	Name       string    `json:"name"`
	Type       string    `json:"type"`
	Salary     int       `json:"salary"`
	SalaryUnit string    `json:"salaryUnit"`
	CreatedAt  time.Time `json:"createdAt"`
	Status     string    `json:"status"`
}

type StudentUserJobMatchDetail struct {
	Id         int    `json:"id"`
	Company    string `json:"company"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Salary     int    `json:"salary"`
	SalaryUnit string `json:"salaryUnit"`
	RegionName string `json:"regionName"`
	Major      string `json:"major"`
}

type JobApplicationProfileInfo struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Major        string `json:"major"`
	StudentName  string `json:"studentName"`
	StudentId    string `json:"studentId"`
	StudentMajor string `json:"studentMajor"`
	Salary       int    `json:"salary"`
	SalaryUnit   string `json:"salaryUnit"`
	Total        int    `json:"total"`
	SalaryPeriod string `json:"salaryPeriod"`
	Status       string `json:"status"`
}

type AdminJobApplicationDetail struct {
	Id           int    `json:"id"`
	Company      string `json:"company"`
	Name         string `json:"name"`
	Major        string `json:"major"`
	StudentName  string `json:"studentName"`
	StudentId    string `json:"studentId"`
	StudentMajor string `json:"studentMajor"`
	Salary       int    `json:"salary"`
	Total        int    `json:"total"`
	SalaryUnit   string `json:"salaryUnit"`
	SalaryPeriod string `json:"salaryPeriod"`
	Status       string `json:"status"`
}

type StudentUserApplicationDetail struct {
	Id           int    `json:"id"`
	Company      string `json:"company"`
	Name         string `json:"name"`
	Major        string `json:"major"`
	StudentName  string `json:"studentName"`
	StudentId    string `json:"studentId"`
	StudentMajor string `json:"studentMajor"`
	Status       string `json:"status"`
}

type AttendanceRecord struct {
	// 主键字段
	ID int `gorm:"column:id;type:int;primaryKey;autoIncrement;comment:主键id" json:"id"`
	// 关联job_applications表的主键
	JobApplicationID int `gorm:"column:job_application_id;type:int;not null;comment:job_applications表数据库主键编号（jobid+studentid确定）" json:"job_application_id"`
	// 打卡日期（注：SQL中用了varchar(64)，结构体对应string类型）
	AttendanceDate string `gorm:"column:attendance_date;type:varchar(64);not null;comment:打卡日期（优化为date类型，贴合日期场景；若需含时分秒可改用datetime）" json:"attendance_date"`
	// 打卡地点
	Location string `gorm:"column:location;type:varchar(500);default:null;comment:打卡地点（支持详细地址/坐标，长度适配高德地址返回）" json:"location"`
}

type AttendanceRecordList struct {
	AttendanceDate string `json:"attendance_date"`
	Location       string `json:"location"`
}
