package dto

type PostJobParams struct {
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

type GetCompanyUserJobListParams struct {
	Name     string `json:"search"`
	Status   string `json:"status"`
	Type     string `json:"type"`
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
}

type UpdateJobParams struct {
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

type GetAdminJobListParams struct {
	Page     int    `json:"page" binding:"required"`
	PageSize int    `json:"pageSize" binding:"required"`
	Search   string `json:"search"`
	Status   string `json:"status"`
	Type     string `json:"type"`
}

type ReviewJobParams struct {
	Id       int    `json:"id" binding:"required"`
	Status   string `json:"status" binding:"required"`
	FailInfo string `json:"failInfo"`
}
