package dto

type SubmitComplaintParams struct {
	Title         string `json:"title"`
	CompanyId     int    `json:"companyId"`
	ComplaintType string `json:"complaintType"`
}
