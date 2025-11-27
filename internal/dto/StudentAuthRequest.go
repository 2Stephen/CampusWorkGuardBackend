package dto

type StudentAuthParams struct {
	ID     string `json:"studentId"`
	School string `json:"school"`
	Vcode  string `json:"vCode"`
	Email  string `json:"email"`
	Code   string `json:"code"`
}

type StudentLoginParams struct {
	SchoolId  string `json:"schoolId"`
	StudentId string `json:"studentId"`
	Password  string `json:"password"`
}
