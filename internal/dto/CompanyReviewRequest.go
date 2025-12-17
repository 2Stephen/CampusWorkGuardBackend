package dto

type CompanyListRequest struct {
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
	Search   string `json:"search"`
	Status   string `json:"status"`
}

type CompanyReviewRequest struct {
	Id       int    `json:"id" binding:"required"`
	Status   string `json:"status"`
	FailInfo string `json:"failInfo"`
}
