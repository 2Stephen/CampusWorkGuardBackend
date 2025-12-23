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
