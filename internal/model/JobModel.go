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
	Address      string    `gorm:"type:varchar(255);comment:详细地址" json:"address"`
	Shift        string    `gorm:"type:varchar(20);comment:工作时段" json:"shift"`
	Experience   string    `gorm:"type:varchar(20);comment:经验要求" json:"experience"`
	PictureList  string    `gorm:"type:text;comment:岗位相关图片" json:"picture_list"`
	CreatedAt    time.Time `gorm:"autoCreateTime;comment:岗位发布时间" json:"created_at"`
	Status       string    `gorm:"type:varchar(20);default:pending;comment:审核状态" json:"status"`
	CompanyID    int       `gorm:"not null;comment:发布公司id" json:"company_id"`
}
