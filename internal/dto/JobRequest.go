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
	//    name: string
	//    status: string
	//    type: string
	//    page: 第几页，从0开始
	//    pageSize: 每页size
	Name     string `json:"name"`
	Status   string `json:"status"`
	Type     string `json:"type"`
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
}
