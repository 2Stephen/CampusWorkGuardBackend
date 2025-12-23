package dto

type SubmitComplaintParams struct {
	Title         string `json:"title"`
	CompanyId     int    `json:"companyId"`
	ComplaintType string `json:"complaintType"`
}

type GetComplaintListParams struct {
	Search   string `form:"search"`
	Page     int    `form:"page" binding:"required"`
	PageSize int    `form:"pageSize" binding:"required"`
}

type CompanyProcessComplaint struct {
	Id             int    `json:"id"`
	CompanyDefense string `json:"companyDefense"`
}

type AdminResolveComplaint struct {
	Id         int    `json:"id"`
	ResultInfo string `json:"resultInfo"`
}
